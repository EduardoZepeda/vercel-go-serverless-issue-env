package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Api(w http.ResponseWriter, r *http.Request) {
	mode, ok := os.LookupEnv("MODE")
	fmt.Fprintf(w, fmt.Sprintf("<h1>Hello from Go! Mode set by MODE environmetal variable is: %s loaded: %t</h1>", mode, ok))
}
