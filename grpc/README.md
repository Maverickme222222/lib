# How to Use the gRPC Interceptors



### Server-side Interceptors

```go

import (
interceptors "github.com/kappapay/backend/lib/golang/src/kappa/grpc"
)

grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.ServerInterceptorsWithTracing(logger)),
	)
	reflection.Register(grpcServer)


```