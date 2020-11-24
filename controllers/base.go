package controllers

import (
	"fmt"
	"gin-dubbogo-consumer/exception"
)

type BaseController struct{}

func (this *BaseController) Catch(response *Response) {
	if err := recover(); err != nil {
		_e := fmt.Sprintf("%s", err)
		response.Fail(exception.NewApiException().GetCode(_e),
			exception.NewApiException().GetMeessage(_e)).Json()
		return
	}
}
