package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(10, 61, 38, 244),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("dial failed,err:%v\n", err)
		return
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Printf("send to server failed,err:%v\n", err)
			return
		}
		//接收数据
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("recv from udp failed,err:%v\n", err)
			return
		}
		fmt.Printf("read from:%v,msg:%v\n", addr, string(buf[:n]))
	}
}
