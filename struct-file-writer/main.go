package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

//Goal is to be able to write custom structs to file.
//The data should be written in a way , that we can read the data back to a struct.
//The writer will write multiple rows of "Person" struct into a file.
//The Struct-File-Reader file will contain the logic to read the file into structs.

func main() {

	data := make([]Person, 10)
	for i := 0; i < 10; i++ {
		data[i] = Person{
			Name: "Name " + strconv.Itoa(i),
			Age:  i,
		}
	}

	//Write all these 10 objects into the file.

	f, err := os.Create("../Structs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	writer := bufio.NewWriter(f)

	defer func (){
		f.Close()
	}()

	for _, d := range data {
		//Convert every data to a byte slice using json marshalling
		marshalledBytes, err := json.Marshal(d)
		if err != nil {
			log.Fatalln(err)
		}
		bytesWritten, err := writer.Write(marshalledBytes)
		if err != nil {
			log.Fatalln(err)
		}
		writer.WriteString("\n")
		fmt.Println("bytes written are ", bytesWritten)
	}
	//Need to flush the buffer.
	writer.Flush()

}
