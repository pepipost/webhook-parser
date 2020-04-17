package connection

import (
  "github.com/go-redis/redis"
  "webhook-parser/internal/config"
)

func ConnectRedis(host string, port string) (*redis.Client, error) {
  conn := redis.NewClient(&redis.Options{
    Addr:     host + ":" + port,
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  _, err := conn.Ping().Result()
  if err != nil {
    return nil, err
  }
  return conn, nil
}


