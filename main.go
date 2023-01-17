package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var str string
	flag.String(str, "", "****")
	flag.Parse()
	src := flag.Args()
	//fmt.Printf("%T %v %d\n", src, src, len(src))
	//fmt.Println("**", src[0], "**")
	e := strings.Fields(src[0])
	fmt.Println(len(e))
}
