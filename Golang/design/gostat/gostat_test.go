package gostat

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// 测试服务注册后的数量
func TestServiceNum(t *testing.T) {
	DefaultSrvManager.Register("simple", func(ic IserviceCtx) {
		for i := 0; i < 100; i++ {
			ic.Sleep(100 * time.Millisecond)
			fmt.Println(i)
		}
		fmt.Println("simple done")
	})

	require.Equal(t, 1, DefaultSrvManager.Count())
	err := DefaultSrvManager.Start("simple")
	require.Nil(t, err)
	time.Sleep(300 * time.Millisecond)
	err = DefaultSrvManager.Stop("simple")
	require.Nil(t, err)

	fmt.Println(DefaultSrvManager.StatAll())

	// 再次启动
	err = DefaultSrvManager.Start("simple")
	require.Nil(t, err)
	time.Sleep(100 * time.Millisecond) // 等待状态变更完成
	fmt.Println(DefaultSrvManager.StatAll())
	time.Sleep(500 * time.Millisecond)
	err = DefaultSrvManager.Stop("simple")
	require.Nil(t, err)
}
