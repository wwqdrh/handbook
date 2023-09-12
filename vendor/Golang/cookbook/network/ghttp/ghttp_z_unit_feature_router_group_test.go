// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/test/gtest"
)

// 执行对象
type GroupObject struct{}

func (o *GroupObject) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *GroupObject) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *GroupObject) Index(r *ghttp.Request) {
	r.Response.Write("Object Index")
}

func (o *GroupObject) Show(r *ghttp.Request) {
	r.Response.Write("Object Show")
}

func (o *GroupObject) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete")
}

func Handler(r *ghttp.Request) {
	r.Response.Write("Handler")
}

func Test_Router_GroupBasic1(t *testing.T) {
	p, _ := gtcp.GetFreePort()
	s := g.Server(p)
	obj := new(GroupObject)
	// 分组路由方法注册
	group := s.Group("/api")
	group.ALL("/handler", Handler)
	group.ALL("/obj", obj)
	group.GET("/obj/my-show", obj, "Show")
	group.REST("/obj/rest", obj)
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent(ctx, "/api/handler"), "Handler")

		t.Assert(client.GetContent(ctx, "/api/obj"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/api/obj/"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/api/obj/index"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/api/obj/delete"), "1Object Delete2")
		t.Assert(client.GetContent(ctx, "/api/obj/my-show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/api/obj/show"), "1Object Show2")
		t.Assert(client.DeleteContent(ctx, "/api/obj/rest"), "1Object Delete2")

		t.Assert(client.DeleteContent(ctx, "/ThisDoesNotExist"), "Not Found")
		t.Assert(client.DeleteContent(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_GroupBuildInVar(t *testing.T) {
	p, _ := gtcp.GetFreePort()
	s := g.Server(p)
	obj := new(GroupObject)
	// 分组路由方法注册
	group := s.Group("/api")
	group.ALL("/{.struct}/{.method}", obj)
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent(ctx, "/api/group-object/index"), "1Object Index2")
		t.Assert(client.GetContent(ctx, "/api/group-object/delete"), "1Object Delete2")
		t.Assert(client.GetContent(ctx, "/api/group-object/show"), "1Object Show2")

		t.Assert(client.DeleteContent(ctx, "/ThisDoesNotExist"), "Not Found")
		t.Assert(client.DeleteContent(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Methods(t *testing.T) {
	p, _ := gtcp.GetFreePort()
	s := g.Server(p)
	obj := new(GroupObject)
	group := s.Group("/")
	group.ALL("/obj", obj, "Show, Delete")
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent(ctx, "/obj/show"), "1Object Show2")
		t.Assert(client.GetContent(ctx, "/obj/delete"), "1Object Delete2")
	})
}

func Test_Router_Group_MultiServer(t *testing.T) {
	p1, _ := gtcp.GetFreePort()
	p2, _ := gtcp.GetFreePort()
	s1 := g.Server(p1)
	s2 := g.Server(p2)
	s1.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/post", func(r *ghttp.Request) {
			r.Response.Write("post1")
		})
	})
	s2.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/post", func(r *ghttp.Request) {
			r.Response.Write("post2")
		})
	})
	s1.SetPort(p1)
	s2.SetPort(p2)
	s1.SetDumpRouterMap(false)
	s2.SetDumpRouterMap(false)
	gtest.Assert(s1.Start(), nil)
	gtest.Assert(s2.Start(), nil)
	defer s1.Shutdown()
	defer s2.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c1 := g.Client()
		c1.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p1))
		c2 := g.Client()
		c2.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p2))
		t.Assert(c1.PostContent(ctx, "/post"), "post1")
		t.Assert(c2.PostContent(ctx, "/post"), "post2")
	})
}

func Test_Router_Group_Map(t *testing.T) {
	testFuncGet := func(r *ghttp.Request) {
		r.Response.Write("get")
	}
	testFuncPost := func(r *ghttp.Request) {
		r.Response.Write("post")
	}
	p, _ := gtcp.GetFreePort()
	s := g.Server(p)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Map(map[string]interface{}{
			"Get: /test": testFuncGet,
			"Post:/test": testFuncPost,
		})
	})
	s.SetPort(p)
	//s.SetDumpRouterMap(false)
	gtest.Assert(s.Start(), nil)
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(c.GetContent(ctx, "/test"), "get")
		t.Assert(c.PostContent(ctx, "/test"), "post")
	})
}

// https://github.com/gogf/gf/issues/1609
func Test_Issue1609(t *testing.T) {
	p, _ := gtcp.GetFreePort()
	s := g.Server(p)
	group := s.Group("/api/get")
	group.GET("/", func(r *ghttp.Request) {
		r.Response.Write("get")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	gtest.Assert(s.Start(), nil)
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(c.GetContent(ctx, "/api/get"), "get")
		t.Assert(c.PostContent(ctx, "/test"), "Not Found")
	})
}
