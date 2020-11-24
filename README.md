# dubbo-go-demo

### 简介
  dubbo-go-demo 解决团队快速了解与使用的问题<br>
  当初所在团队需要解决Go团队与Java团队dubbo服务互通的问题，那么dubbo-go成为团队首选。

### 环境需求
* \>= go 1.14
* \>= dubbo-go v1.5.4
* \>= dubbo-go-hessian2 v1.7.0
* \>= gin v1.6.3
* nacos/nacos-server

### 分支说明
```code
 分支: gin-dubbo-go-consumer（消费端） 
 分支: dubbo-go-producer     (生产端)
```
### 组成说明
  * 1.Docker-compose搭建nacos-server
  * 2.Dubbo-go消费端与Gin框架进行整合与代码分层架构设计（适合中型业务型服务） 
  * 3.Dubbo-go生产端代码分层架构设计
  * 4.Dubbo-go-demo提供用户模块Demo(消费端至生产端完整的CURD)
  * 5.Dubbo-go-demo提供分布式事务Demo

### 消费端代码架构流程
![example-1](https://github.com/MengyangRen/dubbo-go-demo/blob/gin-dubbo-go-consumer/doc/example-01.jpg)
### 生产端代码架构流程
![example-2](https://github.com/MengyangRen/dubbo-go-demo/blob/gin-dubbo-go-consumer/doc/example-02.jpg)

### 其他/思考
  <font color="#dd0000">以上代码架构流程图(排除注册中心)</font><br>
  对于刚入Golang坑普通微服务开发者<br>
  熟悉Dubbo-go的过程中也发现它的易用性的问题，它的入门成本较高,把很多对它感兴趣团队挡在了门外<br>
  另外团队内部推广也成了很大的问题，主要为资料少，成熟案例较少 <br>
  为此编写Dubbo-go-demo能快速让团队内部能正确使用与快速熟悉  <br>

  [基础调试环境配置](https://github.com/MengyangRen/dubbo-go-demo/blob/gin-dubbo-go-consumer/doc/example.md)


