package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/negasus/haproxy-spoe-go/agent"
	"github.com/negasus/haproxy-spoe-go/logger"
	"github.com/negasus/haproxy-spoe-go/message"
	"github.com/negasus/haproxy-spoe-go/request"
)

var cache = make(map[string]*cacheEntry)

type cacheEntry struct {
	Request  string
	Response string
}

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
	if mes, ok := isRequestMessage(req); ok {
		uniqueId, err := getUniqueID(mes)
		if err != nil {
			log.Print(err)
			return
		}

		argBody, ok := mes.KV.Get("body")
		if !ok {
			log.Printf("var 'body' not found in message")
			return
		}

		body, ok := argBody.([]byte)
		if !ok {
			log.Printf("could not assert `body` as []byte")
			return
		}

		if _, ok := cache[uniqueId]; !ok {
			cache[uniqueId] = &cacheEntry{}
		}

		cache[uniqueId].Request = string(body)
	}

	if mes, ok := isResponseMessage(req); ok {
		uniqueId, err := getUniqueID(mes)
		if err != nil {
			log.Print(err)
			return
		}

		argStatus, ok := mes.KV.Get("status")
		if !ok {
			log.Printf("var 'status' not found in message")
			return
		}

		status, ok := argStatus.(int64)
		if !ok {
			log.Printf("could not assert `status` as int64")
			return
		}

		if _, ok := cache[uniqueId]; !ok {
			cache[uniqueId] = &cacheEntry{}
		}

		cache[uniqueId].Response = fmt.Sprintf("%d", status)
	}

	for uniqueId, entry := range cache {
		fmt.Println(uniqueId)
		fmt.Println(entry.Request)
		fmt.Println(entry.Response)
	}
}

func isRequestMessage(req *request.Request) (*message.Message, bool) {
	mes, err := req.Messages.GetByName("request")
	if err != nil {
		return nil, false
	}

	return mes, true
}

func isResponseMessage(req *request.Request) (*message.Message, bool) {
	mes, err := req.Messages.GetByName("response")
	if err != nil {
		return nil, false
	}

	return mes, true
}

func getUniqueID(mes *message.Message) (string, error) {
	uniqueInterface, ok := mes.KV.Get("unique_id")
	if !ok {
		return "", fmt.Errorf("`unique_id` not found in message")
	}

	uniqueId, ok := uniqueInterface.(string)
	if !ok {
		return "", fmt.Errorf("could not assert `unique_id` as string")
	}

	return uniqueId, nil
}
