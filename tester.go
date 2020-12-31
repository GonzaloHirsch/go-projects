package main

import "projects/crypto"

func main() {
	crypto.TestHmac()
	crypto.TestSha256()
	crypto.TestMd5()
}
