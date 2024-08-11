package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// create a topic
func main() {
	// to create topics when auto.create.topics.enable='true'
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:39094", "topic-A", 0)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
}
