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

## Cleaning up

`unsubscribe/main.go` removes a subscription from the BMC. Use the
dumper program to get the URI, then pass it on the command line.

```
go run ./unsubscribe/main.go /redfish/v1/EventService/Subscriptions/f35299c0-25ed-11eb-a058-588a5afa1cc4
```
