package handler

import (
	sqlrepo "github.com/vegggg/user-management/user/sqlrepo"
)

type UserManagement struct {
	sqlrepo *sqlrepo.SqlRepo
}

func NewUserManagement(sqlrepo *sqlrepo.SqlRepo) *UserManagement {
	u := new(UserManagement)
	u.sqlrepo = sqlrepo

	return u
}
