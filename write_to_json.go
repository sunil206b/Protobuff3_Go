package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/sunil206b/Protobuff3_Go/src/simple"
)

func writeToJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func readFromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteTOJSON(pb proto.Message) {
	personAsStr := writeToJSON(pb)
	fmt.Println("JSON string:", personAsStr)
	person2 := &simple.Person{}
	readFromJSON(personAsStr, person2)
	fmt.Println("JSON string to proto struct:", person2)
}
