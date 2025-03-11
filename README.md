# rsync-over-grpc

This is a demo program to show how to use the
[gokrazy/rsync](https://github.com/gokrazy/rsync) module over a gRPC transport.

The `rsyncovergrpc.proto` file shows the required structure to embed the rsync
protocol: you need an initial request which sends the command-line arguments
from client to server (= configures the transfer), then a bi-directional pipe
between client and server to speak the rsync protocol over.
