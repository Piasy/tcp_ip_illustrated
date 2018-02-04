package main

import (
    "net"
    "fmt"

    "golang.org/x/net/ipv4"
    "golang.org/x/net/icmp"
)

func main() {
    addr, _ := net.ResolveIPAddr("ip4", "0.0.0.0")
    conn, _ := net.ListenIP("ip4:icmp", addr)

    buf := make([]byte, 1024)

    for ; ; {
        _, _, err := conn.ReadFrom(buf)
        if err != nil {
            fmt.Printf("Error: %#v\n", err)
            break
        }

        packet, err := icmp.ParseMessage(ipv4.ICMPType(0).Protocol(), buf)
        if err != nil {
            fmt.Printf("Error: %#v\n", err)
            break
        }

        fmt.Printf("receive ICMP, type:%v, code:%d, checksum:%d, ", packet.Type, packet.Code, packet.Checksum)

        echo, ok := packet.Body.(*icmp.Echo)
        if ok {
            fmt.Printf("id:%d, seq:%d\n", echo.ID, echo.Seq)
        } else {
            fmt.Print("\n")
        }
    }
}
