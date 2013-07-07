// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

// This file implements a stand-alone blog server.

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var gopath = os.Getenv("GOPATH")
	var contentPath = gopath + "/src/github.com/tiancaiamao//go.blog/content/"
	var templatePath = gopath + "/src/github.com/tiancaiamao/go.blog/template/"
	var staticPath = gopath + "/src/github.com/tiancaiamao/go.blog/static/"

	
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s, err := NewServer(contentPath, templatePath)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", s)
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
