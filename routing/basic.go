package routing

import (
	"net/http"
	"encoding/json"
)

type JsonReturn struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func (data *JsonReturn) JsonWrite(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := json.Marshal(data)
	if err != nil {
		data.Success = false
		data.Error = err.Error()
		data.Data = nil
		res, _ = json.Marshal(data)
	}
	w.Write(res)
	return err
}
