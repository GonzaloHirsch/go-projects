package main

import (
	"fmt"
	"net/http"
	"projects/blockchain"
	"strconv"
)

///////////////////////////////////////////////////////////////
// Variables
///////////////////////////////////////////////////////////////

var SECURITY = map[string][]string{
	CREATE_TRANSACTION: {
		"POST",
	},
	MINE: {
		"POST",
	},
	CHAIN: {
		"GET",
	},
}

const (
	CREATE_TRANSACTION = "/transactions"
	MINE = "/mine"
	CHAIN = "/chain"
	CREATE_TRANSACTION_SENDER = "sender"
	CREATE_TRANSACTION_RECEIVER = "receiver"
	CREATE_TRANSACTION_AMOUNT = "amount"
)

var b blockchain.Blockchain
const serverAddress string = "SERVER_ADDRESS"

///////////////////////////////////////////////////////////////
// General Methods
///////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////
// Endpoints
///////////////////////////////////////////////////////////////

func createTransaction(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod(CREATE_TRANSACTION, req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Parse the form in order to access it
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get the data to add
	sender, sender_ok := req.PostForm[CREATE_TRANSACTION_SENDER]
	if !sender_ok {
		http.Error(w, "Missing sender parameter.", http.StatusBadRequest)
		return
	}
	receiver, receiver_ok := req.PostForm[CREATE_TRANSACTION_RECEIVER]
	if !receiver_ok {
		http.Error(w, "Missing receiver parameter.", http.StatusBadRequest)
		return
	}
	amount, amount_ok := req.PostForm[CREATE_TRANSACTION_AMOUNT]
	if !amount_ok {
		http.Error(w, "Missing amount parameter.", http.StatusBadRequest)
		return
	}
	_amount, _amount_ok := strconv.ParseFloat(amount[0], 32)
	if _amount_ok != nil {
		http.Error(w, "Invalid amount parameter.", http.StatusBadRequest)
		return
	}

	// Add the transaction
	b.NewTransaction(sender[0], receiver[0], float32(_amount))

	// Creating the hash
	h := fmt.Sprintf("Generated new transaction\n")

	// Sending the response
	fmt.Fprintf(w, h)
}

func mine(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod(MINE, req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Get the last block and it's proof
	lastBlock := b.GetLastBlock()
	lastProof := lastBlock.GetProof()

	// Calculate the proof of work
	newProof := b.ProofOfWork(lastProof)

	// Add the transaction meaning that this server got a coin
	b.NewTransaction("0", serverAddress, 1)

	// Get the last block hash
	lastHash := lastBlock.HashBlock()

	// Generate the block
	b.NewBlock(newProof, lastHash)

	// Creating the hash
	h := fmt.Sprintf("Generated new block!\nProof of Work is: %v\n", newProof)

	// Sending the response
	fmt.Fprintf(w, h)
}

func getChain(w http.ResponseWriter, req *http.Request) {
	if !enforceMethod(CHAIN, req) {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Creating the hash
	h := fmt.Sprintf("%v\n", b)

	// Sending the response
	fmt.Fprintf(w, h)
}

func main() {
	fmt.Println("Starting server on port 8090...")

	// Initializing the blockchain
	b = blockchain.InitBlockchain()

	http.HandleFunc(CREATE_TRANSACTION, createTransaction)
	http.HandleFunc(MINE, mine)
	http.HandleFunc(CHAIN, getChain)

	e := http.ListenAndServe(":8090", nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}
