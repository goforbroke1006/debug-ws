# debug-ws

Simple websocket server for debugging infra.

### How to use

```bash
docker run --name debug-ws-inst -it --rm -p 8080:8080 \
  docker.io/goforbroke1006/debug-ws:latest \
  --prefix=/some/path/
```

and now ws://127.0.0.1:8080/ws is waiting for connection

http://127.0.0.1:8080/ - simplest debug web page