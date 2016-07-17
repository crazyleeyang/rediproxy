package strings

import(
    "github.com/garyburd/redigo/redis"
    "rediproxy/base"
)

func (t *Strings) GET(key string) (r string, err error){
    c := base.RedisClient.Get()
    defer c.Close()

    r, err = redis.String(c.Do("GET", key))
    return
}
