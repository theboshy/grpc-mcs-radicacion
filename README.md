
# GRPC Kinesthetic {radicacion}
Implementacion de servicio / api rest y microservicio grpc integrado al orm (xorm) / grpc y rest end point

Ejecucion  (o cambios agregados) agregados al protobuf

 `#app/mcs/grpc/proto/radicacionmcs.proto build`

```console
~$ protoc  ./app/mcs/grpc/proto/radicacionmcs.proto --go_out=plugins=grpc:./
```

## Librerias utilizadas 

* [golang.org/x/net/context](https://godoc.org/golang.org/x/net/context)
* [google.golang.org/grpc](https://godoc.org/google.golang.org/grpc) 
* [google.golang.org/grpc/reflection](https://godoc.org/google.golang.org/grpc/reflection)
* [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
* [github.com/gin-contrib/cors](https://github.com/gin-contrib/cors)
* [github.com/gorilla/mux](https://github.com/gorilla/mux)
* [github.com/gorilla/handlers](https://github.com/gorilla/handlers)
* [github.com/go-xorm/core](https://github.com/go-xorm/core)
* [github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)
* [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

## Instalacion 

```console
~$ go get -d .
```

## Configuracion DB

> {raiz}/configuration.json

### Test conexion
> {raiz}/main.go
