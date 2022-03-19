package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

//Reads the file line by line and unmarshalls the data into a struct.
func main() {

	f, err := os.Open("Structs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	s := make([]Person, 10)
	//Read data from the file , it becomes important to read data line by line because we want to parse these into structs.
	sc := bufio.NewScanner(f)
	index := 0
	for sc.Scan() {
		p := Person{}
		lineBytes := sc.Bytes()
		err := json.Unmarshal(lineBytes, &p)
		if err != nil {
			log.Fatalln(err)
		}
		s[index] = p
		index++
	}

	fmt.Println(s)
}
