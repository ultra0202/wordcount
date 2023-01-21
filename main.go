package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	e := strings.Fields(str)
	fmt.Println(len(e))
}
