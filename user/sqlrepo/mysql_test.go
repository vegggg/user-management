package sqlrepo

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vegggg/user-management/entity"
)

var mysqlDB *sql.DB
var mysqlStmt stmt

func newMysqlConn() (*sql.DB, error) {
	address := "user:pass@tcp(localhost:33061)/user_management?parseTime=true&loc=Local"
	return sql.Open("mysql", address)
}
func init() {

	db, err := newMysqlConn()
	if err != nil {
		panic(err)
	}
	s := &SqlRepo{
		db:   db,
		stmt: Mysql,
	}
	s.CreateUser(context.Background(), &entity.UserProfile{
		Phone:       "0944146503",
		FirstName:   "Nhat Dinh",
		LastName:    "Nguyen",
		Email:       "ngnhatdinh1110@gmail.com",
		Country:     "VietNam",
		Province:    "abc",
		City:        "Ho Chi Minh",
		AddressLine: "Ho Chi Minh",
		PostalCode:  "700000",
	})
	mysqlDB = db
}
func TestSqlRepo_GetUser(t *testing.T) {
	type fields struct {
		db   *sql.DB
		stmt stmt
	}

	mysqlStmt := Mysql

	type args struct {
		ctx   context.Context
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.UserProfile
		wantErr bool
	}{
		{
			"happy case",
			fields{
				mysqlDB,
				mysqlStmt,
			},
			args{
				context.Background(),
				"0944146504",
			},
			&entity.UserProfile{
				Phone:       "0944146504",
				FirstName:   "Nhat Dinh",
				LastName:    "Nguyen",
				Email:       "ngnhatdinh1110@gmail.com",
				Country:     "VietNam",
				Province:    "abc",
				City:        "Ho Chi Minh",
				AddressLine: "Ho Chi Minh",
				PostalCode:  "700000",
			},
			false,
		},
		{
			"Not found",
			fields{
				mysqlDB,
				mysqlStmt,
			},
			args{
				context.Background(),
				"094414650",
			},
			&entity.UserProfile{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SqlRepo{
				db:   tt.fields.db,
				stmt: tt.fields.stmt,
			}
			got, err := s.GetUser(tt.args.ctx, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqlRepo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SqlRepo.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqlRepo_CreateUser(t *testing.T) {

	mysqlStmt := Mysql

	type fields struct {
		db   *sql.DB
		stmt stmt
	}
	type args struct {
		ctx context.Context
		u   *entity.UserProfile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.UserProfile
		wantErr bool
	}{
		{
			"duplicated phone case",
			fields{
				mysqlDB,
				mysqlStmt,
			},
			args{
				context.Background(),
				&entity.UserProfile{
					Phone:       "0944146503",
					FirstName:   "Nhat Dinh",
					LastName:    "Nguyen",
					Email:       "ngnhatdinh1110@gmail.com",
					Country:     "VietNam",
					Province:    "abc",
					City:        "700000",
					AddressLine: "Ho Chi Minh",
					PostalCode:  "Ho Chi Minh",
				},
			},
			nil,
			true,
		},
		{
			"happy case",
			fields{
				mysqlDB,
				mysqlStmt,
			},
			args{
				context.Background(),
				&entity.UserProfile{
					Phone:       "0944146506",
					FirstName:   "Nhat Dinh",
					LastName:    "Nguyen",
					Email:       "ngnhatdinh1110@gmail.com",
					Country:     "VietNam",
					Province:    "abc",
					City:        "700000",
					AddressLine: "Ho Chi Minh",
					PostalCode:  "Ho Chi Minh",
				},
			},
			&entity.UserProfile{
				Phone:       "0944146506",
				FirstName:   "Nhat Dinh",
				LastName:    "Nguyen",
				Email:       "ngnhatdinh1110@gmail.com",
				Country:     "VietNam",
				Province:    "abc",
				City:        "700000",
				AddressLine: "Ho Chi Minh",
				PostalCode:  "Ho Chi Minh",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SqlRepo{
				db:   tt.fields.db,
				stmt: tt.fields.stmt,
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqlRepo.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SqlRepo.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
