package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var p1 = "$2a$10$4l1rFD3gAYCBU6IDKEnhPuRXZegZTtqMg8.5AlJLoBSScl4vyTJsy"
var p2 = "$2a$10$YtQ0zYhMEcNJMuM.hBYpmOw7.L5GXVuyRM3n3uBVONqwyxA3nW4P6"

func main() {
	e := p2
	p := "Aa123098.."
	//if hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost); err != nil {
	//	return
	//} else {
	//	e = string(hash)
	//}

	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		panic(err)
	}
	fmt.Println(e, p)
}
