package unit

import (
    "errors"
    "testing"

    "server/internal/models"
)

type StudentRepo interface {
    Create(*models.Student) error
    FindByEmail(string) (models.Student, error)
}

type MockRepo struct{
    created bool
}

func (m *MockRepo) Create(s *models.Student) error {
    m.created = true
    s.ID = 1
    return nil
}

func (m *MockRepo) FindByEmail(email string) (models.Student, error) {
    if email == "exists@test.com" {
        return models.Student{ID:1, Email: email}, nil
    }
    return models.Student{}, errors.New("not found")
}

func TestMockRepoBehavior(t *testing.T) {
    var repo StudentRepo = &MockRepo{}
    s := &models.Student{Name: "M", Email: "m@test.com"}
    if err := repo.Create(s); err != nil {
        t.Fatalf("create mock failed: %v", err)
    }
    if s.ID == 0 {
        t.Fatalf("mock create did not set id")
    }
    if _, err := repo.FindByEmail("nope@test.com"); err == nil {
        t.Fatalf("expected not found")
    }
}
