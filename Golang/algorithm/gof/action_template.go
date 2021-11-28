package gof

import "fmt"

// AskForLeaveRequest 请假单接口
type AskForLeaveRequest interface {
	GetName() string
	ReasonForLeave() string
	HowManyDaysForLeave() float32
}

// Company 公司
type Company struct {
	// AskForLeaveRequest 组合AskForLeaveRequest接口对象
	AskForLeaveRequest
}

func NewCompany(request AskForLeaveRequest) *Company {
	return &Company{request}
}

// AskLeave 请假
func (company Company) AskLeave() {
	fmt.Println(fmt.Sprintf("%s 因为 %s，请假 %.1f 天",
		company.GetName(), company.ReasonForLeave(), company.HowManyDaysForLeave()))
}

// 具体实现
// MyAskForLeaveRequest 我的请假单，实现AskForLeaveRequest接口
type MyAskForLeaveRequest struct {
}

// TomAskForLeaveRequest Tom的请假单，实现AskForLeaveRequest接口
type TomAskForLeaveRequest struct {
}

// GetName 请假人姓名
func (request MyAskForLeaveRequest) GetName() string {
	return "kaysun"
}

// ReasonForLeave 请假事由
func (request MyAskForLeaveRequest) ReasonForLeave() string {
	return "给娃打疫苗"
}

// HowManyDaysForLeave 请假多少天
func (request MyAskForLeaveRequest) HowManyDaysForLeave() float32 {
	return 0.5
}

// GetName 请假人姓名
func (request TomAskForLeaveRequest) GetName() string {
	return "tom"
}

// ReasonForLeave 请假事由
func (request TomAskForLeaveRequest) ReasonForLeave() string {
	return "回家探亲"
}

// HowManyDaysForLeave 请假多少天
func (request TomAskForLeaveRequest) HowManyDaysForLeave() float32 {
	return 5
}
