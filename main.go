package main

import (
	"github.com/pkg/browser"
	// https://github.com/gorilla/mux#examples
	"github.com/gorilla/mux"
	"time"
	"net/http"
	"log"
	// The simple and easy way to embed static files into Go binaries.
	"github.com/gobuffalo/packr"
	"fmt"
	"flag"
	"github.com/jlaso/go-redis-tools/routing"
	"github.com/jlaso/go-redis-tools/common"
)


func main() {

	var live string

	cfg := common.GetConfiguration()

	flag.StringVar(&live, "live", "N", "serve files from folder (Y) or from packer.box (N)")
	flag.Parse()

	r := mux.NewRouter()
	r.Host(cfg.Url)
	r.HandleFunc("/redis", routing.RedisHandler)
	r.HandleFunc("/servers", routing.ServersHandler)
	r.HandleFunc("/top-menu", routing.TopMenuHandler)
	r.HandleFunc("/add-server", routing.AddRedisPostHandler).Methods("POST")

	if live == "Y" {
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(packr.NewBox("./files"))))
	} else {
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./files"))))
	}

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.Url,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fullUrl := "http://" + cfg.Url

	fmt.Printf("Opening UI in "+fullUrl+" [live:%s], you can open in a browser if you close it!\n", live)

	browser.OpenURL(fullUrl)

	log.Fatal(srv.ListenAndServe())
}
