package main

import (
	"log"
	"net/http"

	proto "github.com/sk-develop/grpc-sample/hello-api/hello-proto"

	echo "github.com/labstack/echo/v4"
	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	e := echo.New()
	
	conn, err := grpc.Dial(
        address,
        grpc.WithInsecure(),
        grpc.FailOnNonTempDialError(true),
        grpc.WithBlock(), 
    )
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	
	req := &proto.HelloRequest{Name: "せいやですよ！"}
	e.GET("/", func(c echo.Context) error {
		log.Print("リクエストを受け付けました")
		if response, err := client.SayHello(c.Request().Context(), req); err == nil {
			log.Print("レスポンスを返却しました")
			return c.JSON(http.StatusOK, response)
		} else {
			return c.JSON(http.StatusInternalServerError, response)
		}
	})
	
	e.Logger.Fatal(e.Start(":1323"))
}