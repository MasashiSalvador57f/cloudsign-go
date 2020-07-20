package main

import (
	"context"
	"fmt"
	"os"

	"github.com/MasashiSalvador57f/cloudsign-go"
)

func main() {
	clientID := os.Getenv("CLOUDSIGN_CLIENT_ID")
	cli, err := cloudsign.NewClient(clientID, nil, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(cli)
	at, err := cli.IssueAccessToken(context.Background())
	fmt.Println(at)
	fmt.Println(err)
}
