package imageloaderapi

import (
	"fmt"
	"log"
	"net/http"

	"image-loader/mongo"
	dao "image-loader/mongo"

	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

// Server .
type Server struct {
	Router   *mux.Router
	Database mongo.DatabaseService
}

//Init configuration
func (serverConfiguration *Server) InitHTTPServer(databasepath string) {
	serverConfiguration.InitDatabase(databasepath)
	serverConfiguration.Router = mux.NewRouter()
	serverConfiguration.InitRouters()

	serverConfiguration.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Printf("routes %s %s", methods, template)
		fmt.Println()
		return nil
	})
}

//Run .
func (serverConfiguration *Server) Run(host string) {
	fmt.Println("Listening to:", host)
	handler := cors.Default().Handler(serverConfiguration.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}

//InitRouters .
func (serverConfiguration *Server) InitRouters() {
	serverConfiguration.Router.HandleFunc("/raw-images", SetMiddlewareJSON(serverConfiguration.LoadNewRawSatelliteImage)).Methods("POST")
	serverConfiguration.Router.HandleFunc("/raw-images", SetMiddlewareJSON(serverConfiguration.GetRawSatelliteImage)).Methods("GET")
	serverConfiguration.Router.HandleFunc("/processed-images", SetMiddlewareJSON(serverConfiguration.LoadNewProcessedSatelliteImage)).Methods("POST")
	serverConfiguration.Router.HandleFunc("/processed-images", SetMiddlewareJSON(serverConfiguration.GetProcessedSatelliteImage)).Methods("GET")

}

//InitDatabase .
func (serverConfiguration *Server) InitDatabase(databasepath string) {
	imageDao, err := dao.InitiateImageDao(databasepath)

	if err != nil {
		log.Fatalf("failed to connect %v ", err)
	}
	serverConfiguration.Database = mongo.NewImageService(imageDao)

}

//SetMiddlewareJSON .
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
