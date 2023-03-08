package main

import (
	"log"
	"net"
	"os"

	"github.com/negasus/haproxy-spoe-go/agent"
	"github.com/negasus/haproxy-spoe-go/logger"
	"github.com/negasus/haproxy-spoe-go/request"
)

func main() {
	log.Print("Listening on 0.0.0.0:9090")

	listener, err := net.Listen("tcp4", "0.0.0.0:9090")
	if err != nil {
		log.Printf("error create listener, %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	a := agent.New(handler, logger.NewDefaultLog())

	if err := a.Serve(listener); err != nil {
		log.Printf("error agent serve: %+v\n", err)
	}
}

func handler(req *request.Request) {
	log.Printf("Handle request EngineID: '%s', StreamID: '%d', FrameID: '%d' with %d messages\n", req.EngineID, req.StreamID, req.FrameID, req.Messages.Len())

	for _, message := range *req.Messages {
		log.Printf("%v\n", message.Name)
		for _, item := range message.KV.Data() {
			log.Printf("%v\n", item.Name)
			log.Printf("%v\n", item.Value)
		}
		log.Printf(message.Name)
	}

}
