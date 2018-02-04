package main

import (
    "fmt"
    "time"
    "os"
    "encoding/binary"

    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"

    "github.com/Piasy/tcp_ip_illustrated/utils"
)

func main() {
    conn := utils.DialIP("icmp", "0.0.0.0", os.Args[1])

    now := time.Now()
    secondsAfterMidnightUtc := now.Unix() - time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()

    buf := make([]byte, 16)

    binary.BigEndian.PutUint16(buf[0:2], 0)
    binary.BigEndian.PutUint16(buf[2:4], 0)
    binary.BigEndian.PutUint32(buf[4:8], uint32(secondsAfterMidnightUtc*1000))
    binary.BigEndian.PutUint32(buf[8:12], 0)
    binary.BigEndian.PutUint32(buf[12:16], 0)

    message := icmp.Message{Type: ipv4.ICMPTypeTimestamp, Code: 0, Body: &icmp.DefaultMessageBody{Data: buf}}

    utils.SendIcmpPacket(conn, &message)

    fmt.Printf("send ICMP, %s\n", utils.PrintIcmpPacket(&message))

    readBuf := make([]byte, 1024)
    _, _, err := conn.ReadFrom(readBuf)
    if err != nil {
        panic(err)
    }

    reply, err := icmp.ParseMessage(ipv4.ICMPType(0).Protocol(), readBuf)
    if err != nil {
        panic(err)
    }

    fmt.Println("receive ICMP,", utils.PrintIcmpPacket(reply))
}
