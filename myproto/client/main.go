package main

import (
	"bytes"
	"log"
	"net/http"

	pb "myproto/examples/go/tutorialpb"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := &pb.Person{Name: "abc"}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}
	resp, err := http.Post("http://localhost:8080/person", "application/x-protobuf", bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
	// res := &pb.BoolMessage{}
	// err = proto.UnmarshalResponse(resp.Body, res)
	// if err != nil {
	// 	log.Fatalf("Failed to unmarshal response: %v", err)
	// }
	// log.Printf("Received bool value: %v", res.Value)
}
