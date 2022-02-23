package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

/*

When you run this, it will open a TCP server running on localhost port 5000
If you want to test it on your terminal run this

telnet localhost 5000

Then you can type something like hello and see it printed on the server terminal

To quit telnet press Ctrl + ]
*/
func main() {
	PORT := ":5000"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Println("server running on port ", PORT)

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
