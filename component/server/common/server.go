package common

import (
	"net"

	"github.com/spf13/cast"
)

func Run(address string, port uint, callbackFunc func(conn net.Conn)) error {

	ln, err := net.Listen("tcp", address+":"+cast.ToString(port))
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		} else {
			go callbackFunc(conn)
		}

	}

}
