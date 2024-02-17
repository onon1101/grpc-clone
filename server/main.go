package main

import (
  "context"
  "flag"
  "fmt"
  //"io"
  "log"
  "net"

  demo "grpc-demo-clone/proto"

  // "google.golang.org/grpc/credentials"
  "google.golang.org/grpc/reflection"

  "google.golang.org/grpc"
)

type demoServer struct {
  demo.UnimplementedDemoServer
  saveResults []*demo.Response
}

var (
  port = flag.Int("port", 50054, "server port")
)

func (s *demoServer) Add(ctx context.Context, in *demo.TwoNum) (*demo.Response, error) {
  x := in.X
  y := in.Y

  cRes := &demo.Response{
    Result: x+y,
  }

  return cRes, nil
}

func main() {
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
    log.Fatalln(err)
  }

  s := grpc.NewServer()
  demo.RegisterDemoServer(s, &demoServer{})
  reflection.Register(s)
  log.Printf("Server listeing at : %v\n", *port)
  if err:= s.Serve(lis); err != nil {
    log.Fatalln(err)
  }
}
