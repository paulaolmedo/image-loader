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

package server

import (
	"fmt"
	"log"
	"net/http"

	mongo "image-loader/mongo"

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
	// swagger:operation POST /images/raw LoadNewRawSatelliteImage
	//
	// Load new raw image
	//
	// @WIP
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: body
	//   name: raw satellite image
	//   description: Image information
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/RawSatelliteImage"
	// responses:
	//   '200':
	//     description: Raw image details (bytes written into the database)
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/RawSatelliteImage"
	//   '409':
	//     description: Raw image already exists
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	//   '500':
	//     description: Internal Server Error
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	serverConfiguration.Router.HandleFunc("/images/raw", SetMiddlewareJSON(serverConfiguration.LoadNewRawSatelliteImage)).Methods("POST")

	// swagger:operation GET /images/raw/{filename} GetRawSatelliteImage
	//
	// Retrieve raw image
	//
	// @WIP
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: filename
	//   description: raw image identification
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Raw image details (bytes written to local storage)
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/RawSatelliteImage"
	//   '409':
	//     description: I/O conflict while writting to local storage
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	//   '500':
	//     description: Internal Server Error
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	serverConfiguration.Router.HandleFunc("/images/raw", SetMiddlewareJSON(serverConfiguration.GetRawSatelliteImage)).Methods("GET")

	// swagger:operation POST /images/processed LoadNewProcessedSatelliteImage
	//
	// Load new processed image
	//
	// @WIP
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: body
	//   name: processed satellite image
	//   description: Image information
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ProcessedSatelliteImage"
	// responses:
	//   '200':
	//     description: Processed image details (bytes written into the database)
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ProcessedSatelliteImage"
	//   '409':
	//     description: Processed image already exists
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	//   '500':
	//     description: Internal Server Error
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	serverConfiguration.Router.HandleFunc("/images/processed", SetMiddlewareJSON(serverConfiguration.LoadNewProcessedSatelliteImage)).Methods("POST")

	// swagger:operation GET /images/processed/{filename} GetProcessedSatelliteImage
	//
	// Retrieve processed image
	//
	// @WIP
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: filename
	//   description: processed image identification
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: processed image details (bytes written to local storage)
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ProcessedSatelliteImage"
	//   '409':
	//     description: I/O conflict while writting to local storage
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	//   '500':
	//     description: Internal Server Error
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ModelError"
	serverConfiguration.Router.HandleFunc("/images/processed", SetMiddlewareJSON(serverConfiguration.GetProcessedSatelliteImage)).Methods("GET")

}

//InitDatabase .
func (serverConfiguration *Server) InitDatabase(databasepath string) {
	imageDao, err := mongo.InitiateImageDao(databasepath)

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
