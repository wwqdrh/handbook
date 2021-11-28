package gof

import "fmt"

// IUser 用户接口
type IUser interface {
	// CreateUser 创建用户
	CreateUser()
}

// IVirtualAccount 虚拟账户接口
type IVirtualAccount interface {
	// CreateVirtualAcount 创建虚拟账户
	CreateVirtualAcount()
}

// ICoupon 优惠券接口
type ICoupon interface {
	// IssueCoupons 发放优惠券接口
	IssueCoupons()
}

// User 用户，实现用户接口
type User struct {
	// Telephone 电话号码
	Telephone string
}

// CreateUser 创建用户
func (user User) CreateUser() {
	fmt.Println(fmt.Sprintf("创建了手机号码为%s的账户", user.Telephone))
}

// VirtualAccount 虚拟账户，实现虚拟账户接口
type VirtualAccount struct {
	// User 虚拟账户嵌套用户
	User
}

func (virtualAccount VirtualAccount) CreateVirtualAcount() {
	fmt.Println(fmt.Sprintf("为手机号码为%s的用户创建虚拟钱包账户", virtualAccount.Telephone))
}

// Coupon 优惠券，实现优惠券接口
type Coupon struct {
	// User 优惠券嵌套用户
	User
}

func (coupon Coupon) IssueCoupons() {
	fmt.Println(fmt.Sprintf("为手机号码为%s的用户发放优惠券", coupon.Telephone))
}

// RegisterFacade 注册门面
type RegisterFacade struct {
	// IUser 组合用户接口
	IUser
	// IVirtualAccount 组合虚拟账户接口
	IVirtualAccount
	// ICoupon 组合优惠券接口
	ICoupon
}

func (facade RegisterFacade) Register(telephone string) {
	// 创建对象
	user := User{Telephone: telephone}
	virtualAccount := &VirtualAccount{user}
	coupon := &Coupon{user}
	// 创建用户
	user.CreateUser()
	// 创建虚拟账户
	virtualAccount.CreateVirtualAcount()
	// 对新账户发放优惠券
	coupon.IssueCoupons()
}
