package transaction

import (
	"context"
	"testing"
	"time"

	"github.com/dtm-labs/dtmcli/logger"
)

func TestFlushApp(t *testing.T) {
	logger.InitLog("debug")
	_, err := rdb.Set(context.Background(), stockKey, "4", 86400*time.Second).Result()
	logger.FatalIfError(err)
	startSvr()
	select {}
}
