package app

import (
	"fmt"
	"net/http"
)

const CONTENT_TYPE_JSON = "application/json"
const CONTENT_TYPE_XML = "application/xml"

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
