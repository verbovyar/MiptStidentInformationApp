package utils

import (
	"Application/ValidationService/internal/domain"
	"errors"
)

const (
	createError string = "Bad name: "
)

func CreateValidation(student domain.Student) error {
	if student.FirstName == "" {
		return errors.New(createError + " empty string")
	}

	if student.SecondName == "" {
		return errors.New(createError + " empty string")
	}

	if student.Faculty == "" {
		return errors.New(createError + " empty string")
	}

	if student.Room < 0 {
		return errors.New(createError + " incorrect number")
	}

	return nil
}
