package sqlrepo

import "database/sql"

type SqlRepo struct {
	db   *sql.DB
	stmt stmt
}

type stmt struct {
	GetUser    string
	CreateUser string
}

var (
	// Sqlite ..
	Mysql stmt
	// Psql ..
	Psql stmt
)

func init() {
	Mysql = stmt{
		GetUser:    "SELECT * FROM users WHERE phone = ?",
		CreateUser: "INSERT INTO users (phone, first_name, last_name, email, country, province, city, address_line, postal_code) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
	}
	Psql = stmt{
		GetUser:    "SELECT id FROM users WHERE phone = $1",
		CreateUser: "INSERT INTO users (phone, first_name, last_name, email, country, province, city, address_line, postal_code) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
	}
}

func NewMysqlRepo(mysqldb *sql.DB) *SqlRepo {
	return &SqlRepo{
		db:   mysqldb,
		stmt: Mysql,
	}
}

func NewPsqlRepo(psqldb *sql.DB) *SqlRepo {
	return &SqlRepo{
		db:   psqldb,
		stmt: Psql,
	}
}
