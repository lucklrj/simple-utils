package grpc

import (
	"net"

	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	return grpc.NewServer()
}
func Run(address string, port uint, server *grpc.Server) error {

	addressListen, err := net.Listen("tcp", address+":"+cast.ToString(port))
	return err
	if err := server.Serve(addressListen); err != nil {
		return err
	}
	return nil
}
