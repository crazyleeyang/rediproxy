package strings

import(
    "rediproxy/base"
)

func (t *Strings) DEL(key string) (r int32, err error){
    c := base.RedisClient.Get()
    defer c.Close()

    _, err = c.Do("DEL", key)
    return
}

