package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	for i, arg := range os.Args[0:] {
		s := strconv.Itoa(i)+" : "+arg
		fmt.Println(s)
	}
}
