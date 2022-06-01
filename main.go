package main

import (
	
	"net/http"
	"strconv"
	"html/template"
	"log"
	"fmt"
)
func main()(
	
	server :=http.Server{
		Addr:"127.0.0.1:8080",
	}

	server.ListenAndServe();
)