# logstash

Simple formatter for Logstash

## Example

```go
conn, err := net.Dial("tcp", "localhost:3333")
if err != nil {
    log.Fatal(err)
}

log.SetOutput(logstash.New(conn, map[string]interface{}{
    "tags": []string{"vision-api"},
}))

log.Println("Hello!")
```
