package main


import (
    "context"
    "fmt"
    "net"
    "log"

    "calculator/calcpb"

    "google.golang.org/grpc"
)

type server struct{}

func main(){
    //50051 is the default gRPC port
    lis, err := net.Listen("tcp", "0.0.0.0:50051")
    if err != nil{
        log.Fatalf("failed to listen on 50051 : %v\n", err)
    }

    s := grpc.NewServer()
    calcpb.RegisterCalculateServer(s, &server{})

    fmt.Println("server starting on port : 50051")
    if err := s.Serve(lis); err != nil{
        log.Fatalf("failed to start server : %v\n", err)
    }
}

type server struct{}

//new code
func (s *server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error){
        result := req.GetNumOne() + req.GetNumTwo()
        return &calcpb.SumResponse{
            Sum : result,
        }, nil
    }