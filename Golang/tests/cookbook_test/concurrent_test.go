package cookbooktest

import (
	"testing"
	"wwqdrh/handbook/cookbook/concurrent"
)

func TestMultiPrint(t *testing.T) {
	if res := concurrent.MultiPrint(); res != "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728" {
		t.Errorf("交替打印失败，res结果: %s", res)
	} else {
		t.Log("交替打印成功")
	}
}
func TestMockCommunicate(t *testing.T) {
	if res := concurrent.MockCommunicate(); res != "11111" {
		t.Error("两个协程通信失败")
	}
}
