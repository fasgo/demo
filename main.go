package main

import "github.com/fasgo/protoapi"

func main() {
	//
	s := protoapi.NewServer()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
