# Viper Nacos Config Example

> Convention over Configuration

An out-of-the-box example of using [Nacos](https://github.com/alibaba/nacos) configuration management with [Viper](https://github.com/spf13/viper).

## Quick Start

[中文](https://juejin.cn/post/7194335882232397882)

### Remote

[example](./remote)

#### Start Nacos

```shell
docker-compose up
```

#### Config Nacos

Visit `127.0.0.1:8848/nacos` on Browser.

The default username and password is `nacos`.

Add MySQL configuration

- Data ID: `example-remote`
- Group:   `example-remote`

```yaml
mysql:
    host: 127.0.0.1
    port: 3306
    username: root
    password: 114514
    database: example-remote
```

#### Run

```shell
cd remote
go run .
```

### Local

[example](./local)

#### Start Nacos

```shell
docker-compose up
```

#### Run

```shell
cd local
go run .
```

#### Visit Nacos

Visit `127.0.0.1:8848/nacos` on Browser.

The default username and password is `nacos`.

## End

Just an Example