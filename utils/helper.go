package utils

import (
    "fmt"
    "golang.org/x/net/icmp"
)

func PrintIcmpPacket(packet *icmp.Message) string {
    if echo, ok := packet.Body.(*icmp.Echo); ok {
        return fmt.Sprintf("type:%v, code:%d, checksum:%d, id:%d, seq:%d",
            packet.Type, packet.Code, packet.Checksum, echo.ID, echo.Seq)
    } else {
        return fmt.Sprintf("type:%v, code:%d, checksum:%d", packet.Type, packet.Code, packet.Checksum)
    }
}
