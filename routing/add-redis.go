package routing

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/jlaso/go-redis-tools/common"
)

func AddRedisPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	port := r.Form.Get("port")
	cfg := common.GetConfiguration()
	cfg.Port, err = strconv.Atoi(port)
	if err != nil {
		fmt.Println(err)
	}
	cfg.Host = r.Form.Get("host")
	err = cfg.Save()
	if err != nil {
		fmt.Printf("Error saving config file with %s:%d\n", cfg.Host, cfg.Port)
		w.Write([]byte(fmt.Sprintf("Error saving config file with %s:%d\n", cfg.Host, cfg.Port)))
	}else {
		fmt.Printf("%s:%d\n", cfg.Host, cfg.Port)
		w.Write([]byte(fmt.Sprintf("%s:%d\n", cfg.Host, cfg.Port)))
	}
}