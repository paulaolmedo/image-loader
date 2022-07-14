//Copyright (c) 2021 PAULA B. OLMEDO.
//
//This file is part of IMAGE_LOADER
//(see https://github.com/paulaolmedo/image-loader).
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program. If not, see <http://www.gnu.org/licenses/>.

// Image loader
//
// Service to manage satellite imagery and it's data
//
//     Schemes: http
//     Host: 0.0.0.0:8080
//	   BasePath: /
//     Version: 1.0.0
//	   License: GPLv3 https://www.gnu.org/licenses/gpl-3.0.html
//     Contact: Satellite Forecast<info@satelliteforecast.system>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	server "image-loader/server"
	"log"
	"os"

	"github.com/magiconair/properties"
)

//go:generate swagger generate spec -m -o ../docs/swagger.yml

func main() {
	propertiesPath := os.Getenv("IL_PROPERTIES")
	p := properties.MustLoadFile(propertiesPath, properties.UTF8)

	host := p.MustGetString("host")
	databaseHost := p.MustGetString("database_host")

	var configuration server.Server

	environment := p.MustGetString("environment")
	os.Setenv(server.ENVIRONMENT, environment)

	if err := configuration.InitHTTPServer(databaseHost); err != nil {
		log.Fatal(err)
	}

	configuration.Run(host)
}
