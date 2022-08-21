package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *AppHandler) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 8 << 20 // 8 MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// 再次解码，验证是否单个json, (防止：{}{} 出现， Decode 每次只解析一个json)
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("读取JSON错误: 仅仅允许传输单个JSON值")
	}

	return nil
}

func (app *AppHandler) errorResponse(w http.ResponseWriter, msg ...string) {
	w.WriteHeader(http.StatusInternalServerError)
	message := "服务内部错误"
	if len(msg) > 0 {
		message = msg[0]
	}
	w.Write([]byte(message))

}
