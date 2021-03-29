package main

import (
	"net/http"
	"https://github.com/valentin-roche/MicroservicesPOA/blog"
)

func main() {

	service := blogPOA.MakeBlogPostService()

	endpoints := blogPOA.MakeBlogPostEndpoints(service)

	err := http.ListenAndServe(":8000", blogPOA.MakeHTTPHandler(endpoints))
	if err != nil {
		panic(err)
	}
}
