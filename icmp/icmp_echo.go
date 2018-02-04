package main

import (
    "fmt"
    "os"

    "math/rand"

    "golang.org/x/net/ipv4"
    "golang.org/x/net/icmp"

    "github.com/Piasy/tcp_ip_illustrated/utils"
)

func main() {
    conn := utils.DialIP("icmp", "0.0.0.0", os.Args[1])

    body := icmp.Echo{ID: rand.Intn(65536), Seq: 0, Data: []byte("Hello from Piasy :)")}
    message := icmp.Message{Type: ipv4.ICMPTypeEcho, Code: 0, Body: &body}

    utils.SendIcmpPacket(conn, &message)

    fmt.Printf("send ICMP, %s\n", utils.PrintIcmpPacket(&message))

    readBuf := make([]byte, 1024)
    _, _, err := conn.ReadFrom(readBuf)
    if err != nil {
        panic(err)
    }

    reply, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), readBuf)
    if err != nil {
        panic(err)
    }

    fmt.Println("receive ICMP,", utils.PrintIcmpPacket(reply))
}
