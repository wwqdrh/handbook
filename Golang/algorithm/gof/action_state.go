package gof

import "fmt"

type State uint8

const (
	StateNormal State = iota
	StateVip
)

// IUser 用户接口
type IStateUser interface {
	// WatchVideo 看视频
	WatchVideo()
}

// ISwitchState 转换状态接口
type ISwitchState interface {
	// PurchaseVip 购买会员
	PurchaseVip()
	// Expire 会员过期
	Expire()
}

// NormalUser 普通用户，实现IUser接口
type NormalUser struct{}

// VipUser 会员用户，实现IUser接口
type VipUser struct{}

// WatchVideo 看视频
func (user NormalUser) WatchVideo() {
	fmt.Println("看广告中...")
}

// WatchVideo 看视频
func (user VipUser) WatchVideo() {
	fmt.Println("您是尊敬的vip用户，已为您跳过120s广告")
}

// User 用户，实现ISwitchState、IUser接口
type StateUser struct {
	// UserState 用户
	UserState IStateUser
}

// SetUser 设置用户状态
func (user *StateUser) SetUser(userState IStateUser) {
	user.UserState = userState
}

// Expire 会员过期
func (user *StateUser) Expire() {
	user.UserState = NormalUser{}
}

// PurchaseVip 购买vip会员
func (user *StateUser) PurchaseVip() {
	user.UserState = VipUser{}
}

// WatchVideo 看视频
func (user *StateUser) WatchVideo() {
	if user.UserState == nil {
		// 默认为普通用户
		user.SetUser(NormalUser{})
	}
	user.UserState.WatchVideo()
}
