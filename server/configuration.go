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
	"os"

	mongo "image-loader/internal/mongo"

	"github.com/go-openapi/runtime/middleware"
	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

// Server .
type Server struct {
	Router   *mux.Router
	Database mongo.DatabaseService
}

// Init configuration
func (server *Server) InitHTTPServer(databasepath string) error {
	server.InitDatabase(databasepath)
	server.Router = mux.NewRouter()
	server.InitRouters()

	err := server.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err != nil {
			return fmt.Errorf("error getting path templates: %w", err)
		}

		methods, err := route.GetMethods()
		if err != nil {
			return fmt.Errorf("error getting methods: %w", err)
		}

		fmt.Printf("routes %s %s \n", methods, template)

		return nil
	})
	if err != nil {
		return fmt.Errorf("error when initializing routes -> %w", err)
	}

	server.InitOpenAPIRouters()
	server.Router.Handle("/", http.FileServer(http.Dir("./")))

	return nil
}

// Run .
func (server *Server) Run(host string) {
	fmt.Printf("Listening to: %v \n", host)
	handler := cors.Default().Handler(server.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}

// InitRouters .
func (server *Server) InitRouters() {
	// swagger:operation POST /images/raw LoadNewRawSatelliteImage
	//
	// Loads new raw image
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
	server.Router.HandleFunc("/images/raw/", SetMiddlewareJSON(server.LoadNewRawSatelliteImage)).Methods("POST")

	// swagger:operation GET /images/raw GetRawSatelliteImage
	//
	// Retrieves a raw image
	//
	// @WIP
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: query
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
	//     description: I/O conflict while writing to local storage
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
	server.Router.HandleFunc("/images/raw", SetMiddlewareJSON(server.GetRawSatelliteImage)).Methods("GET")

	// swagger:operation POST /images/processed LoadNewProcessedSatelliteImage
	//
	// Loads a new processed image and it's processing results
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
	server.Router.HandleFunc("/images/processed", SetMiddlewareJSON(server.LoadNewProcessedSatelliteImage)).Methods("POST")

	// swagger:operation GET /images/processed GetProcessedSatelliteImage
	//
	// Retrieves processed image
	//
	// @WIP
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: query
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
	//     description: I/O conflict while writing to local storage
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
	server.Router.HandleFunc("/images/processed", SetMiddlewareJSON(server.GetProcessedSatelliteImage)).Methods("GET")
}

func (server *Server) InitOpenAPIRouters() {
	if _, err := os.Stat(swaggerpath); err != nil {
		fmt.Printf("cannot open swagger file %v \n", err)
	}

	server.Router.Handle(swaggerpath, http.FileServer(http.Dir("./")))

	if os.Getenv(ENVIRONMENT) == "production" {
		opts := middleware.RedocOpts{SpecURL: swaggerpath}
		redoc := middleware.Redoc(opts, nil)
		server.Router.Handle("/docs", redoc)
	}

	opts := middleware.SwaggerUIOpts{SpecURL: swaggerpath}
	swagger := middleware.SwaggerUI(opts, nil)
	server.Router.Handle("/docs", swagger)
}

// InitDatabase .
func (server *Server) InitDatabase(databasepath string) {
	imageDao, err := mongo.InitiateImageDao(databasepath)
	if err != nil {
		log.Fatalf(connectionError, err)
	}
	server.Database = mongo.NewImageService(imageDao)
}

// SetMiddlewareJSON .
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, appJSON)
		next(w, r)
	}
}
