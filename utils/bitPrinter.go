package utils

import "fmt"

func PrintByte(a byte) {
	for j := 0; j < 8; j++ {
		mask := byte(1 << uint(7-j))
		if a&mask == mask {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}

func PrintInt8(a int8) {
	for j := 0; j < 8; j++ {
		mask := int8(1 << uint(7-j))
		if a&mask == mask {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}

func PrintInt16(a int16) {
	for j := 0; j < 16; j++ {
		mask := int16(1 << uint(15-j))
		if a&mask == mask {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}

func PrintInt32(a int32) {
	for j := 0; j < 32; j++ {
		mask := int32(1 << uint(31-j))
		if a&mask == mask {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}

func PrintInt64(a int64) {
	for j := 0; j < 64; j++ {
		mask := int64(1 << uint(63-j))
		if a&mask == mask {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
