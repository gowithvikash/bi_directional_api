package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/bi_direction_api/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
func (s *Server) Greet_Everyone(stream pb.GreetService_Greet_EveryoneServer) error {
	fmt.Println("___ Greet_Everyone Function Was Invoked At Client___")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		err = stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello And Welcome , %s !\n", req.Name)})
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil

}
