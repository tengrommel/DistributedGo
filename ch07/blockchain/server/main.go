package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"CryptoChallenge/blockchain/proto"
	"golang.org/x/net/context"
	"CryptoChallenge/blockchain/server/blockchain"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalf("unable to listen on 8080 port: %v", err)
	}
	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}

type Server struct {
	Blockchain *blockchain.Blockchain
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)
	return &proto.AddBlockResponse{
		Hash:block.Hash,
	}, nil
}

func (s *Server) GetBlockchain(context.Context, *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b :=range s.Blockchain.Blocks{
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash:b.PrevBlockHash,
			Hash: b.Hash,
			Data:b.Data,
		})
	}
	return resp, nil
}
