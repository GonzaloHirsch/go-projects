package main

import (
	"fmt"
	"net/http"
	"projects/crypto"
)

var SECURITY = map[string][]string{
	ENCRYPT_SINGLE: {
		"GET",
	},
	ENCRYPT_MULTIPLE: {
		"POST",
	},
	ENCRYPT_MULTIPLE_PARALLEL: {
		"POST",
	},
}

const (
	ENCRYPT_SINGLE                    = "/encrypt-single"
	ENCRYPT_MULTIPLE                  = "/encrypt-multiple"
	ENCRYPT_MULTIPLE_PARALLEL         = "/encrypt-multiple-parallel"
	ENCRYPT_SINGLE_MESSAGE            = "msg"
	ENCRYPT_MULTIPLE_MESSAGE          = "msgs"
	ENCRYPT_MULTIPLE_PARALLEL_MESSAGE = "msgs"
)

// Checks if the request method is the appropiate one
func enforceMethod(key string, req *http.Request) bool {
	method := req.Method
	sec := SECURITY[key]
	for _, s := range sec {
		if s == method {
			return true
		}
	}
	return false
}

func encryptSingle(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod(ENCRYPT_SINGLE, req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Recover the values from the query params
	msgs, ok := req.URL.Query()[ENCRYPT_SINGLE_MESSAGE]
	if !ok {
		http.Error(w, "Missing msg parameter.", http.StatusBadRequest)
		return
	}

	// Creating the hash
	h := fmt.Sprintf("%x\n", crypto.Sha256([]byte(msgs[0])))

	// Sending the response
	fmt.Fprintf(w, h)
}

func encryptMultiple(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod(ENCRYPT_MULTIPLE, req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get messages to hash
	fmt.Println(req.PostForm)
	msgs, ok := req.PostForm[ENCRYPT_MULTIPLE_MESSAGE]
	if !ok {
		http.Error(w, "Missing msgs parameter.", http.StatusBadRequest)
		return
	}

	response := ""

	for _, msg := range msgs {
		// Creating the hash
		h := fmt.Sprintf("%v --> %x\n", msg, crypto.Sha256([]byte(msg)))
		response += h
	}

	// Sending the response
	fmt.Fprintf(w, response)
}

func encryptMultipleParallel(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod(ENCRYPT_MULTIPLE, req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get messages to hash
	fmt.Println(req.PostForm)
	msgs, ok := req.PostForm[ENCRYPT_MULTIPLE_MESSAGE]
	if !ok {
		http.Error(w, "Missing msgs parameter.", http.StatusBadRequest)
		return
	}

	// Sending the threads to run the hash generations
	c := make(chan [32]byte)
	for _, msg := range msgs {
		go func(m string) {
			h := crypto.Sha256([]byte(m))
			c <- h
		}(msg)
	}

	// Retrieving the hashes and building the responses
	response := ""
	for _, msg := range msgs {
		// Creating the hash
		h := fmt.Sprintf("%v --> %x\n", msg, <-c)
		response += h
	}

	// Sending the response
	fmt.Fprintf(w, response)
}

func main() {
	fmt.Println("Starting server on port 8090...")

	http.HandleFunc(ENCRYPT_SINGLE, encryptSingle)
	http.HandleFunc(ENCRYPT_MULTIPLE, encryptMultiple)
	http.HandleFunc(ENCRYPT_MULTIPLE_PARALLEL, encryptMultipleParallel)

	http.ListenAndServe(":8090", nil)
}
