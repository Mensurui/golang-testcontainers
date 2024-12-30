package service_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/Mensurui/golang-testcontainers/project/data"
	"github.com/Mensurui/golang-testcontainers/project/service"
	"github.com/Mensurui/golang-testcontainers/project/testhelpers"
	protos "github.com/Mensurui/golang-testcontainers/protobuf/gen"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	models      *data.Models
	ctx         context.Context
}

func (suite *ServiceTestSuite) SetupSuite() {
	fmt.Println("----------Setup Suite----------")
	suite.ctx = context.Background()
	container, err := testhelpers.CreateContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = container
	db, err := sql.Open("postgres", suite.pgContainer.Connectionstring)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	suite.models = data.NewModels(db)
}

func (suite *ServiceTestSuite) TeardownSuite() {
	fmt.Println("---------- Tear Down Suite----------")
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("Failed to terminate container: %v", err)
	}
}

func (suite *ServiceTestSuite) TestAddUser() {
	fmt.Println("----------Test Add User----------")
	svc := service.NewService(suite.models.User.DB)
	req := &protos.AddUserRequest{
		UserName: "Mensur",
		Email:    "mensur@email.com",
		Age:      5,
	}
	resp, err := svc.AddUser(suite.ctx, req)
	suite.NoError(err)
	suite.Equal("Registered Successfully", resp.Message)
}

func (suite *ServiceTestSuite) TestCheckUser() {
	fmt.Println("----------Test Check User----------")
	type User struct {
		UserName string
		Age      string
	}
	svc := service.NewService(suite.models.User.DB)
	reqAdd := &protos.AddUserRequest{
		UserName: "John",
		Email:    "john@gmail.com",
		Age:      5,
	}
	respAdd, err := svc.AddUser(suite.ctx, reqAdd)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	log.Printf("Log: %v", respAdd)

	reqCheck := &protos.CheckUserRequest{
		UserID: 2,
	}
	userCheck := User{
		UserName: "John",
		Age:      "john@gmail.com",
	}
	respCheck, err := svc.CheckUser(suite.ctx, reqCheck)
	suite.NoError(err)
	suite.Equal(userCheck.UserName, respCheck.GetUser())
	suite.Equal(userCheck.Age, respCheck.GetAge())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
