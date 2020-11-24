package dnc

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 用户消费者结构体
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package controllers
 * @description
 *
 * 说明:
 *
 */
import (
	"context"
	"time"

	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"

	_ "github.com/apache/dubbo-go/protocol/dubbo"

	_ "github.com/apache/dubbo-go/registry/protocol"

	_ "github.com/apache/dubbo-go/filter/filter_impl"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"

	_ "github.com/apache/dubbo-go/cluster/loadbalance"

	_ "github.com/apache/dubbo-go/registry/nacos"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Age       int32     `json:"age"`
	Addr      string    `json:"address"`
	Phone     string    `json:"phone"`
	Introduce string    `json:"introduce"`
	Time      time.Time `json:"requestTime "`
}

type UserPaginationQ struct {
	Users []*User `json:"users"`
	Page  int64   `json:"page"`
	Limit int64   `json:"limit"`
	Total int     `json:"total"`
}

type UserStoreResult struct {
	InsterID int `json:"_insterID"`
	Rows     int `json:"_rows"`
}

type UserProvider struct {
	GetUser func(ctx context.Context, req []interface{}, rsp *User) error
	Users   func(ctx context.Context, req []interface{}, rsp *UserPaginationQ) error
	Update  func(ctx context.Context, req []interface{}, rsp *User) error
	Destroy func(ctx context.Context, req []interface{}, rsp *User) error
	Store   func(ctx context.Context, req []interface{}, rsp *UserStoreResult) error
}

func (u *UserProvider) Reference() string {
	return "UserProvider"
}

func (User) JavaClassName() string {
	return "com.ikurento.user.User"
}

func (UserPaginationQ) JavaClassName() string {
	return "com.ikurento.user.UserPagination"
}

func (up *UserStoreResult) JavaClassName() string {
	return "com.ikurento.user.UserStoreResult"
}
