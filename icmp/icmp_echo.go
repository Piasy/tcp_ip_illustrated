package main

import (
    "net"
    "fmt"

    "math/rand"

    "golang.org/x/net/ipv4"
    "golang.org/x/net/icmp"

    "github.com/Piasy/tcp_ip_illustrated/utils"
)

func main() {
    localAddr, err := net.ResolveIPAddr("ip4", "0.0.0.0")
    if err != nil {
        panic(err)
    }

    remoteAddr, err := net.ResolveIPAddr("ip4", "www.baidu.com")
    if err != nil {
        panic(err)
    }

    conn, err := net.DialIP("ip4:icmp", localAddr, remoteAddr)
    if err != nil {
        panic(err)
    }

    echoBody := icmp.Echo{ID: rand.Intn(65536), Seq: 0, Data: []byte("Hello from Piasy :)")}
    echoMessage := icmp.Message{Type: ipv4.ICMPTypeEcho, Code: 0, Body: &echoBody}

    writeBuf, err := echoMessage.Marshal(nil)
    if err != nil {
        panic(err)
    }

    _, err = conn.Write(writeBuf)
    if err != nil {
        panic(err)
    }

    fmt.Printf("send ICMP, %s\n", utils.PrintIcmpPacket(&echoMessage))

    readBuf := make([]byte, 1024)
    _, _, err = conn.ReadFrom(readBuf)
    if err != nil {
        panic(err)
    }

    echoReply, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), readBuf)
    if err != nil {
        panic(err)
    }

    fmt.Println("receive ICMP,", utils.PrintIcmpPacket(echoReply))
}
