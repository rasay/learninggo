package main

import (
	"net/http"

	"github.com/rkasay/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	// launches an HTTP server
	http.ListenAndServe(":3000", nil)
}
