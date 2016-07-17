namespace go redis.String

service stringRpc {
    string GET(1:string key),
    void SET(1:string key, 2:string value),
    i32 DEL(1:string key),
}

