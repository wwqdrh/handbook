package gof

import "fmt"

// 责任链 命令 迭代器 中介 备忘录 观察者 状态 策略 模板 访问者

// Manager 管理者接口
type Manager interface {
	// Review 审核
	Review(request Request) bool
}

// Request 报销请求
type Request struct {
	// Name 姓名
	Name string
	// Amount 金额
	Amount int
}

// HandlerChain 责任链处理器
type HandlerChain struct {
	// managers 责任链切片，不可导出
	managers []Manager
}

// AddHandler 责任链处理器对象，添加处理器。注：使用指针接收者，需要更改切片
func (chain *HandlerChain) AddHandler(manager Manager) {
	chain.managers = append(chain.managers, manager)
}

// AddHandlers 责任链处理器对象，批量添加处理器。注：使用指针接收者，需要更改切片
func (chain *HandlerChain) AddHandlers(managers ...Manager) {
	chain.managers = append(chain.managers, managers...)
}

// HandleRequest 责任链处理器对象，处理请求
func (chain *HandlerChain) HandleRequest(request Request) error {
	for _, manager := range chain.managers {
		result := manager.Review(request)
		// 如果result为true，则说明已经处理完成
		if result {
			return nil
		}
	}
	return fmt.Errorf("请求未被处理")
}

// Leader 直接上级只能审核500元以下的报销，实现Manager接口
type Leader struct{}

// Director 总监只能审核5000元以下的报销，实现Manager接口
type Director struct{}

// CFO 首席财务官可以审核任意金额的报销，实现Manager接口
type CFO struct{}

// Review Leader实现报销方法
func (leader Leader) Review(request Request) bool {
	if request.Amount < 500 {
		fmt.Println(fmt.Sprintf("leader 处理了%s的%d元报销", request.Name, request.Amount))
		return true
	}
	return false
}

// Review Director实现报销方法
func (director Director) Review(request Request) bool {
	if request.Amount < 5000 {
		fmt.Println(fmt.Sprintf("director 处理了%s的%d元报销", request.Name, request.Amount))
		return true
	}
	return false
}

// Review CFO实现报销方法
func (cfo CFO) Review(request Request) bool {
	fmt.Println(fmt.Sprintf("cfo 处理了%s的%d元报销", request.Name, request.Amount))
	return true
}
