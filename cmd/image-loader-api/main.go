package main

import (
	"context"
	"fmt"
	imageloaderapi "image-loader"
	raw_images "image-loader/gen/raw_images"
	processed_images "image-loader/gen/processed_images"
	"image-loader/mongo"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/namsral/flag"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.

	fset := flag.NewFlagSetWithEnvPrefix("api", "IL", flag.ExitOnError)

	var (
		domainF   = fset.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = fset.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = fset.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = fset.Bool("debug", false, "Log request and response bodies")
		//period, gcPeriod uint
		//properties       string
		mdbStr string
	)
	{
		fset.StringVar(&mdbStr, "mongodb", "mongodb://localhost:27017", "mongodb url")
		//fset.StringVar(&properties, "properties", "/run/secrets/sm.properties", "Path used to retrieve properties")
		//fset.UintVar(&gcPeriod, "gcperiod", MinutesInADay, "garbage collector period in minutes, must be greather than 10, default once per day")
	}

	fset.Parse(os.Args[1:])

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "", 0)
	}
	log.SetPrefix("")
	log.SetFlags(0)

	/*config, err := ReadPropertiesFile(properties)
	if err != nil {
		logger.Fatalf("Cannot read property file defined in path %s", properties)
	}*/

	logger.Println("**** readed config ****")
	logger.Printf("http-port:%s", *httpPortF)
	logger.Printf("mongodb:%s", mdbStr)
	//logger.Printf("properties:%s", properties)
	//logger.Printf("period:%d", period)
	//logger.Printf("gcperiod:%d", gcPeriod)
	logger.Println("**** end readed config ****")

	// Dao creation
	/*var imageDao *mongo.ImageDao
	imageDao, err := mongo.AddNewImageDao(mdbStr, logger)
	if err != nil {
		logger.Fatal(fmt.Errorf("Error creating imageDao: %s", err.Error()))
	}*/

	//mongo client creation
	imageDao, err := mongo.InitiateImageDao(mdbStr)
	if err != nil {
		logger.Fatal(fmt.Errorf("Error creating imageDao: %s", err.Error()))
	}

	// Garbage collector
	/*gc := gc.NewGarbageCollector(gcPeriod, userDao)
	gc.Start()*/

	// Initialize the services.
	var (
		imagesSrv raw_images.Service
		processedImagesSrv processed_images.Service
	)
	{
		imagesSrv = imageloaderapi.NewImageLoader(logger, imageDao)
		processedImagesSrv = imageloaderapi.NewProcessedImageLoader(logger, imageDao)

	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		imagesEndpoints *raw_images.Endpoints
		processedimagesEndpoints *processed_images.Endpoints
	)
	{
		imagesEndpoints = raw_images.NewEndpoints(imagesSrv)
		processedimagesEndpoints = processed_images.NewEndpoints(processedImagesSrv)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.

	addr := "http://0.0.0.0:8080"
	u, err := url.Parse(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
		os.Exit(1)
	}
	if *secureF {
		u.Scheme = "https"
	}
	if *domainF != "" {
		u.Host = *domainF
	}
	if *httpPortF != "" {
		h := strings.Split(u.Host, ":")[0]
		u.Host = h + ":" + *httpPortF
	} else if u.Port() == "" {
		u.Host += ":80"
	}
	handleHTTPServer(ctx, u, imagesEndpoints, processedimagesEndpoints, &wg, errc, logger, *dbgF)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

// ConfigProperties type to store properties values
/*
type ConfigProperties map[string]interface{}

// ReadPropertiesFile and store them in a ConfigProperties type
func ReadPropertiesFile(filepath string) (ConfigProperties, error) {
	config := ConfigProperties{}

	if len(filepath) == 0 {
		return config, nil
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equalsIndex := strings.Index(line, "=")
		isEsProperty := strings.Contains(line, "sm")
		if equalsIndex >= 0 && isEsProperty {
			if key := strings.TrimSpace(line[:equalsIndex]); len(key) > 0 {
				var value string
				if len(line) > equalsIndex {
					value = strings.TrimSpace(line[equalsIndex+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
*/
