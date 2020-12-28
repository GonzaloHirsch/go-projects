package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"reflect"
)

func pad(b []byte, target int) []byte {
	missingBytes := target - len(b)
	filler := make([]byte, missingBytes)
	b = append(b, filler...)
	return b
}

func HmacSha256(key, msg []byte) [SHA_256_OTUPUT_SIZE]byte {
	// Resizing the key to match the block size
	if len(key) > SHA_256_CHUNK_SIZE_BYTES {
		temp := Sha256(key)
		key = temp[:]
	}

	if len(key) < SHA_256_CHUNK_SIZE_BYTES {
		key = pad(key, SHA_256_CHUNK_SIZE_BYTES)
	}

	// Getting inner and outer keys
	oKey, iKey := make([]byte, len(key)), make([]byte, len(key))
	copy(oKey, key)
	copy(iKey, key)
	for i := range key {
		oKey[i] = oKey[i] ^ 0x5c
		iKey[i] = iKey[i] ^ 0x36
	}

	// Getting the first pass
	_innerHash := Sha256(append(iKey, msg...))
	innerHash := _innerHash[:]

	// Calculating the second pass
	return Sha256(append(oKey, innerHash...))
}

func TestHmac() {
	key1, msg1 := []byte("key"), []byte("The quick brown fox jumps over the lazy dog")
	_res1 := HmacSha256(key1, msg1)
	res1 := _res1[:]
	h1 := hmac.New(sha256.New, key1)
	h1.Write(msg1)
	sha1 := h1.Sum(nil)
	fmt.Printf("HMAC-SHA256 sum for %v with key %v results are: \n\tMINE: %x\n\tREAL: %x\nRESULTS ARE EQUAL?: %v\n", string(msg1), string(key1), res1, sha1, reflect.DeepEqual(res1, sha1))

	key2, msg2 := []byte("keykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykeykey"), []byte("abc")
	_res2 := HmacSha256(key2, msg2)
	res2 := _res2[:]
	h2 := hmac.New(sha256.New, key2)
	h2.Write(msg2)
	sha2 := h2.Sum(nil)
	fmt.Printf("HMAC-SHA256 sum for %v with key %v results are: \n\tMINE: %x\n\tREAL: %x\nRESULTS ARE EQUAL?: %v\n", string(msg2), string(key2), res2, sha2, reflect.DeepEqual(res2, sha2))
}

