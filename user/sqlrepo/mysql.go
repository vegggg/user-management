package sqlrepo

import (
	"context"

	"github.com/golang/glog"
	"github.com/vegggg/user-management/entity"
)

func (s *SqlRepo) GetUser(ctx context.Context, phone string) (*entity.UserProfile, error) {
	stmt := s.stmt.GetUser

	glog.Infoln("get user: ", phone)
	rows, err := s.db.QueryContext(ctx, stmt, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(entity.UserProfile)
	for rows.Next() {
		err = rows.Scan(u.Phone, &u.Phone, &u.FirstName, &u.LastName, &u.Country, &u.Province, &u.City, &u.AddressLine, &u.PostalCode)
		if err != nil {
			return nil, err
		}
		break // break here to ensure only one user is scanned
	}

	return u, nil
}
func (s *SqlRepo) CreateUser(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error) {
	stmt := s.stmt.CreateUser

	glog.Infoln("create user: ", u)
	// (phone, first_name, last_name, email, country, province, city, address_line, postal_code)
	_, err := s.db.ExecContext(ctx, stmt, &u.Phone, &u.FirstName, &u.LastName, &u.Country, &u.Province, &u.City, &u.AddressLine, &u.PostalCode)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return u, nil
}
