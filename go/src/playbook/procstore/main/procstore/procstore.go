package main

import (
	"flag"
	"fmt"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pbabi "playbook/apis/procstore"
	"playbook/procstore"
)

var host = flag.String("db-host", "127.0.0.1:9080", "address of the backing dgraph instance")
var port = flag.Int("p", 3030, "port of this service")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect with db host.../n%s", err)
	}

	log.Println("dialed dgraph")

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to construct listener on port %d...\n%s", *port, err)
	}
	log.Printf("listening at %s", lis.Addr())

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pbabi.RegisterProcStoreServer(grpcServer, procstore.NewProcStoreServer(dg))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpcserver...\n%s", err)
	}

}
