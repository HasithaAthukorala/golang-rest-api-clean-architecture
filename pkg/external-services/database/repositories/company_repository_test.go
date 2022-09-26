package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"golang-rest-api-clean-architecture/pkg/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetCompanies(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("unexpected error while opening the db connection: %v", err)
	}
	repo := GetCompanyRepository(conn)

	company := &entities.Company{
		Name:    "TestName",
		Country: "TestCountry",
		Code:    "TestCode",
		Website: "http://test-site.com",
		Phone:   "+21349112123",
	}
	const addCompanyQuery = "INSERT INTO `companies` (`name`,`country`,`code`,`website`,`phone`) VALUES (?,?,?,?,?)"

	mock.ExpectBegin()
	mock.ExpectExec(addCompanyQuery).
		WithArgs(company.Name, company.Country, company.Code, company.Website, company.Phone).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = repo.AddCompany(company); err != nil {
		t.Errorf("failed to insert to the db: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("expectations were not met : %v", err)
	}
}
