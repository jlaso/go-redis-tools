package routing

import (
	"net/http"
	"github.com/jlaso/go-redis-tools/common"
)

func ServersHandler(w http.ResponseWriter, r *http.Request) {

	servers := common.GetServers()

	ret := JsonReturn{
		Success: true,
		Error: "",
		Data: servers,
	}

	ret.JsonWrite(w)
}
