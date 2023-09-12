package network

import "testing"

func TestParseURL(t *testing.T) {
	ParseURL("postgres://user:pass@host.com:5432/path?k=v#f")
	ParseURL("postgres://user:pass@host.com:5432")
	ParseURL("etcd://root:123456@192.168.110.113:2379/dtmservice")
}
