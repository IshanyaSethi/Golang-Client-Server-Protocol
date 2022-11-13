// Create a GRPC server which creates a firebase analytics custom event

package firebase

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "firebase"
)

type server struct{}

func (s *server) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Event, error) {
	log.Printf("Received: %v", in)
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFirebaseAnalyticsServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
