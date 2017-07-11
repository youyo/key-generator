# SSL Key generator

## ContainerImage

https://hub.docker.com/r/youyo/sslkey-generator/

```
$ docker run -d -p 1323:1323 youyo/sslkey-generator:latest
```

## API

```
$ curl -X POST -H "Content-Type: application/json" -d '{
  "common_name": "test.com",
  "country": "JP",
  "state": "Miyagi",
  "locality": "Sendai",
  "organization_name": "example inc.",
  "organizational_unit_name": "sales"
}' http://your-site-address:1323/generate
```
