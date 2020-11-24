module dev-dubbo-producer

require (
	github.com/apache/dubbo-go v1.5.4
	github.com/apache/dubbo-go-hessian2 v1.7.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dubbogo/gost v1.9.2
	github.com/emicklei/go-restful/v3 v3.0.0
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922
	google.golang.org/grpc v1.26.0
)

go 1.14

replace gitlab.stagingvip.net/publicGroup/public v0.0.0-20201029110713-ec938beba922 => D:/workspace/tcp/go-env/src/gitlab.stagingvip.net/publicGroup/public
