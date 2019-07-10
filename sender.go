// UPnP, SSDP, finding router

package main

import (
    "fmt"
    "net"
    "strings"
)

func main() {

    ipv4 := "239.255.255.250"
    port := ":1900"
    addr := ipv4 + port

    format := `M-SEARCH * HTTP/1.1
HOST: 239.255.255.250:1900
MAN: "ssdp:discover"
MX: 3
USER-AGENT: sender/1.0
ST: urn:schemas-upnp-org:service:WANPPPConnection:1

`
//ST: urn:schemas-upnp-org:device:InternetGatewayDevice:1

    udp_addr, err := net.ResolveUDPAddr("udp", addr)
    _error(err)

    listener, err := net.ListenMulticastUDP("udp", nil, udp_addr)
    _error(err)
    defer listener.Close()


    conn, err := net.Dial("udp", addr)
    _error(err)
    defer conn.Close()

    message := []byte(strings.Replace(format, "\n", "\r\n", -1))
    conn.Write([]byte(message))
    fmt.Printf("%s\n", message)

    buffer := make([]byte, 1240)
    for {
        length, remoteAddress, err := listener.ReadFrom(buffer)
        _error(err)

        fmt.Printf("Sender: %v\n", remoteAddress)
        fmt.Printf("Contents: %s\n", string(buffer[:length]))
    }

}

func _error(_err error) {
    if _err != nil {
        panic(_err)
    }
}
