package controllers

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 用户控制层
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
	"cloudlive/filter"
	"cloudlive/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Base *BaseController
}

func (this *UserController) User(c *gin.Context) {
	//异常处理
	defer this.Base.Catch(NewResponse(c))
	//参数校验与过滤层
	param := filter.NewUserFilter(c).User()
	//服务调用层与响应
	NewResponse(c).Success(map[string]interface{}{
		"controllers.user": service.NewUserService().User(param)}).Json()

}

func (this *UserController) Users(c *gin.Context) {
	param := filter.NewUserFilter(c).Users()
	NewResponse(c).Success(map[string]interface{}{
		"controllers.users": service.NewUserService().Users(param)}).Json()
}

func (this *UserController) Update(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	param := filter.NewUserFilter(c).Update()
	service.NewUserService().Update(param)
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

func (this *UserController) Destroy(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	param := filter.NewUserFilter(c).User()
	service.NewUserService().Destroy(param)
	NewResponse(c).Success(map[string]interface{}{}).Json()
}

func (this *UserController) Store(c *gin.Context) {
	defer this.Base.Catch(NewResponse(c))
	param := filter.NewUserFilter(c).Store()
	NewResponse(c).Success(map[string]interface{}{
		"controllers.store": service.NewUserService().Store(param)}).Json()
}
