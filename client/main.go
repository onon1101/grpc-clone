package main

import (
  "context"
  "flag"
  "fmt"
  demo "grpc-demo-clone/proto"
  "log"
  "time"

  "google.golang.org/grpc"
)

func PrintAddResult(client demo.DemoClient, twoNum *demo.TwoNum) {
  fmt.Printf("normal ")
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
  resp, err := client.Add(ctx, twoNum)
  if err != nil {
    log.Fatalf("v.GetFeatrues(_) = _, %v",client, err)
  }
  fmt.Printf("context: %v", resp.Result)
}

var (
  serverAddr = flag.String("server_addr", "localhost:50054", "server format")

)

func main() {
  flag.Parse()

  conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect %v", err)
  }
  defer conn.Close()
  client := demo.NewDemoClient(conn)

  fmt.Printf("First Request...\n")
  PrintAddResult(client, &demo.TwoNum{X: 10, Y: 2})
  fmt.Printf("\n\n")
}
