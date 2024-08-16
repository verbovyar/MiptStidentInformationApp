package addHandler

import (
	"StudentsInfoService/internal/domain"
	"StudentsInfoService/internal/repositories/interfaces"
	"context"
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"github.com/allegro/bigcache/v3"
	"log"
	"time"
)

type AddHandler struct {
	producer      sarama.SyncProducer
	consumerGroup sarama.ConsumerGroup
	data          interfaces.Repository
	cache         *bigcache.BigCache
}

type addRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Age        uint   `json:"age"`
	Faculty    string `json:"faculty"`
	Hostel     uint   `json:"hostel"`
	Room       uint   `json:"room"`
}

type addResponse struct {
	Id uint `json:"id"`
}

func NewAddHandler(producer sarama.SyncProducer, data interfaces.Repository, consumerGroup sarama.ConsumerGroup, cache *bigcache.BigCache) *AddHandler {
	return &AddHandler{
		producer:      producer,
		data:          data,
		consumerGroup: consumerGroup,
		cache:         cache,
	}
}

func (h *AddHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *AddHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *AddHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var request addRequest
		err := json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Printf("income data %v: %v", string(msg.Value), err)
			continue
		}

		student, err := domain.New(request.FirstName, request.SecondName, request.Age, request.Faculty, request.Hostel, request.Room)

		data, err := h.cache.Get("task:Create")
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			err = h.data.Create(student)
			if err != nil {
				return err
			}

			marshalData, _ := json.Marshal(student)
			h.cache.Set("task:Create", marshalData)
		} else {
			json.Unmarshal(data, &student)
		}

		addResp := addResponse{Id: student.Id}
		response, _ := json.Marshal(&addResp)

		producerMsg := &sarama.ProducerMessage{
			Topic:     "AddResponse",
			Partition: -1,
			Value:     sarama.ByteEncoder(response),
		}

		_, _, err = h.producer.SendMessage(producerMsg)

	}

	return nil
}

func AddClaim(addHandler *AddHandler) {
	for {
		if err := addHandler.consumerGroup.Consume(context.Background(), []string{"AddRequest"}, addHandler); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 10)
		}
	}
}
