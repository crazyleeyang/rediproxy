package base

import(
    "time"
    "os"
    "fmt"
    "github.com/garyburd/redigo/redis"
    "github.com/Unknwon/goconfig"
)

var(
    RedisClient     *redis.Pool
)

func init() {
    cfg, err := goconfig.LoadConfigFile("conf/rediproxy.ini")
    if err != nil {
        return
    }

    redisIP, err := cfg.GetValue("redis", "redisIP")
    if err != nil {
        return
    }

    redisPort, err := cfg.GetValue("redis", "redisPort")
    if err != nil {
        return
    }

    redisMaxidle, err := cfg.Int("redis", "redisMaxidle")
    if err != nil {
        return
    }

    redisMaxactive, err := cfg.Int("redis", "redisMaxactive")
    if err != nil {
        return
    }

    redisIdleTimeout, err := cfg.Int("redis", "redisIdleTimeout")
    if err != nil {
        return
    }

    redisDb, err := cfg.Int("redis", "redisDb")
    if err != nil {
        return
    }

    RedisClient = &redis.Pool{
        MaxIdle: redisMaxidle,
        MaxActive: redisMaxactive,
        IdleTimeout: time.Duration( redisIdleTimeout ) * time.Second,

        Dial: func() (redis.Conn, error) {
            connStr := redisIP + ":" + redisPort
            c, err := redis.Dial("tcp", connStr)
            if err != nil {
                fmt.Println("redis.Dial failed:")
                os.Exit(1)
            }
            // 选择db
            c.Do("SELECT", redisDb)
            return c, nil
        },
    }
}

