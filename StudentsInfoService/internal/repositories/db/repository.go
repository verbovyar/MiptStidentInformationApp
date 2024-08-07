package db

import (
	"StudentsInfoService/internal/domain"
	"StudentsInfoService/internal/repositories/interfaces"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentsRepository struct {
	Pool  *pgxpool.Pool
	iface interfaces.Repository
}

func New(pool *pgxpool.Pool) *StudentsRepository {
	return &StudentsRepository{
		Pool: pool,
	}
}

func (r *StudentsRepository) Create(user *domain.Student) error {
	ctx := context.Background()

	var facultyId uint
	query := `SELECT facultyId FROM Faculty WHERE facultyName = $1`
	err := r.Pool.QueryRow(ctx, query, user.Faculty).Scan(&facultyId)
	if err != nil {
		fmt.Printf("Create Faculty error %v\n", err)
		return err
	}

	var hostelId uint
	query = `SELECT hostelId FROM Hostel WHERE hostelId = $1`
	err = r.Pool.QueryRow(ctx, query, user.Hostel).Scan(&hostelId)
	if err != nil {
		fmt.Printf("Create Hostel error %v\n", err)
		return err
	}

	query = `INSERT INTO Student (secondName, firstName, age, faculty_id, hostel_id, room) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.Pool.Query(ctx, query, user.SecondName, user.FirstName, user.Age, facultyId, hostelId, user.Room)
	if err != nil {
		fmt.Printf("Create Insert error %v\n", err)
		return err
	}

	var id int
	query = `SELECT id FROM Student WHERE firstName = $1`
	err = r.Pool.QueryRow(ctx, query, user.FirstName).Scan(&id)
	if err != nil {
		fmt.Printf("Create Id error %v\n", err)
		return err
	}
	user.Id = uint(id)

	return nil
}

func (r *StudentsRepository) Read() []*domain.Student {
	return nil
}

func (r *StudentsRepository) Update(user *domain.Student, id uint) error {
	ctx := context.Background()

	var facultyId uint
	query := `SELECT facultyId FROM Faculty WHERE facultyName = $1`
	err := r.Pool.QueryRow(ctx, query, user.Faculty).Scan(&facultyId)
	if err != nil {
		fmt.Printf("Update Faculty error %v\n", err)
		return err
	}

	var hostelId uint
	query = `SELECT hostelId FROM Hostel WHERE hostelId = $1`
	err = r.Pool.QueryRow(ctx, query, user.Hostel).Scan(&hostelId)
	if err != nil {
		fmt.Printf("Update Hostel error %v\n", err)
		return err
	}

	query = `UPDATE Student SET secondName = $1, firstName = $2, age = $3, faculty_id = $4, hostel_id = $5, room = $6 WHERE id = $7`
	_, err = r.Pool.Query(ctx, query, user.SecondName, user.FirstName, user.Age, facultyId, hostelId, user.Room, id)
	if err != nil {
		fmt.Printf("Update Insert error %v\n", err)
		return err
	}

	return nil
}

func (r *StudentsRepository) Delete(id uint) error {
	ctx := context.Background()

	query := `DELETE FROM Student WHERE id = $1`
	_, err := r.Pool.Query(ctx, query, id)
	if err != nil {
		fmt.Printf("Delete error %v\n", err)
		return err
	}

	return nil
}
