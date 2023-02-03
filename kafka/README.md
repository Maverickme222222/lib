### Publish Messages to Kafka


```go

package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka/producer"
)

func main() {

	var cfg struct {
		conf.Version

		Kafka struct {
			BrokerURL   string `conf:"env:KAFKA_BROKER_URL,required:True"`
			Topic       string `conf:"env:KAFKA_TOPIC_ENTITY_CREATED"`
			Username    string `conf:"env:KAFKA_USERNAME"`
			Password    string `conf:"env:KAFKA_PASSWORD"`
			GroupID     string `conf:"env:KAFKA_GROUP_ID"`
			DisableTLS  bool   `conf:"env:KAFKA_DISABLE_TLS"`
			DisableSASL bool   `conf:"env:KAFKA_DISABLE_SASL"`
		}
	}
	prod := producer.NewKappaProducer(kafka.Config{
		BrokerURL:   cfg.Kafka.BrokerURL,
		Topic:       cfg.Kafka.Topic,
		Username:    cfg.Kafka.Username,
		Password:    cfg.Kafka.Password,
		GroupID:     cfg.Kafka.GroupID,
		DisableTLS:  cfg.Kafka.DisableTLS,
		DisableSASL: cfg.Kafka.DisableSASL,
	})

	user := models.User{
		Id:         uuid.New().String(),
		FirstName:  "Kevin",
		MiddleName: "Hungai",
		LastName:   "Amuhinda",
		Email: &models.Email{
			Address: "hungaikevin@gmail.com",
		},
		Phone: &models.Phone{
			Number: "0789898982",
		},
	}

	if err := prod.Send(context.Background(), []byte("key"), user); err != nil {
		fmt.Println(err.Error())
	}
}

```


### Consume Messages from Kafka

```go

package main

import (
	"context"
	"expvar"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ardanlabs/conf/v3"
	"github.com/pkg/errors"

	"github.com/joho/godotenv"
	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka/consumer"
)

func main() {


	var cfg struct {
		conf.Version

		Kafka struct {
			BrokerURL   string `conf:"env:KAFKA_BROKER_URL,required:True"`
			Topic       string `conf:"env:KAFKA_TOPIC_ENTITY_CREATED"`
			Username    string `conf:"env:KAFKA_USERNAME"`
			Password    string `conf:"env:KAFKA_PASSWORD"`
			GroupID     string `conf:"env:KAFKA_GROUP_ID"`
			DisableTLS  bool   `conf:"env:KAFKA_DISABLE_TLS"`
			DisableSASL bool   `conf:"env:KAFKA_DISABLE_SASL"`
		}
	}


	con := consumer.NewKappaConsumer(kafka.Config{
		BrokerURL:             cfg.Kafka.BrokerURL,
		Topic:                 cfg.Kafka.Topic,
		Username:              cfg.Kafka.Username,
		Password:              cfg.Kafka.Password,
		GroupID:               cfg.Kafka.GroupID,
		EnableDeadLetterTopic: true,
		DisableTLS:            cfg.Kafka.DisableTLS,
		DisableSASL:           cfg.Kafka.DisableSASL,
	})

	defer con.Close()


	for {
		user := models.User{}
		err := con.Read(context.Background(), &user)
		if err != nil {
			fmt.Println(err)
			consumerErrs <- err
			continue
		}

		if user.FirstName == "" {
			log.Error().Msg("first name not defined for user")
			continue
		}
		fmt.Println("message from kafka:", user)
		con.Commit(ctx)
	}

}


```