package handlers

import (
	"StudentsInfoService/internal/domain"
	"StudentsInfoService/internal/repositories/db"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Handlers struct {
	Router        *mux.Router
	Data          *db.StudentsRepository
	Producer      *sarama.SyncProducer
	ConsumerGroup *sarama.ConsumerGroup
}

func New(repo *db.StudentsRepository, producer *sarama.SyncProducer, consumerGroup *sarama.ConsumerGroup) *Handlers {
	h := &Handlers{}

	router := mux.NewRouter()
	h.Router = router
	h.Data = repo
	h.Producer = producer
	h.ConsumerGroup = consumerGroup

	// может принимать регулярку в path, вместо корневого пути
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

// ------------------------

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

	h.Data.Create(&student) // TODO handle error

	w.Write([]byte("Юзер успешно добавлен"))
}

// ------------------------
type readRequest struct {
}

func (h *Handlers) Read(w http.ResponseWriter, r *http.Request) {
	var list []*domain.Student
	list = h.Data.Read()

	for _, value := range list {
		v, err := json.Marshal(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(v)
	}
}

// ------------------------

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.Data.Update(&student, student.Id) // TODO handle error

	w.Write([]byte("Юзер успешно обновлен"))
}

// ------------------------
type deleteRequest struct {
	Id uint `json:"id"`
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var student deleteRequest
	if err = json.Unmarshal(body, &student); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.Data.Delete(student.Id) // TODO handle error

	w.Write([]byte("Юзер удален"))
}
