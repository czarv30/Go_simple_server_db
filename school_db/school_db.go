package school_db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type SchoolDb struct {
	db *sql.DB
}

type student struct {
	StudentID   int       `json:"student_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func InitSchoolDb(dsn string) (*SchoolDb, error) { // name "dsn" -- data source name
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &SchoolDb{db: db}, nil
}

func (repo *SchoolDb) Close() error {
	return repo.db.Close()
}

func (repo *SchoolDb) GetAllStudents() ([]student, error) {
	rows, err := repo.db.Query("SELECT student_id, first_name, last_name, birth_date FROM students")

	var students []student // A student "slice," similar to python list.

	for rows.Next() {
		var s student
		rows.Scan(&s.StudentID, &s.FirstName, &s.LastName, &s.DateOfBirth)
		students = append(students, s)
	}

	return students, err
}

func (repo *SchoolDb) PostStudent(s student) error {
	_, err := repo.db.Exec("INSERT INTO students (first_name, last_name, birth_date) VALUES ($1, $2, $3)", s.FirstName, s.LastName, s.DateOfBirth)
	return err
}
