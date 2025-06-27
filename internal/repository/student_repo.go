package repository

import (
	"database/sql"
	"sre-bootcamp/internal/model"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(student *model.Student) error {
	_, err := r.db.Exec("INSERT INTO students (id, name, email) VALUES ($1, $2, $3)", student.ID, student.Name, student.Email)
	return err
}

func (r *StudentRepository) FindAll() ([]model.Student, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var stu model.Student
		if err := rows.Scan(&stu.ID, &stu.Name, &stu.Email); err != nil {
			return nil, err
		}
		students = append(students, stu)
	}
	return students, nil
}

func (r *StudentRepository) FindByID(id string) (*model.Student, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM students WHERE id=$1", id)
	var stu model.Student
	if err := row.Scan(&stu.ID, &stu.Name, &stu.Email); err != nil {
		return nil, err
	}
	return &stu, nil
}

func (r *StudentRepository) Update(id string, student *model.Student) error {
	_, err := r.db.Exec("UPDATE students SET name=$1, email=$2 WHERE id=$3", student.Name, student.Email, id)
	return err
}

func (r *StudentRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM students WHERE id=$1", id)
	return err
}
