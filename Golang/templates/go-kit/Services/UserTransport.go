package Services

import (
	"_examples/go-kit/util"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid:    uid,
			Method: r.Method,
		}, nil
	}
	return nil, errors.New("参数错误")

}

func EncodeUserResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

/**
error 错误信息处理
*/
func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter){
	contextType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("content-type", contextType)
	if baseError, ok := err.(*util.BaseError); ok {
		w.WriteHeader(baseError.Code)
	}else{
		w.WriteHeader(500)
	}
	_, _ = w.Write(body)
}