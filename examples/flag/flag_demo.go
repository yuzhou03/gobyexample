package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "File contains values from sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort Algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile:", *infile, "outfile:", *outfile, "algorithm", *algorithm)
	}
}
