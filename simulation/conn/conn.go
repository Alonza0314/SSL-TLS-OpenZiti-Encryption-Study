package conn

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

type ListenerInterface func(string, HandlerInterface)

type HandlerInterface func(conn net.Conn)

func TCPListener(serverCfg string, handler HandlerInterface) {
	log.Println("=> Server Starting...")

	viper.SetConfigFile(serverCfg)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config:\n\t%s\n", err.Error())
	}

	host, port := viper.GetString("server.addr"), viper.GetInt("server.port")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Fatalf("failed to listen on %s:%v:\n\t%s\n", host, port, err.Error())
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
				log.Printf("failed to accept connection:\n\t%s\n", err.Error())
				continue
			}
			log.Println(NEW, "CONNECTION from", conn.RemoteAddr())
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
	if err := conn.SetReadDeadline(time.Now().Add(CONN_TIMEOUT)); err != nil {
		log.Printf("failed to set read timeoute:\n\t%s\n", err.Error())
		return
	}
	var req request
	if err := json.NewDecoder(conn).Decode(&req); err != nil {
		log.Printf("failed to decode request:\n\t%s\n", err.Error())
		return
	}
	log.Println(RECEIVED, CLIENT_HELLO, "from", conn.RemoteAddr())
	if err := conn.SetReadDeadline(time.Time{}); err != nil {
		log.Printf("failed to reset read timeout:\n\t%s\n", err.Error())
		return
	}
}
