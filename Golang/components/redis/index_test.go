package redis

import "testing"

func TestGetString(t *testing.T) {
	if StringGet("key") != "" {
		t.Error("发生错误")
	}
	if StringSet("key", "1") {
		if StringGet("key") != "1" {
			t.Error("发生错误")
		}
	} else {
		t.Error("发生错误")
	}
	if !StringDel("key") {
		t.Error("发生错误")
	}
}

func TestPipeline(t *testing.T) {
	PipelineIncr()
}

func TestTransaction(t *testing.T) {
	StringDel("index")
	if err := Transaction("index"); err != nil {
		t.Error("发生错误")
	} else {
		if StringGet("index") != "1" {
			t.Error("发生错误")
		}
	}
}

func TestPubSub(t *testing.T) {
	TryPub("testpub", "hello")
	if TrySub("testpub") != "hello" {
		t.Error("发生错误")
	}
}
