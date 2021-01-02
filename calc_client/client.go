package main

import (
    "context"
    "fmt"
    "log"
	"calculator/calcpb"
	
    "google.golang.org/grpc"
)

func main(){
    //grpc.WithInsecure() means our connection is not encrypted.
    //we'll cover securing our gRPC connections with SSL/TLS in a separate post.
    conn,err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil{
        log.Fatalf("error : %v\n", err)
    }
    defer conn.Close()

    c := calcpb.NewCalculateClient(conn)
	fmt.Println("Client successfully initialized")

    //new code
    unaryCall(c)
    }

    //new code
    func unaryCall(c calcpb.CalculateClient){
        req := &calcpb.SumRequest{
            NumOne : 6,
            NumTwo: 7,
        }
        res, err := c.Sum(context.Background(), req)
        if err != nil{
            log.Fatalf("error calling sum endpoint : %v\n", err)
        }
        log.Printf("Sum response from server: %v\n", res.Sum)
    }

}
