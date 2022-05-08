package main

// import (
// 	"testing"
// 	"wwqdrh/handbook/net/gcasbin"
// )

// func TestRbacBasicEnforcer(t *testing.T) {
// 	e, err := gcasbin.BasicEnforcer("rbac-model.conf", "rbac-policy.csv")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	requests := [][3]string{
// 		{"alice", "data1", "read"},
// 		{"alice", "data1", "write"},
// 		{"alice", "data2", "write"},
// 	}
// 	except := []bool{true, false, true}

// 	for i, item := range requests {
// 		if e.Enforce(item[0], item[1], item[2]) != except[i] {
// 			t.Error("认证失败", item, "期望为", except[i])
// 		}
// 	}
// }
