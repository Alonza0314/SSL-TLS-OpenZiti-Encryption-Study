package network

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ListenerInterface func(string, int, HandlerInterface)

type HandlerInterface func(conn net.Conn)

func TCPListener(host string, port int, handler HandlerInterface) {
	log.Println("=> Server Starting...")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Fatalf("Failed to listen on %s:%v: %s\n", host, port, err.Error())
	}
	defer listener.Close()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println()
		log.Println("==> Server Stopping...")
		cancel()
		listener.Close()
	}()

	log.Println("===== Server Started =====")
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			log.Println("===== Server Stopped =====")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				if ne, ok := err.(*net.OpError); ok && ne.Err.Error() == "use of closed network connection" {
					continue
				}
				log.Printf("Failed to accept connection: %s\n", err.Error())
				continue
			}
			wg.Add(1)
			go func(conn net.Conn) {
				defer wg.Done()
				defer conn.Close()
				handler(conn)
			}(conn)
		}
	}
}

func TCPHandler(conn net.Conn) {
}
