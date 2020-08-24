# Simple gRPC Server

## Run
Server:
```
docker run -p 9292:9292 -it gucvitaly/simple-grpc server
```
Client:
```
docker run --net=host -e SERVER_ADDR=localhost:9292 -it gucvitaly/simple-grpc client
```
