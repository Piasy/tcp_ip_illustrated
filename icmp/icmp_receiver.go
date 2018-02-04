package main

import (
    "net"
    "fmt"

    "golang.org/x/net/ipv4"
    "golang.org/x/net/icmp"

    "github.com/Piasy/tcp_ip_illustrated/utils"
)

func main() {
    addr, err := net.ResolveIPAddr("ip4", "0.0.0.0")
    if err != nil {
        panic(err)
    }

    conn, err := net.ListenIP("ip4:icmp", addr)
    if err != nil {
        panic(err)
    }

    buf := make([]byte, 1024)

    for ; ; {
        _, _, err = conn.ReadFrom(buf)
        if err != nil {
            panic(err)
        }

        packet, err := icmp.ParseMessage(ipv4.ICMPType(0).Protocol(), buf)
        if err != nil {
            panic(err)
        }

        fmt.Println("receive ICMP,", utils.PrintIcmpPacket(packet))
    }
}
