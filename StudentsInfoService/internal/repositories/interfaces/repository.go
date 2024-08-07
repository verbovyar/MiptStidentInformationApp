package interfaces

import (
	"StudentsInfoService/internal/domain"
)

type Repository interface {
	Create(player *domain.Student) error
	Read() []*domain.Student
	Update(user *domain.Student, id uint) error
	Delete(id uint) error
}
