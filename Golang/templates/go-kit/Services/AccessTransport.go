package Services

import (
	"context"
	"errors"
	"github.com/tidwall/gjson"
	"net/http"
	"io/ioutil"
)

func DecodeAccessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	result := gjson.Parse(string(body))
	if result.IsObject() {
		username := result.Get("username")
		userpass := result.Get("userpass")
		return AccessRequest{Username: username.String(), Userpass: userpass.String(), Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}