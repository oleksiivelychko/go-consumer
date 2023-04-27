# go-consumer

### Independent microservice consumes messages from third-party channel.

⚠️ Third-party producer [go-queue-service](https://github.com/oleksiivelychko/go-queue-service) must be running before.

📌 Build binary:
```
go build -o bin/app -v .
```

📌 Start consumer:
```
go build -o bin/app -v . && MQ_USER=rabbit MQ_PASS=secret MQ_HOST=go-queue-service.local MQ_PORT=5672 MQ_NAME=go-queue ./bin/app
```

![Send and receive message from queue](social_preview.png)
