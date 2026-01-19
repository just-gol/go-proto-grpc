package handler

import (
	"context"

	log "go-micro.dev/v5/logger"

	pb "service/proto"
)

type Service struct{}

// Return a new handler
func New() *Service {
	return &Service{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Service) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("Received Service.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Service) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.Service_StreamStream) error {
	log.Infof("Received Service.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}
