package main

import (
	// "fmt"
	// "mime/multipart"
	"net/http"

	// "log"

	// "github.com/getkin/kin-openapi/openapi3"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/long2ice/swagin"
	// "github.com/long2ice/swagin/router"
	// "github.com/long2ice/swagin/security"
	// "github.com/long2ice/swagin/swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/wwqdrh/handbook/goswag/docs"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Param req body Hreq true "-"
// @Success 200 {object} Hres
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

type H struct {
	HelloReq Hreq

	HelloRes Hres
}

type Hreq struct {
	Name string `json:"name"`
}

type Hres struct {
	Code string `json:"code"`
}

func oldversion() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8081")

}

// type TestQueryReq struct {
// 	Name     string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
// 	Token    string `header:"token" validate:"required" json:"token" default:"test"`
// 	Optional string `query:"optional" json:"optional"`
// }

// func TestQuery(c *gin.Context, req TestQueryReq) {
// 	user := c.MustGet(security.Credentials).(*security.User)
// 	fmt.Println(user)
// 	c.JSON(http.StatusOK, req)
// }

// type TestQueryListReq struct {
// 	Name  string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
// 	Token string `header:"token" validate:"required" json:"token" default:"test"`
// }

// func TestQueryList(c *gin.Context, req TestQueryListReq) {
// 	user := c.MustGet(security.Credentials).(*security.User)
// 	fmt.Println(user)
// 	c.JSON(http.StatusOK, []TestQueryListReq{req})
// }

// type TestQueryPathReq struct {
// 	Name  string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
// 	ID    int    `uri:"id" validate:"required" json:"id" description:"id of model" default:"1"`
// 	Token string `header:"token" validate:"required" json:"token" default:"test"`
// }

// func TestQueryPath(c *gin.Context, req TestQueryPathReq) {
// 	c.JSON(http.StatusOK, req)
// }

// type TestFormReq struct {
// 	ID   int    `query:"id" validate:"required" json:"id" description:"id of model" default:"1"`
// 	Name string `form:"name" validate:"required" json:"name" description:"name of model" default:"test"`
// 	List []int  `form:"list" validate:"required" json:"list" description:"list of model"`
// }

// func TestForm(c *gin.Context, req TestFormReq) {
// 	c.JSON(http.StatusOK, req)
// }

// type TestNoModelReq struct {
// 	Authorization string `header:"authorization" validate:"required" json:"authorization" default:"authorization"`
// 	Token         string `header:"token" binding:"required" json:"token" default:"token"`
// }

// func TestNoModel(c *gin.Context, req TestNoModelReq) {
// 	c.JSON(http.StatusOK, req)
// }

// type TestFileReq struct {
// 	File *multipart.FileHeader `form:"file" validate:"required" description:"file upload"`
// }

// func TestFile(c *gin.Context, req TestFileReq) {
// 	c.JSON(http.StatusOK, gin.H{"file": req.File.Filename})
// }

// // TestQueryNoReq if there is no req body and query
// func TestQueryNoReq(c *gin.Context) {
// 	c.JSON(http.StatusOK, "{}")
// }

// func NewSwagger() *swagger.Swagger {
// 	return swagger.New("SwaGin", "Swagger + Gin = SwaGin", "0.1.0",
// 		swagger.License(&openapi3.License{
// 			Name: "Apache License 2.0",
// 			URL:  "https://github.com/long2ice/swagin/blob/dev/LICENSE",
// 		}),
// 		swagger.Contact(&openapi3.Contact{
// 			Name:  "long2ice",
// 			URL:   "https://github.com/long2ice/swagin",
// 			Email: "long2ice@gmail.com",
// 		}),
// 		swagger.TermsOfService("https://github.com/long2ice"),
// 	)
// }

// var (
// 	query = router.New(
// 		TestQuery,
// 		router.Summary("Test query"),
// 		router.Description("Test query model"),
// 		router.Security(&security.Basic{}),
// 		router.Responses(router.Response{
// 			"200": router.ResponseItem{
// 				Model:       TestQueryReq{},
// 				Description: "response model description",
// 			},
// 		}),
// 	)
// 	queryList = router.New(
// 		TestQueryList,
// 		router.Summary("Test query list"),
// 		router.Description("Test query list model"),
// 		router.Security(&security.Basic{}),
// 		router.Responses(router.Response{
// 			"200": router.ResponseItem{
// 				Model: []TestQueryListReq{},
// 			},
// 		}),
// 	)
// 	noModel = router.New(
// 		TestNoModel,
// 		router.Summary("Test no model"),
// 		router.Description("Test no model"),
// 	)
// 	queryPath = router.New(
// 		TestQueryPath,
// 		router.Summary("Test query path"),
// 		router.Description("Test query path model"),
// 	)
// 	formEncode = router.New(
// 		TestForm,
// 		router.Summary("Test form"),
// 		router.ContentType(binding.MIMEPOSTForm, router.ContentTypeRequest),
// 	)
// 	body = router.New(
// 		TestForm,
// 		router.Summary("Test json body"),
// 		router.Responses(router.Response{
// 			"200": router.ResponseItem{
// 				Model: TestFormReq{},
// 			},
// 		}),
// 	)
// 	file = router.New(
// 		TestFile,
// 		router.Summary("Test file upload"),
// 		router.ContentType(binding.MIMEMultipartPOSTForm, router.ContentTypeRequest),
// 	)
// )

// func newversion() {
// 	app := swagin.New(NewSwagger()).WithErrorHandler(func(ctx *gin.Context, err error, status int) {
// 		ctx.AbortWithStatusJSON(status, gin.H{
// 			"error": err.Error(),
// 		})
// 	})
// 	subApp := swagin.New(NewSwagger())
// 	subApp.GET("/noModel", noModel)
// 	app.Mount("/sub", subApp)
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"*"},
// 		AllowMethods:     []string{"*"},
// 		AllowHeaders:     []string{"*"},
// 		AllowCredentials: true,
// 	}))
// 	queryGroup := app.Group("/query", swagin.Tags("Query"))
// 	queryGroup.GET("/list", queryList)
// 	queryGroup.GET("/:id", queryPath)
// 	queryGroup.DELETE("", query)

// 	app.GET("/noModel", noModel)

// 	formGroup := app.Group("/form", swagin.Tags("Form"), swagin.Security(&security.Bearer{}))
// 	formGroup.POST("/encoded", formEncode)
// 	formGroup.PUT("", body)
// 	formGroup.POST("/file", file)

// 	if err := app.Run(); err != nil {
// 		log.Panic(err)
// 	}
// }
func main() {
	oldversion()
}
