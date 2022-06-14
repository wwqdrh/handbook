package transaction

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"encoding/json"
	"net/http"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lithammer/shortuuid"
)

// WrapHandler used by examples. much more simpler than WrapHandler2
func WrapHandler(fn func(*gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		began := time.Now()
		ret := fn(c)
		status, res := dtmcli.Result2HttpJSON(ret)

		b, _ := json.Marshal(res)
		if status == http.StatusOK || status == http.StatusTooEarly {
			logger.Infof("%2dms %d %s %s %s", time.Since(began).Milliseconds(), status, c.Request.Method, c.Request.RequestURI, string(b))
		} else {
			logger.Errorf("%2dms %d %s %s %s", time.Since(began).Milliseconds(), status, c.Request.Method, c.Request.RequestURI, string(b))
		}
		c.JSON(status, res)
	}
}

func MustBarrierFrom(c *gin.Context) *dtmcli.BranchBarrier {
	bb, err := dtmcli.BarrierFromQuery(c.Request.URL.Query())
	if err != nil {
		panic(err)
	}
	return bb
}

var redisOption = redis.Options{
	Addr:     "localhost:6379",
	Username: "root",
	Password: "123456",
}

var DtmServer = "http://localhost:36789/api/dtmsvr"

const BusiAPI = "/api/busi"
const BusiPort = 8081

var BusiUrl = fmt.Sprintf("http://localhost:%d%s", BusiPort, BusiAPI)

var rdb = redis.NewClient(&redisOption)

var stockKey = "{a}--stock-1"
var orderCreated int64

func startSvr() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	addRoutes(app)
	log.Printf("flash sales examples listening at %d", BusiPort)
	go app.Run(fmt.Sprintf(":%d", BusiPort))
	time.Sleep(100 * time.Millisecond)
}

func addRoutes(app *gin.Engine) {
	app.GET(BusiAPI+"/redisQueryPrepared", WrapHandler(func(c *gin.Context) interface{} {
		bb := MustBarrierFrom(c)
		return bb.RedisQueryPrepared(rdb, 7*86400)
	}))
	app.POST(BusiAPI+"/createOrder", WrapHandler(func(c *gin.Context) interface{} {
		logger.Infof("createOrder ------")
		atomic.AddInt64(&orderCreated, 1)
		return nil
	}))
	app.Any(BusiAPI+"/flashSales", WrapHandler(func(c *gin.Context) interface{} {
		gid := "{a}-" + shortuuid.New() // gid should contain same {a} as stockKey, so that the data will be in same redis slot
		msg := dtmcli.NewMsg(DtmServer, gid).
			Add(BusiUrl+"/createOrder", nil)
		return msg.DoAndSubmit(BusiUrl+"/redisQueryPrepared", func(bb *dtmcli.BranchBarrier) error {
			return bb.RedisCheckAdjustAmount(rdb, stockKey, -1, 86400)
		})
	}))
	app.Any(BusiAPI+"/flashSales-crash", WrapHandler(func(c *gin.Context) interface{} {
		gid := "{a}-" + shortuuid.New() // gid should contain same {a} as stockKey, so that the data will be in same redis slot
		msg := dtmcli.NewMsg(DtmServer, gid).
			Add(BusiUrl+"/createOrder", nil)
		msg.TimeoutToFail = 3
		return msg.DoAndSubmit(BusiUrl+"/redisQueryPrepared", func(bb *dtmcli.BranchBarrier) error {
			bb.RedisCheckAdjustAmount(rdb, stockKey, -1, 86400)
			select {} // mock crash
		})
	}))

	app.Any(BusiAPI+"/flashSales-batch", WrapHandler(func(c *gin.Context) interface{} {
		logger.InitLog("info")
		atomic.StoreInt64(&orderCreated, 0)
		_, err := rdb.Set(context.Background(), stockKey, "4", 86400*time.Second).Result()
		logger.FatalIfError(err)
		rest := dtmcli.GetRestyClient()
		go func() {
			rest.R().Post(BusiUrl + "/flashSales-crash")
		}()
		logger.Infof("sleeping 0.5s for a flash-sale request to go crash")
		time.Sleep(500 * time.Millisecond)
		for i := 0; i < 10; i++ {
			go func() {
				rest.R().Post(BusiUrl + "/flashSales")
			}()
		}
		logger.Infof("sleeping 0.5s for flash sale to finish normal requests")
		time.Sleep(500 * time.Millisecond)
		n := atomic.LoadInt64(&orderCreated)
		logger.Infof("normally created %d orders", n)
		logger.Infof("waiting for all orders created")
		for n < 4 {
			logger.Infof("total order created is: %d", n)
			time.Sleep(2 * time.Second)
			n = atomic.LoadInt64(&orderCreated)
		}
		logger.Infof("total order created is: %d", n)
		logger.InitLog("debug")
		return nil
	}))
}
