package redis

import (
    "fmt"
    "github.com/Unknwon/goconfig"
    "github.com/garyburd/redigo/redis"
)

type Link struct {
    Conn redis.Conn
}

func (l *Link) Init() {
    var err error
    configPath := "./config.ini"
    config, load_conf_err := goconfig.LoadConfigFile(configPath)
    if load_conf_err != nil {
        fmt.Println(load_conf_err)
    }
    host, _ := config.GetValue("redis", "host")
    l.Conn, err = redis.Dial("tcp", host)

    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }

    //defer l.Conn.Close()
}

func (l *Link) Get(key string) string {
    value, err := redis.String(l.Conn.Do("GET", key))
    if err != nil {
        fmt.Println("err while getting:", err)
    }

    return value
}

func (l *Link) BrPop(key ...interface{}) string {
    value, err := redis.Strings(l.Conn.Do("brpop", key...))

    if err != nil {
        fmt.Printf("err while getting:", err)
    }
    return value[1]
}
