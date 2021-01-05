package main

import (
	"fmt"
	"net/http"
	"projects/urlShortener"
	"strings"
)

///////////////////////////////////////////////////////////////
// Variables
///////////////////////////////////////////////////////////////

const (
	ADD_URL            = "/urls"
	DELETE_URL         = "/urls/delete"
	GET_URL            = "/goto/"
	GET_URL_VISITS     = "/urls/visits/"
	GET_URL_VISITS_ALL = "/urls/visits/all"
	GET_URLS_ALL       = "/urls/all"
	PARAM_URL          = "url"
	PARAM_ALIAS        = "alias"
)

var s urlShortener.Shortener

///////////////////////////////////////////////////////////////
// General Methods
///////////////////////////////////////////////////////////////

// Checks if the request method is the appropiate one
func enforceMethod(m string, req *http.Request) bool {
	method := req.Method
	if m == method {
		return true
	}
	return false
}

///////////////////////////////////////////////////////////////
// Endpoints
///////////////////////////////////////////////////////////////

func addUrl(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod("POST", req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get the data to add
	url, url_ok := req.PostForm[PARAM_URL]
	if !url_ok {
		http.Error(w, "Missing url parameter.", http.StatusBadRequest)
		return
	}
	alias, alias_ok := req.PostForm[PARAM_ALIAS]
	if !alias_ok {
		http.Error(w, "Missing alias parameter.", http.StatusBadRequest)
		return
	}

	// Add the url
	e := s.AddUrl(alias[0], url[0])
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	// Generate the response
	h := fmt.Sprintf("Added url \"%v\" under alias \"%v\"\n", url[0], alias[0])

	// Sending the response
	fmt.Fprintf(w, h)
}

func removeUrl(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod("DELETE", req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get the data to add
	alias, alias_ok := req.URL.Query()[PARAM_ALIAS]
	if !alias_ok {
		http.Error(w, "Missing alias parameter.", http.StatusBadRequest)
		return
	}

	// Add the url
	val, e := s.RemoveUrl(alias[0])
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	// Generate the response
	h := fmt.Sprintf("Removed url \"%v\" using alias \"%v\"\n", val, alias[0])

	// Sending the response
	fmt.Fprintf(w, h)
}

func getUrlVisits(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod("GET", req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get the data to add
	alias := strings.TrimPrefix(req.URL.Path, GET_URL_VISITS)

	// Add the url
	val, e := s.GetUrlVisits(alias)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	// Generate the response
	h := fmt.Sprintf("Alias \"%v\" has %v visits\n", alias,  val)

	// Sending the response
	fmt.Fprintf(w, h)
}

func getAllVisits(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod("GET", req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Add the url
	val := s.GetGeneralVisits()

	// Generate the response
	h := fmt.Sprintf("Visits %v\n", val)

	// Sending the response
	fmt.Fprintf(w, h)
}

func getAllUrls(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod("GET", req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Add the url
	val := s.GetGeneralUrls()

	// Generate the response
	h := fmt.Sprintf("Urls %v\n", val)

	// Sending the response
	fmt.Fprintf(w, h)
}

func getUrl(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod("GET", req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get the data to add
	alias := strings.TrimPrefix(req.URL.Path, GET_URL)

	// Add the url
	val, e := s.GetUrl(alias)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	val = strings.TrimPrefix(val, "https://")
	val = strings.TrimPrefix(val, "http://")

	http.Redirect(w, req, "http://" + val, http.StatusMovedPermanently)
}

func main() {
	fmt.Println("Starting server on port 8090...")

	// Initializing the shortener
	s = urlShortener.InitShortener()

	http.HandleFunc(GET_URL, getUrl)
	http.HandleFunc(GET_URLS_ALL, getAllUrls)
	http.HandleFunc(GET_URL_VISITS_ALL, getAllVisits)
	http.HandleFunc(GET_URL_VISITS, getUrlVisits)
	http.HandleFunc(DELETE_URL, removeUrl)
	http.HandleFunc(ADD_URL, addUrl)

	e := http.ListenAndServe(":8090", nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}
