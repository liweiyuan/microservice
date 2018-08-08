package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"microservice/transport"
	"net/http"
	"log"
	"fmt"
	"microservice/service"
)

func main() {
	svc := service.StringServiceImpl{}
	uppercaseHandler := httptransport.NewServer(
		transport.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)
	countHandler := httptransport.NewServer(
		transport.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)
	fmt.Printf("listener to the port : 8080")

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


