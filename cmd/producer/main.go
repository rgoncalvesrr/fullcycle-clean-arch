package main

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	rabbitmq.Publish(context.Background(), ch, "amq.direct", "Publicado via GO")
}
