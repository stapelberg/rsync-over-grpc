# rsync-over-grpc

This is a demo program to show how to use the
[gokrazy/rsync](https://github.com/gokrazy/rsync) module over a gRPC transport.

The [`rsyncovergrpc.proto`
file](https://github.com/stapelberg/rsync-over-grpc/blob/main/internal/proto/rsyncovergrpc.proto)
shows the required structure to embed the rsync protocol: you need an initial
request which sends the command-line arguments from client to server (=
configures the transfer), then a bi-directional pipe between client and server
to speak the rsync protocol over.

## Usage

```
% go install github.com/stapelberg/rsync-over-grpc/cmd/...@latest
```

Then, open two terminal windows. In the first one, run:

```
% grpc-rsync-server
2025/03/11 20:21:08 server.go:65: server listening at [::]:50051
```

In the second one, start the client to start the transfer of
`/usr/share/man/man7` to `/tmp/dest`:

```
% grpc-rsync-client
2025/03/11 20:31:44 clientmaincmd.go:231: remote protocol: 27
2025/03/11 20:31:44 clientmaincmd.go:292: receiving to dest=/tmp/dest
[…]
2025/03/11 20:31:52 do.go:140: server sent stats: read=16, written=41270, size=3858067
2025/03/11 20:31:52 client.go:58: client done, err: <nil>
```

In the first one, the log output will read something like:

```
[…]
2025/03/11 20:31:32 server.go:35: handling Rsync request args:"--server" args:"--sender" args:"-logDtpr" args:"." args:"/usr/share/man/man7"
2025/03/11 20:31:32 rsyncd.go:475: exclusion list read (entries: 0)
2025/03/11 20:31:32 do.go:45: SendFileList(modPath="/", paths=["/usr/share/man/man7"])
2025/03/11 20:31:32 rsyncd.go:482: handleConnSender done. stats: &{Read:20 Written:41286 Size:3858067}
2025/03/11 20:31:32 grpcpipe.go:67: io.EOF from client
2025/03/11 20:31:32 grpcpipe.go:30: closing streams
2025/03/11 20:31:32 grpcpipe.go:40: context done: context canceled
2025/03/11 20:31:32 grpcpipe.go:30: closing streams
```
