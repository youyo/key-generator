# SSL Key generator

## Website

https://sslkey-generator.arukascloud.io/

## ContainerImage

https://hub.docker.com/r/youyo/sslkey-generator/

```
$ docker run -d -p 1323:1323 youyo/sslkey-generator:v5
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
}' https://sslkey-generator.arukascloud.io/generate
```
