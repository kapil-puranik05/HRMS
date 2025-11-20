package unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"server/internal/databases"
	"server/internal/models"
	"server/internal/routes"
)

func setup(t *testing.T) *gin.Engine {
	os.Setenv("TEST_EMAIL", "admin@example.com")
	os.Setenv("TEST_PASSWORD", "secret")
	os.Setenv("JWT_KEY", "testkey")

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&models.Student{}, &models.Attendance{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	databases.DB = db
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routes.StudentRoutes(r)
	routes.AttendanceRoutes(r)
	routes.LoginRoutes(r)
	return r
}

func TestControllers_CreateAndGetStudent(t *testing.T) {
	r := setup(t)

	// create
	payload := map[string]string{"name": "A", "email": "a@test.com", "department": "D"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/students/", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	// need login token; bypass by setting env and calling /public/login
	// create login first
	login := map[string]string{"email": "admin@example.com", "password": "secret"}
	lb, _ := json.Marshal(login)
	lreq := httptest.NewRequest(http.MethodPost, "/public/login", bytes.NewBuffer(lb))
	lreq.Header.Set("Content-Type", "application/json")
	lrw := httptest.NewRecorder()
	r.ServeHTTP(lrw, lreq)
	var lresp map[string]interface{}
	_ = json.Unmarshal(lrw.Body.Bytes(), &lresp)
	token := "Bearer " + lresp["token"].(string)

	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("create got %d", w.Code)
	}

	// get list
	req = httptest.NewRequest(http.MethodGet, "/students/", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("list got %d", w.Code)
	}

	// get single
	req = httptest.NewRequest(http.MethodGet, "/students/1", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("get got %d", w.Code)
	}

	// update
	upd := map[string]string{"name": "B"}
	ub, _ := json.Marshal(upd)
	req = httptest.NewRequest(http.MethodPut, "/students/1", bytes.NewBuffer(ub))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("update got %d", w.Code)
	}

	// delete
	req = httptest.NewRequest(http.MethodDelete, "/students/1", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("delete got %d", w.Code)
	}
}
