package strings

import(
    "rediproxy/base"
)

func (t *Strings)SET(key string, value string) (err error){
    c := base.RedisClient.Get()
    defer c.Close()

    _, err = c.Do("SET", key, value)
    return
}
