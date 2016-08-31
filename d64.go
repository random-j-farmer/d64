package d64

import "fmt"

// ALPHABET is used by the d64 encoding, in this order
const ALPHABET = ".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

// invertedAlphabet gives the value of a digit. -1 if not in the alphabet
var invertedAlphabet [256]int8

func init() {
	for i := 0; i < len(invertedAlphabet); i++ {
		invertedAlphabet[i] = -1
	}
	for i := 0; i < len(ALPHABET); i++ {
		invertedAlphabet[ALPHABET[i]] = int8(i)
	}
}

// EncodeUInt64 encodes the number n, zero-padding on the left so minWidth is reached.
//
// Encoding negative numbers looses the nice sort order properties of d64,
// that's why there is only a function for unsigned
//
func EncodeUInt64(n uint64, minWidth int) string {
	if minWidth == 0 {
		minWidth = 1
	}
	buff := make([]byte, minWidth)
	for i := 0; i < len(buff); i++ {
		buff[i] = ALPHABET[0]
	}

	i := 0
	for {
		if i == len(buff) {
			buff = append(buff, '.')
		}

		digit := n & 0x3F
		buff[i] = ALPHABET[digit]
		n = n >> 6
		if n == 0 {
			return string(reverse(buff))
		}
		i++
	}
}

// DecodeUInt64 decodes a number from a string.
//
// If the string contains anything not in the ALPHABET,
// an error is returned.
func DecodeUInt64(s string) (uint64, error) {
	var n uint64

	for i := 0; i < len(s); i++ {
		v := invertedAlphabet[s[i]]
		if v < 0 {
			return 0, fmt.Errorf("invalid d64 digit %#x(%c) at byte %d in: %s", s[i], byte(s[i]), i, s)
		}
		if v := invertedAlphabet[s[i]]; v >= 0 {
			n = (n << 6) + uint64(v)
		}
	}

	return n, nil
}

// EncodeBytes encodes a blob of bytes.
func EncodeBytes(src []byte) []byte {
	srcLen := len(src)
	dst := make([]byte, (srcLen/3+1)*4)

	var hang byte
	var dstPos int
	for i := 0; i < srcLen; i++ {
		v := src[i]
		switch i % 3 {
		case 0:
			dst[dstPos] = ALPHABET[v>>2]
			dstPos++
			hang = (v & 0x3) << 4
		case 1:
			dst[dstPos] = ALPHABET[hang|v>>4]
			dstPos++
			hang = (v & 0xF) << 2
		case 2:
			dst[dstPos] = ALPHABET[hang|v>>6]
			dstPos++
			dst[dstPos] = ALPHABET[v&0x3F]
			dstPos++
			hang = 0
		}
	}

	if (srcLen % 3) > 0 {
		dst[dstPos] = ALPHABET[hang]
		dstPos++
	}

	return dst[:dstPos]
}

// DecodeBytes decodes a d64 string into a blob of bytes
//
// It the input contains bytes not in the ALPHABET, an
// error is returned.
func DecodeBytes(src []byte) ([]byte, error) {
	srcLen := len(src)
	dst := make([]byte, (srcLen/4+1)*3)

	var hang byte
	var dstPos int
	for i := 0; i < srcLen; i++ {
		vint8 := invertedAlphabet[src[i]]
		if vint8 < 0 {
			return nil, fmt.Errorf("invalid d64 digit %#x(%c) at byte %d in: %s", src[i], src[i], i, src)
		}
		v := byte(vint8)

		switch i % 4 {
		case 0:
			hang = v << 2
		case 1:
			dst[dstPos] = hang | v>>4
			dstPos++
			hang = byte(v << 4)
		case 2:
			dst[dstPos] = hang | v>>2
			dstPos++
			hang = v << 6
		case 3:
			dst[dstPos] = hang | v
			dstPos++
		}
	}

	return dst[:dstPos], nil
}

func reverse(b []byte) []byte {
	length := len(b)
	lm1 := length - 1
	for i := 0; i < length/2; i++ {
		b[i], b[lm1-i] = b[lm1-i], b[i]
	}
	return b
}
