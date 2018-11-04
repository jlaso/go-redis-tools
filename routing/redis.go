package routing

import (
	"github.com/go-redis/redis"
	"fmt"
	"strings"
	"net/http"
	"github.com/jlaso/go-redis-tools/common"
)

func RedisHandler(w http.ResponseWriter, r *http.Request) {

	cfg := common.GetConfiguration()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res := "["

	keys, err := client.Keys("*").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, key := range keys {
		v, err := client.Get(key).Result()
		if err != nil {
			fmt.Println(err.Error())
		}
		res += fmt.Sprintf(`{"key":"%s","value":"%s"},`, key, v)
	}

	res = strings.Trim(res, ",")

	res += "]"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write([]byte(res))
}
