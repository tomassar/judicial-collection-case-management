package cases

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestPostCase_success(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	caseRepo := NewCaseRepository(db)
	caseService := NewCaseService(caseRepo)
	caseController := NewCaseRoutes(caseService)

	responseRecorder := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(responseRecorder)
	engine.POST("/cases", caseController.createCase)

	// Expect the Begin transaction
	mock.ExpectBegin()

	// Expect the query operation
	expectedSQL := "INSERT INTO \"cases\" (.+) VALUES (.+)"
	mock.ExpectQuery(expectedSQL).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Victoria Contreras", 500000, "active").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Expect the Commit transaction
	mock.ExpectCommit()

	requestBody := `{"debtor_name": "Victoria Contreras", "amount": 500000, "status": "active"}`
	ctx.Request = httptest.NewRequest(http.MethodPost, "/cases", bytes.NewBuffer([]byte(requestBody)))
	engine.ServeHTTP(responseRecorder, ctx.Request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}
