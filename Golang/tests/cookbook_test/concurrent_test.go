package cookbooktest

import (
	"testing"
	"wwqdrh/handbook/cookbook/concurrent"
)

func ExampleMultiPrint() {
	concurrent.MultiPrint()

	// Output: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
}
func TestMockCommunicate(t *testing.T) {
	if res := concurrent.MockCommunicate(); res != "11111" {
		t.Error("两个协程通信失败")
	}
}
