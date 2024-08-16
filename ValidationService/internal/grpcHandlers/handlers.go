package grpcHandlers

import (
	"Application/ValidationService/api/api/ValidationApiPb"
	"context"
)

type Handlers struct {
	ValidationApiPb.UnimplementedValidationServiceServer

	client ValidationApiPb.ValidationServiceClient
}

func New(client ValidationApiPb.ValidationServiceClient) *Handlers {
	return &Handlers{client: client}
}

func (h *Handlers) Read(ctx context.Context, in *ValidationApiPb.ReadRequest) (*ValidationApiPb.ReadResponse, error) {
	resp, read := h.client.Read(ctx, in)
	if read != nil {
		return nil, read
	}

	students := make([]*ValidationApiPb.ReadResponse_Student, len(resp.Students))
	for i, student := range resp.Students {
		students[i] = &ValidationApiPb.ReadResponse_Student{
			FirstName:  student.FirstName,
			SecondName: student.SecondName,
			Age:        student.Age,
			Faculty:    student.Faculty,
			Hostel:     student.Hostel,
			Room:       student.Room,
			Id:         student.Id,
		}
	}

	readResponse := ValidationApiPb.ReadResponse{Students: students}

	return &readResponse, nil
}
