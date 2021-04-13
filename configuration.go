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
