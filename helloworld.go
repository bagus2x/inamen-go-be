package main

import (
	"strconv"
)

type Manusia struct {
	Name interface{}
	No   interface{}
}

func main() {
	a := "-1"
	aInt, _ := strconv.Atoi(a)
	println(uint64(aInt))
}
