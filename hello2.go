package main

import (
	"fmt"
	"github.com/xtothp00/hello2/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf("%s\n", morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World!", "Hello Go!"))
}
