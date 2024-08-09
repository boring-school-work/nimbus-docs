package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// define base url and port
var (
	baseURL string = "/nimbus/api/v1/"
	port    string = ":3211"
)

// getURI returns a string that represents the HTTP method and path
func getURI(method, endpoint string) string {
	return method + " " + baseURL + endpoint
}

// SwaggerSpecHandler handles a GET request to serve the swagger config file
func SwaggerSpecHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)
	http.ServeFile(w, r, "./swagger.yaml")
}

func main() {
	mux := http.NewServeMux()

	// serve swagger docs
	mux.HandleFunc(getURI("GET", "swagger/swagger.yaml"), SwaggerSpecHandler)

	// serve Swagger UI
	mux.HandleFunc(
		getURI("GET", "swagger/ui/"),
		httpSwagger.Handler(httpSwagger.URL("http://localhost:3211/nimbus/api/v1/swagger/swagger.yaml")),
	)

	log.Printf("Starting api docs at http://localhost%s%s%s\n", port, baseURL, "swagger/ui")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Println(err)
	}
}
