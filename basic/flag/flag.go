package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", "name", "say hello to who?")
	flag.Parse()

	fmt.Println(*name)

}
