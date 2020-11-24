package service

import (
	"cloudlive/exception"
	dnc "cloudlive/registry/dnc/app"
	"context"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 用户服务层
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package service
 * @description
 *
 * 说明:
 *
 */

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) User(param map[string]string) *dnc.User {
	user := &dnc.User{}
	dnc.UserPvder.GetUser(context.TODO(), []interface{}{param["id"]}, user)
	return user
}

func (this *UserService) Users(param map[string]string) *dnc.UserPaginationQ {
	upq := &dnc.UserPaginationQ{}
	dnc.UserPvder.Users(context.TODO(), []interface{}{param["page"], param["limit"]}, upq)
	return upq
}

func (this *UserService) Update(param map[string]string) {
	if err := dnc.UserPvder.Update(context.TODO(), []interface{}{
		param["id"], param["name"], param["addr"]}, &dnc.User{}); err != nil {
		exception.NewApiException().Throw(
			756,
			"User information update failed",
		)
	}
}

func (this *UserService) Destroy(param map[string]string) {
	if err := dnc.UserPvder.Destroy(context.TODO(), []interface{}{
		param["id"]}, &dnc.User{}); err != nil {
		exception.NewApiException().Throw(
			776,
			"User destruction failure",
		)
	}
}

func (this *UserService) Store(param []interface{}) *dnc.UserStoreResult {
	usr := &dnc.UserStoreResult{}
	if err := dnc.UserPvder.Store(context.TODO(), param, usr); err != nil {
		exception.NewApiException().Throw(
			779,
			"User creation failure",
		)
	}
	return usr
}
