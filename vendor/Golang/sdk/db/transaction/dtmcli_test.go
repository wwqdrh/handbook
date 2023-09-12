package transaction

import "testing"

func TestSimpleDtmcli(t *testing.T) {
	QsStartSvr()
	_ = QsFireRequest()
	select {}
}
