# etcd_service_discovery
etcd service discovery example

Use etcd for service discovery and auto remove outdated workers.

## run server

```bash
# run server
go run server/main.go

# run worker to register
go run worker/main.go
```

## result

After 5 seconds, worker register information will removed from etcd, the server will receive the notification.

```bash
➜  server git:(master) ✗ go run main.go
2018-11-30 18:51:40.871320 I | PUT "service/worker1" : "here is worker1"
2018-11-30 18:51:44.587254 I | PUT "service/worker1" : "here is worker1"
2018-11-30 18:51:49.929817 I | DELETE "service/worker1" : ""
```