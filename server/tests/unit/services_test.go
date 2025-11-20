package unit

import (
    "testing"

    "github.com/glebarez/sqlite"
    "gorm.io/gorm"

    "server/internal/databases"
    "server/internal/models"
    "server/internal/services"
)

func setupService(t *testing.T) {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        t.Fatalf("open db: %v", err)
    }
    if err := db.AutoMigrate(&models.Student{}, &models.Attendance{}); err != nil {
        t.Fatalf("migrate: %v", err)
    }
    databases.DB = db
}

func TestServices_BasicCRUD(t *testing.T) {
    setupService(t)
    s := models.Student{Name: "Svc", Email: "svc@test.com", Department: "X"}
    created, err := services.CreateStudent(s)
    if err != nil {
        t.Fatalf("create failed: %v", err)
    }
    if created.ID == 0 {
        t.Fatalf("expected id")
    }
    _, err = services.GetStudent(created.ID)
    if err != nil {
        t.Fatalf("get failed: %v", err)
    }
    students, err := services.GetStudents()
    if err != nil || len(students) == 0 {
        t.Fatalf("list failed: %v", err)
    }
    created.Name = "Svc2"
    if err := services.UpdateStudent(&created); err != nil {
        t.Fatalf("update failed: %v", err)
    }
    if err := services.DeleteStudentByID(created.ID); err != nil {
        t.Fatalf("delete failed: %v", err)
    }
}
