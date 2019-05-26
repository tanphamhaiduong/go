package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/handler"
	"github.com/tanphamhaiduong/go/delta/server/modules"
)

func main() {
	db := initDBConnection()
	resolvers := modules.NewResolver(db)
	schema, err := modules.MakeSchema(resolvers)
	if err != nil {
		log.Fatal(err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Listening on localhost:8080")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
