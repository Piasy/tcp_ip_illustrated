package utils

import (
    "fmt"
    "net"

    "encoding/binary"
    "golang.org/x/net/icmp"
)

func PrintIcmpPacket(packet *icmp.Message) string {
    switch body := packet.Body.(type) {
    case *icmp.Echo:
        return fmt.Sprintf("type:%v, code:%d, checksum:%d, id:%d, seq:%d",
            packet.Type, packet.Code, packet.Checksum, body.ID, body.Seq)
    case *icmp.DefaultMessageBody:
        id := binary.BigEndian.Uint16(body.Data[0:2])
        seq := binary.BigEndian.Uint16(body.Data[2:4])

        var w1, w2, w3 uint32

        if len(body.Data) >= 8 {
            w1 = binary.BigEndian.Uint32(body.Data[4:8])
        }
        if len(body.Data) >= 12 {
            w2 = binary.BigEndian.Uint32(body.Data[8:12])
        }
        if len(body.Data) >= 16 {
            w3 = binary.BigEndian.Uint32(body.Data[12:16])
        }

        return fmt.Sprintf("type:%v, code:%d, checksum:%d, id:%d, seq:%d, w1:%d, w2:%d, w3:%d",
            packet.Type, packet.Code, packet.Checksum, id, seq, w1, w2, w3)
    default:
        return fmt.Sprintf("type:%v, code:%d, checksum:%d", packet.Type, packet.Code, packet.Checksum)
    }
}

func DialIP(proto, srcAddr, dstAddr string) *net.IPConn {
    localAddr, err := net.ResolveIPAddr("ip4", srcAddr)
    if err != nil {
        panic(err)
    }

    remoteAddr, err := net.ResolveIPAddr("ip4", dstAddr)
    if err != nil {
        panic(err)
    }

    conn, err := net.DialIP("ip4:"+proto, localAddr, remoteAddr)
    if err != nil {
        panic(err)
    }

    return conn
}

func SendIcmpPacket(conn *net.IPConn, packet *icmp.Message) {
    writeBuf, err := packet.Marshal(nil)
    if err != nil {
        panic(err)
    }

    _, err = conn.Write(writeBuf)
    if err != nil {
        panic(err)
    }
}
