package main

import (
	"fmt"
	"operate/core"
	"testing"
)

func TestMain01(t *testing.T) {
	err := core.OutFile("", "./", 0)
	fmt.Println(err)
}
