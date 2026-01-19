package handler

import (
	"context"

	log "go-micro.dev/v5/logger"

	pb "greeter/proto"
)

type Greeter struct{}

// Return a new handler
func New() *Greeter {
	return &Greeter{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Greeter) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("Received Greeter.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Greeter) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.Greeter_StreamStream) error {
	log.Infof("Received Greeter.Stream request with count: %d", req.Count)

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
