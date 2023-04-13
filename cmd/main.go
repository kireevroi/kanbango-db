package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kireevroi/kanbango/auth/internal/db"
	"github.com/kireevroi/kanbango/auth/internal/userproto"
	"github.com/kireevroi/kanbango/auth/pkg/onstart"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	d := db.NewDB()
	onstart.LoadEnv()
	d.Connect(os.Getenv("DBURL"))
	userproto.RegisterUserServiceServer(s, &userproto.Server{DB: d})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}