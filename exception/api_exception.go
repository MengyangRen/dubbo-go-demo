package exception

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 定义一个基础APi异常类
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package exception
 * @description
 *
 *
 */

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	EX_NOT_MODIFIED         = 304
	EX_REQUEST              = 400
	EX_REQUEST_UNAUTHORIZED = 401
	EX_SERVER               = 500
	EX_SERVER_NOTFOUND      = 501
	EX_REQUEST_INVAL        = 601
	EX_RESP_EMPTY           = 602

	//user defined Exception must be bigger than 700
	EX_USER_TOKEN_NOTEXIST = 714
)

var (
	Messages = map[int]string{
		EX_NOT_MODIFIED:         "Not Modified",
		EX_REQUEST:              "Bad request",
		EX_REQUEST_UNAUTHORIZED: "Unauthorized",
		EX_SERVER:               "Internal Server Error",
		EX_SERVER_NOTFOUND:      "Service not found",
		EX_REQUEST_INVAL:        "Invalid argument",
		EX_RESP_EMPTY:           "Result empty",
		EX_USER_TOKEN_NOTEXIST:  "The user token is invalid",
	}
)

type ApiException struct {
	messages map[int]string
}

func NewApiException() *ApiException {
	apiEx := new(ApiException)
	apiEx.messages = make(map[int]string)
	apiEx.messages = Messages
	return apiEx
}
func (this *ApiException) Throw(code int, msg string) error {
	_s := fmt.Sprintf("%d", code) + "_" + msg
	panic(errors.New(_s))
}

func (this *ApiException) GetCode(e string) int {
	code, _ := strconv.ParseInt(
		strings.Split(e, "_")[0],
		10,
		64,
	)
	return int(code)
}

func (this *ApiException) GetMeessage(e string) string {
	return strings.Split(e, "_")[1]
}
