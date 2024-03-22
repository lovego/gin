package gin

import (
	"context"
	"log"
	"net"
	"syscall"
)

var listenControl func(network, address string, c syscall.RawConn) error

func getListener(addr string) net.Listener {
	listenConfig := net.ListenConfig{Control: listenControl}
	listener, err := listenConfig.Listen(context.Background(), `tcp`, addr)
	if err != nil {
		log.Panic(err)
	}
	log.Println(`backend started.(` + addr + `)`)
	return listener.(*net.TCPListener)
}
