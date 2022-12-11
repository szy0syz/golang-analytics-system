# golang-analytics-system

> Kafka, Go, Postgres &amp; GraphQL

## Notes

```bash
go run -mod=mod github.com/99designs/gqlgen generate
```

<img width="875" alt="image" src="https://user-images.githubusercontent.com/10555820/206857554-990ff34e-238d-4faf-a8b1-40681bd7a9e9.png">

<img width="951" alt="image" src="https://user-images.githubusercontent.com/10555820/206883172-679c9af2-8e24-4c17-b108-8213fab0b736.png">

### Issue-1

> Consumer error: GroupCoordinator: localhost:9092: Connect to ipv6#[::1]:9092 failed: Connection refused (after 1ms in state CONNECT) (<nil>)
>

- solution

```go
c, err := kafka.NewConsumer(&kafka.ConfigMap{
  "bootstrap.servers":     "localhost",
  "group.id":              "myGroup",
  "auto.offset.reset":     "earliest",
  "broker.address.family": "v4",
 })
```

### Issue-2

> 无限打印 Local：Timed out 问题

```go
 for {
  msg, err := c.ReadMessage(time.Second)
  }
``

- 这里没消息来时，会不停打印 `Local: Timed out`
- 解决方案暂时改 `timeout: -1`
- 目前能力无法解决 😅

## Flow

- `Consumer`

<img width="871" alt="image" src="https://user-images.githubusercontent.com/10555820/206913030-5ae189ef-9d8e-42d7-b155-a9f2d586dbfb.png">

<img width="714" alt="image" src="https://user-images.githubusercontent.com/10555820/206913078-d2e61cdb-705f-4ac2-b3c8-a9fe40d3e790.png">
