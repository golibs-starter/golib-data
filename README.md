# Golib Data

Database solutions for Golang project. Includes:

* MySQL database
* Postgres database
* Sqlite database
* Redis database

### Setup instruction

Base setup, see [GoLib Instruction](https://gitlab.com/golibs-starter/golib/-/blob/main/README.md)

Both `go get` and `go mod` are supported.

```shell
go get gitlab.com/golibs-starter/golib-data
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
    "database/sql"
    "github.com/redis/go-redis/v9"
    "gitlab.com/golibs-starter/golib-data"
    "gitlab.com/golibs-starter/golib-data/testutil"
    "go.uber.org/fx"
    "gorm.io/gorm"
)

func main() {
    fx.New(
        // When you want to use redis
        golibdata.RedisOpt(),

        // When you want to use datasource
        golibdata.DatasourceOpt(),

        // Demo way to using redis
        fx.Provide(funcUseRedis),
        fx.Provide(funcUseOrm),
        fx.Provide(funcUseNativeDbConnection),

        // ==================== TEST UTILS =================
        // A useful util to easy to interact with database in test.
        golibdataTestUtil.EnableDatabaseTestUtilOpt(),
        golibdataTestUtil.EnableRedisTestUtilOpt(),

        // This useful when you want to truncate some tables before test.
        // Eg: https://gitlab.com/golibs-starter/golib-sample/-/tree/develop/src/public/testing/create_order_controller_test.go
        golibdataTestUtil.TruncateTablesOpt("table1", "table2"),
    )
}

func funcUseRedis(redisClient *redis.Client) {
    // do something with redis client
}

func funcUseOrm(db *gorm.DB) {
    // do something with gorm
}

func funcUseNativeDbConnection(db *sql.DB) {
    // do something with the native database connection
}
```

### Configuration

```yaml
app:
    # Configuration available for golib.DatasourceOpt()
    datasource:
        # SQL driver. Supports: mysql, postgres, sqlite
        driver: mysql

        # Define the database host
        host: localhost

        # Define the database port
        port: 3306

        # Define the database name
        database: sample

        # Define the username for authentication
        username: root

        # Define the password for authentication
        password: secret

        # Extra params to add to the connection string
        params: parseTime=true

        # When dsn is provided, it will override all above connection configs
        dsn: user1@tcp(127.0.0.1:3306)/demo

        # The maximum number of open connections to the database.
        # Default 10 connections
        maxOpenConns: 10

        # The maximum amount of time a connection may be reused.
        # Default 30m
        connMaxLifetime: 30m

        # The maximum number of connections in the idle connection pool.
        # Default 2
        maxIdleConns: 2

        # The maximum amount of time a connection may be idle.
        # Default 10m
        connMaxIdleTime: 10m

        # The query log level.
        # There are four levels: "SILENT", "ERROR", "WARN", "INFO".
        # Default SILENT
        logLevel: SILENT

    # Configuration available for golib.RedisOpt()
    redis:
        # Define the redis host. Default: `localhost`
        host: localhost

        # Define the redis port. Default: `6379`
        port: 6379

        # Define the redis database. Default: `0`
        database: 0

        # Define the redis username
        username: ""

        # Define the redis password
        password: ""

        # Enable or disable TLS
        enableTLS: true

        # Maximum number of socket connections.
        # Default 10 connections
        poolSize: 10

        # Connection age at which client retires (closes) the connection.
        # Default is Zero means not close aged connections.
        maxConnAge: 0

        # Minimum number of idle connections which is useful when establishing
        # new connection is slow.
        # Default 2 idle connections
        minIdleConns: 2

        # Amount of time after which client closes idle connections.
        # Should be less than server's timeout.
        # Default is 5 minutes. -1 disables idle timeout check.
        idleTimeout: 5m
```
