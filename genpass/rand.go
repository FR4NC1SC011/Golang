package main

import (
	"crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

func init() {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + math_rand.Intn(max-min)
}

/*
func main() {
	password := randomString(12)
	fmt.Println(password)
}
*/
