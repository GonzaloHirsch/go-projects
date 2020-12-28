package crypto

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

// Struct to be able to reset values more easily
type digest struct {
	H [8]uint32
}

// Initializing round constants
var K = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

// Initializing hash values
var H = [8]uint32{
	0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
}

const SHA_256_CHUNK_SIZE_BITS = 512
const SHA_256_CHUNK_SIZE_BYTES = 64
const SHA_256_FILL_LIMIT = 16
const SHA_256_SCHEDULE_SIZE = 64
const SHA_256_OTUPUT_SIZE = 32

// Right rotation function
func RotR(u uint32, times int) uint32 {
	return (u >> times) | (u << (32 - times))
}

// Right shift function
func ShiftR(u uint32, times int) uint32 {
	return u >> times
}

// Converts a byte array into chunks of the same size
func chunkerizeSlice(s []byte, chunkSize int) [][]byte {
	// Calculating number of chunks
	n := int(len(s) / chunkSize)
	// Allocating space for the chunks needed
	var chunks = make([][]byte, n)
	// Putting each chunk in place
	for i := 0; i < n; i++ {
		chunks[i] = s[(chunkSize * i):(chunkSize * (i + 1))]
	}
	return chunks
}

// Converts the resulting hash of 32 bit words into a byte array
func hashToByteArray(hash []uint32) [SHA_256_OTUPUT_SIZE]byte {
	// Allocating the space
	arr := make([]byte, 0, len(hash)*4)

	// Appending the bytes
	for _, h := range hash {
		temp := make([]byte, 4)
		binary.BigEndian.PutUint32(temp, h)
		arr = append(arr, temp...)
	}

	var _arr [SHA_256_OTUPUT_SIZE]byte
	copy(_arr[:], arr)

	return _arr
}

// Resets the H values of the digest struct
func (dig *digest) reset() {
	dig.H = H
}

// Calculates the 256 checksum of the data
func (dig *digest) checkSum256(_s []byte) [SHA_256_OTUPUT_SIZE]byte {
	// Preprocessing the message

	// Length of the original message
	l := len(_s)

	// Transform the original array to a slice
	s := _s[:]

	// Append a '1' bit to the message, it is a 1000|0000 -> 0x80 (1 bit with trailing 0s)
	s = append(s, 0x80)

	// Calculate k given L + 1 + K + 64 is multiple of 512
	// Multiply l by 8 to have it in bits
	k := float64(512 - (((l * 8) + 8 + 64) % 512))

	// Calculate number of bytes to append
	kb := int(math.Ceil(k / 8.0))

	// Creating array to append
	kbs := make([]byte, kb)

	// Append the k 0 bits
	s = append(s, kbs...)

	// Converting the length to a big endian 64 bit number
	_l := uint64(l * 8)
	lbs := make([]byte, 8)
	binary.BigEndian.PutUint64(lbs, _l)

	// Append the l bits
	s = append(s, lbs...)

	// Transforming into chunks
	chunks := chunkerizeSlice(s, SHA_256_CHUNK_SIZE_BYTES)

	// Rounds

	// Iterate each chunk
	for _, chunk := range chunks {
		// Creating the 64 entry 32-bit word array
		w := make([]uint32, SHA_256_SCHEDULE_SIZE)

		// Filling the first 16 positions with the chunk
		for i := 0; i < SHA_256_FILL_LIMIT; i++ {
			w[i] = binary.BigEndian.Uint32(chunk[(i * 4):((i + 1) * 4)])
		}

		// Extending the original words
		for i := 16; i < SHA_256_SCHEDULE_SIZE; i++ {
			s0 := (RotR(w[i-15], 7)) ^ (RotR(w[i-15], 18)) ^ (ShiftR(w[i-15], 3))
			s1 := (RotR(w[i-2], 17)) ^ (RotR(w[i-2], 19)) ^ (ShiftR(w[i-2], 10))
			w[i] = w[i-16] + s0 + w[i-7] + s1
		}

		// Initializing the variables
		a, b, c, d, e, f, g, h := dig.H[0], dig.H[1], dig.H[2], dig.H[3], dig.H[4], dig.H[5], dig.H[6], dig.H[7]

		// Compression loop
		for i := 0; i < SHA_256_SCHEDULE_SIZE; i++ {
			s1 := (RotR(e, 6)) ^ (RotR(e, 11)) ^ (RotR(e, 25))
			ch := (e & f) ^ ((^e) & g)
			temp1 := h + s1 + ch + K[i] + w[i]
			s0 := (RotR(a, 2)) ^ (RotR(a, 13)) ^ (RotR(a, 22))
			maj := (a & b) ^ (a & c) ^ (b & c)
			temp2 := s0 + maj

			h, g, f, e, d, c, b, a = g, f, e, (d+temp1)>>0, c, b, a, (temp1+temp2)>>0
		}

		// Add compressed to current hash value
		dig.H[0], dig.H[1], dig.H[2], dig.H[3], dig.H[4], dig.H[5], dig.H[6], dig.H[7] = (dig.H[0]+a)>>0, (dig.H[1]+b)>>0, (dig.H[2]+c)>>0, (dig.H[3]+d)>>0, (dig.H[4]+e)>>0, (dig.H[5]+f)>>0, (dig.H[6]+g)>>0, (dig.H[7]+h)>>0
	}

	return hashToByteArray(dig.H[:])
}

// Function to calculate the Sha256 of a given byte array message
func Sha256(s []byte) [SHA_256_OTUPUT_SIZE]byte {
	d := digest{}
	d.reset()
	return d.checkSum256(s)
}

func TestSha256() {
	s1, s2, s3 := "abcde", "abc", ""
	b1 := Sha256([]byte(s1))
	b2 := Sha256([]byte(s2))
	b3 := Sha256([]byte(s3))
	sum1 := sha256.Sum256([]byte(s1))
	sum2 := sha256.Sum256([]byte(s2))
	sum3 := sha256.Sum256([]byte(s3))
	fmt.Printf("SHA256 sum for %v results are: \n\tMINE: %x\n\tREAL: %x\nRESULTS ARE EQUAL?: %v\n", s1, b1, sum1, sum1 == b1)
	fmt.Printf("SHA256 sum for %v results are: \n\tMINE: %x\n\tREAL: %x\nRESULTS ARE EQUAL?: %v\n", s2, b2, sum2, sum2 == b2)
	fmt.Printf("SHA256 sum for %v results are: \n\tMINE: %x\n\tREAL: %x\nRESULTS ARE EQUAL?: %v\n", s3, b3, sum3, sum3 == b3)
}
