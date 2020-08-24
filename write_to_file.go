package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/sunil206b/Protobuff3_Go/src/simple"
)

func writeToFIle(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("something went wrong when retuning the file", err)
		return err
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't put the bytes into the protocoll struct", err)
		return err
	}
	return nil
}

func readAndWriteToFile(pb proto.Message) {
	writeToFIle("simple.bin", pb)
	person1 := &simple.Person{}
	readFromFile("simple.bin", person1)
	fmt.Println("Content from the file:", person1)
}
