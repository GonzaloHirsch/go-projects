package crypto

import "encoding/binary"

// Right rotation function
func RotL(u uint32, times int) uint32 {
	return (u << times) | (u >> (32 - times))
}

// Right rotation function
func RotR(u uint32, times int) uint32 {
	return (u >> times) | (u << (32 - times))
}

// Right shift function
func ShiftR(u uint32, times int) uint32 {
	return u >> times
}

// Left shift function
func ShiftL(u uint32, times int) uint32 {
	return u << times
}

// Converts a byte array into chunks of the same size
func ChunkerizeSlice(s []byte, chunkSize int) [][]byte {
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
func HashToByteSlice(hash []uint32) []byte {
	// Allocating the space
	arr := make([]byte, 0, len(hash)*4)

	// Appending the bytes
	for _, h := range hash {
		temp := make([]byte, 4)
		binary.BigEndian.PutUint32(temp, h)
		arr = append(arr, temp...)
	}

	return arr
}
