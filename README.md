# Go: Sample Code for Using Socket Options

This repository has two TCP servers: `./cmd/server1` and `./cmd/server2`.
Even though they listen on same port and host, they can be served since `SO_REUSEADDR` socket option is enabled.

See: https://github.com/110y/sockoptgo/blob/master/internl/server/server.go#L81
