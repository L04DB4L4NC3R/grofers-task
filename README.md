### Grofers task
grof CLI

<br/>

[![](https://img.shields.io/badge/docs%20-view%20API%20documentation-orange.svg?style=for-the-badge&logo=appveyor)](https://angadsharma1016.github.io/grofers-task/)


[![demo](https://img.shields.io/badge/view%20demo-youtube-orange.svg?style=for-the-badge&logo=appveyor)]()

<br/>
<br/>



<details>

<summary>Technology stack used</summary>

<br/>

- [X] Golang + MySQL
- [X] NATS as a high availability messaging queue service
- [X] docker + docker-compose
- [X] bash scripts
- [X] apidoc documentation

### Why NATS?
---
[NATS](https://github.com/nats-io/go-nats.git) is an event sourcing tool which we will be using to publish logs and distribute related data between different services. The reason for NATS is:

* RabbitMQ is limited to HTTP and HTTPS natively

* NATS supports gRPC and works fast due to being type safe as well as removing extra overhead of marshalling and unmarshalling

* Publishing events on a different thread and subscribing from a different thread allows non-blocking log sourcing.

* NATS is very high throughput when it comes to requests per second

<br />

![NATS](./images/nats.png)

<br />
</details>

<br/>
<br/>


### How to setup

```
$ git clone https://github.com/angadsharma1016/grofers-task.git && cd grofers-task

$ chmod +x ./bin/*

$ ./bin/setup
```

Wait for config after ./bin/setup, you can use `docker-compose logs -f` to monitor changes.

<br/>
<br/>

### How to run

```
$ grof
```

After running `grof` you will see something like this:


```
================================= grof CLI ==============================


Usage: 
grof put [key] [value] --------------------> Set key-value
grof get [key] ----------------------------> Get value from key
grof delete [key] -------------------------> Delete from DB
grof update [key] [new value] -------------> Update DB with new value
grof watch --------------------------------> Subscribe to changes in DB
grof watch --------------------------------> Watch for realtime DB changes
```

<br/>
<br/>

### How to build from source
If you want to build from source simple do the following

```
go build -o ./bin/grof ./main.go 
sudo cp ./bin/grof /use/local/bin/grof
```