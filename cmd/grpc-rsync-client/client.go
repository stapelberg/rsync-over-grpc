package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/gokrazy/rsync/rsyncclient"
	"github.com/stapelberg/rsync-over-grpc/internal/grpcpipe"
	pb "github.com/stapelberg/rsync-over-grpc/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func demoTransfer() error {
	flag.Parse()

	args, src, dest := []string{"-av"}, "/usr/share/man/man7", "/tmp/dest"
	client, err := rsyncclient.New(args)
	if err != nil {
		return err
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRsyncClient(conn)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	stream, err := c.Rsync(ctx)
	if err != nil {
		return fmt.Errorf("could not create stream: %v", err)
	}
	if err := stream.Send(&pb.TransferChunk{
		Request: &pb.TransferRequest{
			Args: client.ServerCommandOptions(src),
		},
	}); err != nil {
		return fmt.Errorf("starting stream: %v", err)
	}

	eg, ctx, rw, err := grpcpipe.Adapter(stream)
	if err != nil {
		return err
	}
	eg.Go(func() error {
		defer stream.CloseSend()
		_, err := client.Run(ctx, rw, []string{dest})
		log.Printf("client done, err: %v", err)
		return err
	})
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := demoTransfer(); err != nil {
		log.Fatal(err)
	}
}
