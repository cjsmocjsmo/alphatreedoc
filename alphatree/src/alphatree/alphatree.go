///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2019, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
package main

import (
	"os"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func ShowMain(w http.ResponseWriter, r *http.Request) {
	tmppath := "./static/templates/alphatree.template"
	tmpl := template.Must(template.ParseFiles(tmppath))
	tmpl.Execute(w, tmpl)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/alphatree", ShowMain)
	r.PathPrefix("/static").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	raw_port := os.Getenv("PORT")
	var port string
	if raw_port != "" {
		port =  ":" + raw_port
	} else {
		port = ":80"	
	}
	http.ListenAndServe(port, (r))
}