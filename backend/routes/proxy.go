package routes

import (
	"io"
	"log"
	"net/http"
	"strings"
	"../appconfigs"
)

func ReloadProxy(w http.ResponseWriter, r *http.Request , router http.HandlerFunc) {

	path := r.URL.Path
	if !strings.Contains(path, ".") { // if it's not a resource then continue to the router as normally.
		router(w, r)
		return
	}


	log.Println("Debug, Hot reload", r.Host)


	str_H5URL := "http://"+ appconfigs.GetInst().G_H5Host + r.RequestURI
	resp, err := http.Get( str_H5URL )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
