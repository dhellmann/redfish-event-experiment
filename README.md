# Experimenting with Redfish Event Subscription

## Config files

Copy `config.yaml.example` to `config.yaml` and modify the BMC
settings and the IP address for the receiver endpoint.

## Data dumper

`dump/main.go` shows some settings for the BMC, including existing
event subscriptions.

```
go run ./dump/main.go
```

## Receiver

`receiver/main.go` is a web server that dumps the event notifications
to the console. To run it

1. First, use `make_certs.sh` to generate certificates (you may need
   to modify `openssl.conf` to change the hostname for the certs).
2. Then run

   ```
   go run ./receiver/main.go
   ```

## Register the receiver

`register/main.go` configures the BMC to send events to the receiver
service.

```
go run ./register/main.go
```

## Testing

After the receiver is registered, using the dumper should trigger a
login event notification. The output on the console will look
something like this.

```
$ go run ./receiver/main.go
listening on https://10.8.1.133:9090

[POST] /: "{\"Context\":\"Public\",\"EventId\":\"8491\",\"EventTimestamp\":\"2020-11-13T15:22:43-0500\",\"EventType\":\"Alert\",\"MemberId\":\"f35299c0-25ed-11eb-a058-588a5afa1cc4\",\"Message\":\"Successfully logged in using Username, from 10.8.1.133 and REDFISH.\",\"MessageArgs\":[\"Username\",\"10.8.1.133\",\"REDFISH\"],\"MessageArgs@odata.count\":3,\"MessageId\":\"USR0030\",\"OriginOfCondition\":\"iDRAC.Embedded.1\",\"Severity\":\"Informational\"}\n"
Accept: */*
Content-Length: 394
Content-Type: application/json
```

The web UI for the BMC also includes a way to send test events. The
input must be a valid message ID, as defined by the hardware. One
value that works for Dell is `CPU0001`.

## Cleaning up

`unsubscribe/main.go` removes a subscription from the BMC. Use the
dumper program to get the URI, then pass it on the command line.

```
go run ./unsubscribe/main.go /redfish/v1/EventService/Subscriptions/f35299c0-25ed-11eb-a058-588a5afa1cc4
```
