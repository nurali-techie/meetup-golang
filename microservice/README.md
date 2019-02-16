**Run User and Role Services**

first, change dir to the target version (v1, v2, v3, v4) and then run below command to start user and role service.

- v1 - monolothic (user service in "go")
```
go run usersvc/usersvc1.go
```

- v2 - rpc (user service in "go" and role service in "go")
```
go run usersvc/usersvc2.go
go run rolesvc/rolesvc2.go
```

- v3 - grpc (user service in "go" and role service in "go")
```
protoc -I model/ model/model3.proto --go_out=plugins=grpc:model
go run usersvc/usersvc3.go
go run rolesvc/rolesvc3.go
```

- v4 - grpc (user service in "go" and role service in "python")

```
protoc -I model/ model/model4.proto --go_out=plugins=grpc:model
python -m grpc_tools.protoc -I model/ --python_out=. --grpc_python_out=. model/model4.proto
go run usersvc/usersvc4.go
python rolesvc4.py
```

**curl client**

```
curl http://localhost:8080/users?id=1
curl http://localhost:8080/users?id=2
```

**references:**

https://golang.org/pkg/net/rpc/

https://grpc.io/docs/quickstart/go.html

https://github.com/grpc/grpc-go/tree/master/examples/helloworld

https://grpc.io/docs/quickstart/python.html

https://github.com/grpc/grpc/tree/master/examples/python/helloworld

https://www.youtube.com/watch?v=FK9xCK0JoTw

https://youtu.be/FK9xCK0JoTw?t=1027
