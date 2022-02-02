package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/teethew/protobuf/src/simple"
	"google.golang.org/protobuf/proto"
)

func main() {
	
	sm := doSimple()
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Printf("Data has been written to the file %s!\n", fname)
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read from file", err)
		return err
	}

	
	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Can't desserialize to string", err)
		return err
	}


	fmt.Printf("Data has been read from the file %s!\n", fname)
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id: 1,
		IsSimple: true,
		Name: "Message created successfully",
		SampleList: []int32 {7, 0, 0, 4},
	}

	fmt.Println(sm.String())

	sm.Name = "The protocol buffers methods are working fine"

	return &sm
}