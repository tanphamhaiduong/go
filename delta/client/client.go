package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/tanphamhaiduong/go/delta/pb"
	"google.golang.org/grpc"
)

const port = ":9000"

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()
	// Remove cert
	// creds, err := credentials.NewClientTLSFromFile("./cert.pem", "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewDeltaServiceClient(conn)
	switch *option {
	case 1:
		CompanyList(client)
		return
	case 2:
		CompanyCreate(client)
		return
	}

}

// CompanyList ...
func CompanyList(client pb.DeltaServiceClient) {
	res, err := client.CompanyList(context.Background(), &pb.CompanyListRequest{
		PageSize: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Company)
}

// CompanyCreate ...
func CompanyCreate(client pb.DeltaServiceClient) {
	res, err := client.CompanyInsert(context.Background(), &pb.CompanyInsertRequest{
		Name:   "Test company card",
		Status: "active",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Company)
}
