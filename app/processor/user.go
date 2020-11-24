package processor

import (
	"context"
	"dev-dubbo-producer/app/service"

	"time"
)

type User struct {
	Id        string
	Name      string
	City      string
	Age       int32
	Addr      string
	Phone     string
	Introduce string
	Time      time.Time
}

type UserPaginationQ struct {
	Users []*User
	Page  int64
	Limit int64
	Total int
}

type UserStoreResult struct {
	InsterID int
	Rows     int
}

type UserProvider struct{}

func (u *UserProvider) Users(ctx context.Context, req []interface{}) (*UserPaginationQ, error) {
	page, limit, count, uMaps := service.NewUserService().Users(req)
	return &UserPaginationQ{Users: NewUserWrap().BatchTransForm(uMaps), Page: page, Limit: limit, Total: count}, nil
}

func (u *UserProvider) GetUser(ctx context.Context, req []interface{}) (*User, error) {
	return NewUserWrap().TransForm(service.NewUserService().User(req)), nil
}

func (u *UserProvider) Update(ctx context.Context, req []interface{}) (*User, error) {
	if err := service.NewUserService().Update(req); err != nil {
		return &User{}, err
	}
	return &User{}, nil
}
func (u *UserProvider) Destroy(ctx context.Context, req []interface{}) (*User, error) {
	if err := service.NewUserService().Destroy(req); err != nil {
		return &User{}, err
	}
	return &User{}, nil
}

func (u *UserProvider) Store(ctx context.Context, req []interface{}) (*UserStoreResult, error) {
	_insterID, _rows, err := service.NewUserService().Store(req)
	if err != nil {
		return nil, err
	}
	return &UserStoreResult{InsterID: _insterID, Rows: _rows}, nil
}

func (u *UserProvider) Reference() string {
	return "UserProvider"
}

func (u User) JavaClassName() string {
	return "com.ikurento.user.User"
}

func (up *UserPaginationQ) JavaClassName() string {
	return "com.ikurento.user.UserPagination"
}

func (up *UserStoreResult) JavaClassName() string {
	return "com.ikurento.user.UserStoreResult"
}