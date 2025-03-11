package grpcpipe

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/stapelberg/rsync-over-grpc/internal/proto"
	"golang.org/x/sync/errgroup"
)

// Stream is implemented by both, the grpc.BidiStreamingServer and
// grpc.BidiStreamingClient Rsync types.
type Stream interface {
	Context() context.Context
	Recv() (*pb.TransferChunk, error)
	Send(*pb.TransferChunk) error
}

// Adapter returns a managed errgroup which contains goroutines that adapt
// between a gRPC stream and an io.ReadWriter.
func Adapter(stream Stream) (*errgroup.Group, context.Context, io.ReadWriter, error) {
	ctx, cancel := context.WithCancel(stream.Context())
	// cancel is eventually called by the goroutine below

	fromStreamRd, fromStreamWr := io.Pipe()
	toStreamRd, toStreamWr := io.Pipe()
	close := func() {
		log.Printf("closing streams")
		fromStreamWr.Close()
		toStreamWr.Close()
		cancel()
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		for range ctx.Done() {
			// goroutine will spend most of its time here
		}
		log.Printf("context done: %v", ctx.Err())
		close()
		return nil
	})
	eg.Go(func() error {
		var buf [4096]byte
		for {
			// Read from stdoutrd and send RPC requests
			n, err := toStreamRd.Read(buf[:])
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			if err := stream.Send(&pb.TransferChunk{
				Chunk: buf[:n],
			}); err != nil {
				return fmt.Errorf("sending stream chunk: %v", err)
			}
		}
	})
	eg.Go(func() error {
		// Recv() RPC requests and write to stdinwr
		for {
			chunk, err := stream.Recv()
			if err == io.EOF {
				log.Printf("io.EOF from client")
				close()
				return nil
			}
			if err != nil {
				return fmt.Errorf("rsync RPC failed: %v", err)
			}
			if _, err := fromStreamWr.Write(chunk.GetChunk()); err != nil {
				return fmt.Errorf("receiving stream chunk: %v", err)
			}
		}
	})
	rw := &struct {
		io.Reader
		io.Writer
	}{
		Reader: fromStreamRd,
		Writer: toStreamWr,
	}
	return eg, ctx, rw, nil
}
