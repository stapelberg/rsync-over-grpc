package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/gokrazy/rsync/rsyncd"
	"github.com/stapelberg/rsync-over-grpc/internal/grpcpipe"
	pb "github.com/stapelberg/rsync-over-grpc/internal/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedRsyncServer
}

func (s *server) Rsync(stream grpc.BidiStreamingServer[pb.TransferChunk, pb.TransferChunk]) error {
	// Wait for the first message from the client, which initiates the stream
	first, err := stream.Recv()
	if err == io.EOF {
		return fmt.Errorf("stream closed before any data was received?!")
	}
	if first.Request == nil {
		return fmt.Errorf("first message must set request field")
	}
	if len(first.GetChunk()) > 0 {
		return fmt.Errorf("first message must not contain any data chunks")
	}

	log.Printf("handling Rsync request %+v", first.Request)

	eg, ctx, rw, err := grpcpipe.Adapter(stream)
	if err != nil {
		return err
	}
	rsync, err := rsyncd.NewServer(nil)
	if err != nil {
		return err
	}
	eg.Go(func() error {
		args := first.GetRequest().GetArgs()
		conn := rsyncd.NewConnection(rw, rw, "<gRPC>")
		return rsync.HandleConnArgs(ctx, conn, nil, args)
	})
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRsyncServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
