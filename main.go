package imageloaderapi

import (
	"github.com/magiconair/properties"
)

func main() {
	p := properties.MustLoadFile("app/app.properties", properties.UTF8)

	host := p.MustGetString("host")
	databaseHost := p.MustGetString("database_host")

	configuration := Server{}
	configuration.InitHTTPServer(databaseHost)
	configuration.Run(host)

}
