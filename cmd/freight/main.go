package freight

import (
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucasrgt/go-microservice/internal/freight/domain/entities"
	"github.com/lucasrgt/go-microservice/internal/freight/domain/usecases"
	"github.com/lucasrgt/go-microservice/internal/freight/infra/repositories"
	"github.com/lucasrgt/go-microservice/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/routes?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	repository := repositories.NewRouteRepositorySql(db)
	freight := entities.NewFreight(10)
	createRouteUsecase := usecases.NewCreateRouteUseCase(repository, freight)
	changeRouteStatusUsecase := usecases.NewChangeRouteStatusUseCase(repository)

	msgChan := make(chan *ckafka.Message)
	topics := []string{"routes"}
	servers := "host.docker.internal:9094"
	go kafka.Consume(topics, servers, msgChan)

	for msg := range msgChan {
		input := usecases.CreateRouteInput{}
		err := json.Unmarshal(msg.Value, &input)
		if err != nil {
			panic(err)
		}

		switch input.Event {
		case "RouteCreated":
			output, err := createRouteUsecase.Call(input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(output)

		case "RouteStarted", "RouteFinished":
			input := usecases.ChangeRouteStatusInput{}
			err := json.Unmarshal(msg.Value, &input)
			if err != nil {
				panic(err)
			}
			output, err := changeRouteStatusUsecase.Call(input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(output)
		}
	}
}
