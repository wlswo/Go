package main

import (
	f "fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	/* func Dial(network, address string) (Conn, error)
	   프로토콜, IP주소 , 포트 번호를 설정하여 서버에 연결
	*/
	client, err := net.Dial("tcp", "localhost:8000") //TCP 프로토콜 로컬에 연결

	if err != nil {
		f.Println(err)
		return
	}

	defer client.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data)
			if err != nil {
				f.Println(err)
				return
			}
			f.Println(string(data[:n]))
			time.Sleep(1 * time.Second) // 지연 호출
		}
	}(client)

	go func(c net.Conn) {
		i := 0

		for {
			s := "Hello" + strconv.Itoa(i)

			_, err := c.Write([]byte(s))

			if err != nil {
				f.Println(err)
				return
			}
			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	f.Scanln()

}
