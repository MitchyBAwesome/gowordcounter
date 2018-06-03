package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile(os.Args[1]) // Read in the file as specified in the command 1 argument

	if err != nil {
		panic(err)
	}

	res := (strings.Split(string(data), " ")) // Convert the []byte in to a string and split the string at each 'space'.

	fmt.Printf("There are %v words in your document", len(res)) // Output the number of words.
}
