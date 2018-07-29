package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"CryptoChallenge/blockchain/proto"
	"context"
	"time"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "add the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("cannot dial server: %v", err)
	}
	client = proto.NewBlockchainClient(conn)
	if *addFlag{
		addBlock()
	}
	if *listFlag{
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil{
		log.Fatal("unable to add block: %v", err)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain()  {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil{
		log.Fatalf("unable to get blockchain: %v", err)
	}
	log.Print("blocks:")
	for _, b := range bc.Blocks{
		log.Printf("hash : %s, prev hash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
	}
}
