# SSL Key generator

https://sslkey-generator.herokuapp.com/

## ContainerImage

https://hub.docker.com/r/youyo/sslkey-generator/

```
$ docker container run -d -p 1323:1323 -e PORT=1323 youyo/sslkey-generator:
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
}' http://your-site-address/generate
```
