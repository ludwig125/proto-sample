package main

import (
	"io"
	"log"
	"net/http"

	pb "myproto/examples/go/tutorialpb"

	"github.com/golang/protobuf/proto"
)

func personHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var req pb.Person
	if err := proto.Unmarshal(b, &req); err != nil {
		http.Error(w, "Failed to unmarshal request", http.StatusBadRequest)
		return
	}
	log.Printf("Received value: %v", req.Name)
	// res := &pb.BoolMessage{Value: req.Value}
	// w.Header().Set("Content-Type", "application/x-protobuf")
	// w.WriteHeader(http.StatusOK)
	// if err := proto.MarshalResponse(w, res); err != nil {
	// 	http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
	// 	return
	// }
}

func main() {
	http.HandleFunc("/person", personHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
