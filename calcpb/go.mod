module github.com/fuskovic/grpcDemo/calculator

go 1.15

require (
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/fuskovic/grpcDemo/calculator/calcpb => ./