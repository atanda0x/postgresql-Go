package main

import (
	"log"

	base62 "github.com/atanda0x/postgresql-Go/base62"
)

// Func base62 from the other folder  was used
func main() {
	x := 10000
	base64String := base62.ToBase62(x)
	log.Println(base64String)
	normalNumber := base62.ToBase10(base64String)
	log.Println(normalNumber)
}
