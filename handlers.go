package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	exeTmpl(w, r, nil, "main.tmpl")
}
func assets(w http.ResponseWriter, r *http.Request) {
	exeTmpl(w, r, nil, "assets-page.tmpl")
}
func info(w http.ResponseWriter, r *http.Request) {
	exeTmpl(w, r, nil, "info-page.tmpl")
}
