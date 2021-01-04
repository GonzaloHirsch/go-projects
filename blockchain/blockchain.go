package blockchain

import (
	"fmt"
	"projects/crypto"
	"reflect"
	"time"
)

// Creation based on https://hackernoon.com/learn-blockchains-by-building-one-117428612f46

///////////////////////////////////////////////////////////////
// Structs
///////////////////////////////////////////////////////////////

type block struct {
	index        uint64
	timestamp    int64
	transactions []transaction
	proof        uint64
	previousHash string
}

type transaction struct {
	sender   string
	receiver string
	amount   float32
}

type Blockchain struct {
	blocks              []block
	currentTransactions []transaction
	currentIndex        uint64
	lastBlock           block
}

///////////////////////////////////////////////////////////////
// Methods
///////////////////////////////////////////////////////////////

var instance Blockchain
var proofTest = []byte{0, 0}

func InitBlockchain() Blockchain {
	instance = Blockchain{
		blocks:              make([]block, 0),
		currentTransactions: make([]transaction, 0),
		currentIndex:        1,
		lastBlock:           block{},
	}
	// Create origin block
	instance.NewBlock(100, "")
	return instance
}

func GetBlockchainInstance() Blockchain {
	return instance
}

/**
Adds a new transaction to the list
*/
func (I *Blockchain) NewTransaction(sender, receiver string, amount float32) {
	// Create the transaction
	t := transaction{
		sender: sender, receiver: receiver, amount: amount,
	}
	// Append it to the current transactions list
	I.currentTransactions = append(I.currentTransactions, t)
}

/**
Adds a new block to the chain
*/
func (I *Blockchain) NewBlock(proof uint64, previousHash string) {
	// In order to be able to create the first block
	if previousHash == "" {
		previousHash = "1"
	}
	// Create the block
	b := block{
		index:        I.currentIndex,
		timestamp:    time.Now().UnixNano(),
		proof:        proof,
		previousHash: previousHash,
		transactions: I.currentTransactions,
	}
	// Append the new block
	I.blocks = append(I.blocks, b)
	// Reset the transactions
	I.currentTransactions = make([]transaction, 0)
	// Update the index
	I.currentIndex += 1
	// Modify the last block variable
	I.lastBlock = b
}

// Calculates the Hash of a given block using the string representation of the hash
func (B *block) HashBlock() string {
	// Convert the string representation to a byte array
	// Calculate the SHA256 of the block
	hash := crypto.Sha256([]byte(B.String()))
	// Return the hex representation of the hash
	return fmt.Sprintf("%x", hash)
}

/**
Function to calculate the proof of work required for the next block
*/
func (I *Blockchain) ProofOfWork(previousProof uint64) uint64 {
	var newProof uint64 = 0
	// Calculate until a good proof can be found
	for !testProof(previousProof, newProof) {
		newProof += 1
	}
	return newProof
}

/**
Getter for the last block of the chain
 */
func (I *Blockchain) GetLastBlock() block{
	return I.lastBlock
}

/**
Getter for the proof of a block
 */
func (B *block) GetProof() uint64{
	return B.proof
}

///////////////////////////////////////////////////////////////
// Stringers
///////////////////////////////////////////////////////////////

/**
String representation of the blockchain
*/
func (I Blockchain) String() string {
	return fmt.Sprint(I.blocks)
}

/**
String representation of a block
*/
func (B block) String() string {
	return fmt.Sprintf("{\n\tindex:%v,\n\ttimestamp:%v,\n\tproof:%v,\n\tpreviousHash:%v,\n\ttransactions:\n\t\t%v\n}", B.index, B.timestamp, B.proof, B.previousHash, B.transactions)
}

/**
String representation of a transaction
*/
func (T transaction) String() string {
	return fmt.Sprintf("{\n\t\t\tsender:%v,\n\t\t\treceiver:%v,\n\t\t\tamount:%v\n\t\t}", T.sender, T.receiver, T.amount)
}


///////////////////////////////////////////////////////////////
// Internal Methods
///////////////////////////////////////////////////////////////

/**
Tests the given proof using the SHA256 representation of PP', where P is the old proof and P' is the new proof
Given proof should have 2 leading 0s, being 00.....
*/
func testProof(previousProof, newProof uint64) bool {
	// Concatenate the previous proof with the new proof
	// Calculate sha256 of the data
	shaResult := crypto.Sha256([]byte(fmt.Sprintf("%v%v", previousProof, newProof)))
	return reflect.DeepEqual(shaResult[:2], proofTest)
}
