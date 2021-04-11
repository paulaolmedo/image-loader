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
	 "net/http"
 
	 mux "github.com/gorilla/mux"
 )
 
 type Route struct {
	 Name        string
	 Method      string
	 Pattern     string
	 HandlerFunc http.HandlerFunc
 }
 
 type Routes []Route
 
 //NewRouter .
 func NewRouter() *mux.Router {
	 router := mux.NewRouter().StrictSlash(true)
 
	 for _, route := range routes {
		 var handler http.Handler
		 handler = route.HandlerFunc
		 handler = Logger(handler, route.Name)
 
		 router.
			 Methods(route.Method).
			 Path(route.Pattern).
			 Name(route.Name).
			 Handler(handler)
	 }
 
	 return router
 }
 
 func Index(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprintf(w, "Hello World!")
 }
 
 var routes = Routes{
	 Route{
		 "Index",
		 "GET",
		 "/v1/",
		 Index,
	 },
 }
 