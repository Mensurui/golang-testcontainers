package main

import (
	"database/sql"
	"flag"
	"log"
	"net"
	"os"

	"github.com/Mensurui/golang-testcontainers/project/service"
	protos "github.com/Mensurui/golang-testcontainers/protobuf/gen"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	port string
	db   struct {
		address string
	}
}

func main() {
	var cfg config
	dbAddress := os.Getenv("TESTCONADDR")
	testingPort := os.Getenv("TESTPORTADDR")
	flag.StringVar(&cfg.db.address, "addr", dbAddress, "Use to update the address of the database address.")
	flag.StringVar(&cfg.port, "port", testingPort, "Use to update the port number of the service.")

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Database connected")
	defer db.Close()

	gs := grpc.NewServer()
	service := service.NewService(db)

	protos.RegisterUserServiceServer(gs, service)
	reflection.Register(gs)

	l, err := net.Listen("tcp", testingPort)
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	err = gs.Serve(l)
	if err != nil {
		log.Fatalf("Error serving on grpc: %v", err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	address := cfg.db.address
	db, err := sql.Open("postgres", address)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
