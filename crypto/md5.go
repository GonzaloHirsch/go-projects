package crypto

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math"
)

const MD5_CHUNK_SIZE_BITS = 512
const MD5_CHUNK_SIZE_BYTES = 64
const MD5_FILL_LIMIT = 16
const MD5_SCHEDULE_SIZE = 16
const MD5_OTUPUT_SIZE = 16 // 128-bit output

// Struct to be able to reset values more easily
type md5Digest struct {
	VARS [4]uint32
}

// Initializing round constants
var MD5_K = [64]uint32{
	0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee, 0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
	0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be, 0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
	0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa, 0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
	0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed, 0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
	0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c, 0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
	0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05, 0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
	0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039, 0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
	0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1, 0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
}

// Initializing the shift amounts
var S = [64]int{
	7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
	5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
	4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
	6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
}

// Initial value for the ABCD variables
var VARS = [4]uint32{
	0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476,
}

// Resets the H values of the md5Digest struct
func (dig *md5Digest) reset() {
	dig.VARS = VARS
}

// Pads the given message bytes
func padMessage(_s []byte) []byte {
	// Preprocessing the message

	// Length of the original message
	l := len(_s)

	// Transform the original array to a slice
	s := _s[:]

	// Append a '1' bit to the message, it is a 1000|0000 -> 0x80 (1 bit with trailing 0s)
	s = append(s, 0x80)

	// Calculate k given L + 1 + K + 64 is multiple of 512
	// Multiply l by 8 to have it in bits
	k := float64(MD5_CHUNK_SIZE_BITS - (((l * 8) + 8 + 64) % MD5_CHUNK_SIZE_BITS))

	// Calculate number of bytes to append
	kb := int(math.Ceil(k / 8.0))

	// Creating array to append
	kbs := make([]byte, kb)

	// Append the k 0 bits
	s = append(s, kbs...)

	// Converting the length to a big endian 64 bit number
	_l := uint64(l * 8)
	lbs := make([]byte, 8)
	binary.LittleEndian.PutUint64(lbs, _l)

	// Append the l bits
	s = append(s, lbs...)

	return s
}

// Calculates the md5 checksum of the data
// Algorithm taken from https://en.wikipedia.org/wiki/MD5
func (dig *md5Digest) checkSumMd5(_s []byte) [MD5_OTUPUT_SIZE]byte {
	// Padding the original message
	s := padMessage(_s)

	// Transforming into chunks
	chunks := ChunkerizeSlice(s, MD5_CHUNK_SIZE_BYTES)

	// Rounds

	// Iterate each chunk
	for _, chunk := range chunks {
		// Creating the 16 entry 32-bit word array
		w := make([]uint32, MD5_SCHEDULE_SIZE)

		// Filling the 16 positions with the chunk
		for i := 0; i < MD5_FILL_LIMIT; i++ {
			w[i] = binary.LittleEndian.Uint32(chunk[(i * 4):((i + 1) * 4)])
		}

		// Initializing the variables
		a, b, c, d := dig.VARS[0], dig.VARS[1], dig.VARS[2], dig.VARS[3]

		// Compression loop
		for i := 0; i < MD5_CHUNK_SIZE_BYTES; i++ {
			var f, g uint32
			if 0 <= i && i <= 15 {
				f = (b & c) | ((^b) & d)
				g = uint32(i)
			} else if 16 <= i && i <= 31 {
				f = (d & b) | ((^d) & c)
				g = (5*uint32(i) + 1) & 0x0F
			} else if 32 <= i && i <= 47 {
				f = b ^ c ^ d
				g = (3*uint32(i) + 5) & 0x0F
			} else if 48 <= i && i <= 63 {
				f = c ^ (b | (^d))
				g = (7 * uint32(i)) & 0x0F
			}
			temp := d
			d, c, b, a = c, b, (b+RotL(f+a+MD5_K[i]+w[g], S[i]))>>0, temp
		}

		// Add compressed to current hash value
		dig.VARS[0], dig.VARS[1], dig.VARS[2], dig.VARS[3] = (dig.VARS[0]+a)>>0, (dig.VARS[1]+b)>>0, (dig.VARS[2]+c)>>0, (dig.VARS[3]+d)>>0
	}

	var _arr [16]byte

	binary.Write(bytes.NewBuffer(_arr[:0]), binary.LittleEndian, dig.VARS[:])

	//copy(_arr[:], HashToByteSlice(dig.VARS[:]))

	return _arr
}

// Function to calculate the Sha256 of a given byte array message
func Md5(s []byte) [MD5_OTUPUT_SIZE]byte {
	d := md5Digest{}
	d.reset()
	return d.checkSumMd5(s)
}

func TestMd5() {
	var cases = []string{
		"The quick brown fox jumps over the lazy dog", "The quick brown fox jumps over the lazy dog.", "", "abc",
	}

	for _, c := range cases {
		b := Md5([]byte(c))
		s := md5.Sum([]byte(c))
		fmt.Printf("MD5 sum for %v results are: \n\tMINE: %x\n\tREAL: %x\nRESULTS ARE EQUAL?: %v\n", c, b, s, s == b)
	}
}
