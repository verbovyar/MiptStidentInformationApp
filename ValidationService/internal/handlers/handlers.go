package handlers

import (
	"Application/ValidationService/internal/domain"
	"Application/ValidationService/utils"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Handlers struct {
	Router   *mux.Router
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
}

func New(producer sarama.SyncProducer, consumer sarama.Consumer) *Handlers {
	h := &Handlers{}
	router := mux.NewRouter()
	h.Router = router
	h.Producer = producer
	h.Consumer = consumer

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.Read(w, r)
		case http.MethodPut:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			log.Fatal("unknown method")
		}
	})

	return h
}

// Create ------------------------
func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var student domain.Student
	if err = json.Unmarshal(body, &student); err != nil {
		return
	}

	if err = utils.CreateValidation(student); err != nil {
		w.Write([]byte("Неверные данные"))
		return
	}

	msg := &sarama.ProducerMessage{
		Topic:     "AddRequest",
		Partition: -1,
		Value:     sarama.ByteEncoder(body),
	}
	_, _, err = h.Producer.SendMessage(msg)
	if err != nil {
		//TODO Logger
		error.Error(err)
	}

	//time.Sleep(time.Second * 2)
	log.Println("TAG1")
	claim, err := h.Consumer.ConsumePartition("AddResponse", 0, sarama.OffsetOldest)
	if err != nil {
		//TODO Logger
		error.Error(err)
	}

	log.Println("TAG2")
	select {
	case err = <-claim.Errors():
		log.Println(err)
	default:
	}

	w.Write([]byte("Юзер успешно добавлен"))
	log.Println("TAG3")
}

// Read ------------------------
type readRequest struct {
}

func (h *Handlers) Read(w http.ResponseWriter, r *http.Request) {

}

// Update ------------------------
func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ------------------------
type deleteRequest struct {
	Id uint `json:"id"`
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {

}
