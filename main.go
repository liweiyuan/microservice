package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"microservice/transport"
	"net/http"
	"fmt"
	"microservice/service"
)

func main() {

	//log
	//logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.StringService
	svc = service.StringServiceImpl{} //接口与实现类的调用

	//svc = loggingMiddleware{logger, svc}

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
	http.ListenAndServe(":8080", nil)

	//日志打印
	//logger.Log("msg", "HTTP", "addr", "8080")
	//logger.Log("err", http.ListenAndServe("8080", nil))

}
