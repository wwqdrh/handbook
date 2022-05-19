package http_server

import (
	// _ "git.zx-tech.net/ljhua/huoban_erp/controller/docs"

	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// 提供引入swag的样例代码

func SwaggerInstall(handler *http.Handler) {
	handler.Handle("/swagger/", httpSwagger.Handler())
}

// 示例
// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]

// produce、accept都是支持mimetype类型
// json	application/json
// xml	text/xml
// plain	text/plain
// html	text/html
// mpfd	multipart/form-data
// x-www-form-urlencoded	application/x-www-form-urlencoded
// json-api	application/vnd.api+json
// json-stream	application/x-json-stream
// octet-stream	application/octet-stream
// png	image/png
// jpeg	image/jpeg
// gif	image/gif

//
// @description This is the first line
// @description This is the second line
// @description And so forth.

// 参数
/// ...
// @Param group_id   path int true "Group ID"
// @Param account_id path int true "Account ID"
// ...
// @Router /examples/groups/{group_id}/accounts/{account_id} [get]
/// ...
// @Param group_id path int true "Group ID"
// @Param user_id  path int true "User ID"
// ...
// @Router /examples/groups/{group_id}/user/{user_id}/address [put]
// @Router /examples/user/{user_id}/address [put]
// @Param email body string true "message/rfc822" SchemaExample(Subject: Testmail\r\n\r\nBody Message\r\n)

// 响应 code {类型} 具体值
// @Success 200 {array} model.Account <-- This is a user defined struct.
// JSONResult's data field will be overridden by the specific type proto.Order
// @success 200 {object} jsonresult.JSONResult{data=proto.Order} "desc"
// @Success      200              {string}  string    "ok"
// @failure      400              {string}  string    "error"
// @response     default          {string}  string    "other error"
// @Header       200              {string}  Location  "/entity/1"
// @Header       200,400,default  {string}  Token     "token"
// @Header       all              {string}  Token2    "token2"

// param
// query
// path
// header
// body
// formData

// 类型也可以描述
// Account model info
// @Description User account information
// @Description with user id and username
// type Account struct {
// 	// ID this is userid
// 	ID   int    `json:"id"`
// 	Name string `json:"name"` // This is Name
// }

// secutiry
// @securityDefinitions.basic BasicAuth

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @Security ApiKeyAuth
// @Security OAuth2Application[write, admin]
