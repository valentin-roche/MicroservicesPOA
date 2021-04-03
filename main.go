package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	blogPOA "github.com/valentin-roche/MicroservicesPOA/blog"
)

func main() {

	service := blogPOA.NewInmemBlogPostService()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := http.ListenAndServe(":8000", blogPOA.MakeHTTPHandler(service, logger))
	if err != nil {
		panic(err)
	}
}
