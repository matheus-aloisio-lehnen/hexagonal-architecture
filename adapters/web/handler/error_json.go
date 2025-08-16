package handler

import "encoding/json"

func jsonError(msg string) []byte {
	erro := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	r, err := json.Marshal(erro)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
