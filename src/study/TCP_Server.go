package main

import (
	f "fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
			continue
		} else if conn != nil {
			f.Println("연결완료")
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			f.Println(err)
			return
		}

		f.Println(string(data[:n]))
		_, err = c.Write(data[:n])
		if err != nil {
			f.Println(err)
			return
		}
	}
}
