package service

import (
	"sre-bootcamp/internal/db"
	"sre-bootcamp/internal/model"
	"sre-bootcamp/internal/repository"

	"github.com/google/uuid"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService() *StudentService {
	db := db.GetDB()
	return &StudentService{repo: repository.NewStudentRepository(db)}
}

func (s *StudentService) CreateStudent(stu *model.Student) (*model.Student, error) {
	stu.ID = uuid.NewString()
	if err := s.repo.Create(stu); err != nil {
		return nil, err
	}
	return stu, nil
}

func (s *StudentService) GetAllStudents() ([]model.Student, error) {
	return s.repo.FindAll()
}

func (s *StudentService) GetStudent(id string) (*model.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) UpdateStudent(id string, stu *model.Student) (*model.Student, error) {
	if err := s.repo.Update(id, stu); err != nil {
		return nil, err
	}
	stu.ID = id
	return stu, nil
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.repo.Delete(id)
}
