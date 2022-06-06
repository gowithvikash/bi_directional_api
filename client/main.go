package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/bi_direction_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	do_Greet_Everyone(c)

}

func do_Greet_Everyone(c pb.GreetServiceClient) {
	fmt.Println("\n_______________ Action Number : 01 _______________")
	fmt.Println("_____  do_Greet_Everyone Function Was Invoked At Client  ____")
	stream, err := c.Greet_Everyone(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var reqs = []*pb.GreetRequest{{Name: "Bijender Kumar"}, {Name: "Vikash Parashar"}, {Name: "Khushboo Panday"}, {Name: "Niyati Parashar"}, {Name: "Ritika Parashar"}, {Name: "Rampati Devi"}}

	waitc := make(chan struct{})
	go func() {
		for _, v := range reqs {
			stream.Send(v)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("___ do_Greet_Everyone_Result: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc

}
