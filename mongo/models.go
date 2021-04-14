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

package mongo

const (
	idField = "_id"
	db      = "safo"
)

// SessionState type
type SessionState string

// SessionState values
const (
	Pending     SessionState = "PENDING"
	Active      SessionState = "ACTIVE"
	Blocked     SessionState = "BLOCKED"
	Invalidated SessionState = "INVALIDATED"
)
