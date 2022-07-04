package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	prv, err := NewProvider("client")
	if err != nil {
		log.Fatalln(err)
	}
	defer prv.Close(ctx)

	rtr := http.DefaultServeMux
	rtr.HandleFunc("/api/v1/users", HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, span := NewSpan(r.Context(), "Controller.Create", nil)
		defer span.End()

		fmt.Println(span.SpanContext().TraceID().String())
		req, _ := http.NewRequest("GET", "http://localhost:8081/api/v1/users", nil)
		InjectHeaders(ctx, req)

		cli := http.Client{}
		cli.Do(req)

	}, "users_create"))

	if err := http.ListenAndServe(":8080", rtr); err != nil {
		log.Fatalln(err)
	}
}
