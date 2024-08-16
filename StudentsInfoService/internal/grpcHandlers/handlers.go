package grpcHandlers

import (
	"StudentsInfoService/api/api/ServiceApiPb"
	"StudentsInfoService/internal/repositories/interfaces"
	"context"
)

type Handlers struct {
	ServiceApiPb.UnimplementedStudentsInfoServiceServer

	Data interfaces.Repository
}

func New(data interfaces.Repository) *Handlers {
	return &Handlers{Data: data}
}

func (h *Handlers) Read(ctx context.Context, in *ServiceApiPb.ReadRequest) (*ServiceApiPb.ReadResponse, error) {
	info := h.Data.Read()
	students := make([]*ServiceApiPb.ReadResponse_Student, len(info))
	for i, student := range info {
		students[i] = &ServiceApiPb.ReadResponse_Student{
			FirstName:  student.GetFirstName(),
			SecondName: student.GetSecondName(),
			Age:        uint32(student.GetAge()),
			Faculty:    student.GetFaculty(),
			Hostel:     uint32(student.GetHostel()),
			Room:       uint32(student.GetRoom()),
			Id:         uint64(student.GetId()),
		}
	}

	response := ServiceApiPb.ReadResponse{Students: students}

	return &response, nil
}
