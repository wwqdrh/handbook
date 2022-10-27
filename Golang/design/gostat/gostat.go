package gostat

import (
	"errors"
	"time"

	"github.com/wwqdrh/logger"
)

var ErrCustomStop = errors.New("custom stop")

type (
	TaskHandler func(IserviceCtx)

	IserviceCtx interface {
		Status() string
		Handle(IserviceCtx)
		Stop() error
		Sleep(time.Duration) // 停止任务的核心
	}
)

var DefaultSrvManager = &serviceManager{
	srvs: make(map[string]IserviceCtx, 0),
}

type serviceManager struct {
	srvs map[string]IserviceCtx
}

type serviceCtx struct {
	handle TaskHandler
	stop   bool
	stoped chan struct{}
}

func (s *serviceCtx) Status() string {
	if s.stop {
		return "停止"
	} else {
		return "启动"
	}
}

func (s *serviceCtx) Handle(IserviceCtx) {
	defer func() {
		err := recover()
		if err != nil {
			if val, ok := err.(error); ok && errors.Is(val, ErrCustomStop) {
				logger.DefaultLogger.Info("task exit")
			}
		}
		s.stoped <- struct{}{}
		s.stop = true
	}()
	s.stop = false
	s.handle(s)
}

func (s *serviceCtx) Sleep(slp time.Duration) {
	tmout := time.After(slp)
	for {
		select {
		case <-tmout:
			return
		default:
			if s.stop {
				panic(ErrCustomStop)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *serviceCtx) Stop() error {
	s.stop = true
	select {
	case <-s.stoped:
		return nil
	case <-time.After(10 * time.Second):
		return errors.New("退出失败 请检查任务中是否调用了sleep主动将控制权让出")
	}
}

// 注册新的服务
func (s *serviceManager) Register(srvName string, handle TaskHandler) {
	s.srvs[srvName] = &serviceCtx{
		handle: handle,
		stop:   true,
		stoped: make(chan struct{}, 1),
	}
}

// 获取所有的stat服务
func (s *serviceManager) StatAll() map[string]string {
	res := map[string]string{}
	for name, ctx := range s.srvs {
		res[name] = ctx.Status()
	}
	return res
}

func (s *serviceManager) Count() int {
	return len(s.srvs)
}

func (s *serviceManager) Start(srvName string) error {
	val, ok := s.srvs[srvName]
	if !ok {
		return errors.New(srvName + "不存在")
	}
	go val.Handle(val)
	return nil
}

func (s *serviceManager) Stop(srvName string) error {
	val, ok := s.srvs[srvName]
	if !ok {
		return errors.New(srvName + "不存在")
	}
	return val.Stop()
}
