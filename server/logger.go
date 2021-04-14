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
	 "log"
	 "net/http"
	 "time"
 )
 
 func Logger(inner http.Handler, name string) http.Handler {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 start := time.Now()
 
		 inner.ServeHTTP(w, r)
 
		 log.Printf(
			 "%s %s %s %s",
			 r.Method,
			 r.RequestURI,
			 name,
			 time.Since(start),
		 )
	 })
 }
 