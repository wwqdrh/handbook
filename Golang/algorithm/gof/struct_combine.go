package gof

import "fmt"

// 组合模式

// UIComponent UI组件接口，对于任何UI控件都适用。
type UIComponent interface {
	// PrintUIComponent 打印UI组件
	PrintUIComponent()
	// GetUIControlName 获取控件名字
	GetUIControlName() string
	// GetConcreteUIControlName 获取控件具体名字
	GetConcreteUIControlName() string
}

// UIComponentAddtion UI组件附加接口，使用接口隔离原则，保证不需要实现接口声明方法的结构体，没有额外负担。仅对容器类型对UI控件适用。
type UIComponentAddtion interface {
	// AddUIComponent 添加UI组件
	AddUIComponent(component UIComponent)
	// AddUIComponents 添加UI组件列表
	AddUIComponents(components []UIComponent)
	// GetUIComponentList 获取UI组件列表
	GetUIComponentList() []UIComponent
}

// UIAttr UI属性
type UIAttr struct {
	// UI 名字
	Name string
}

////////////////////
// 各个控件
////////////////////
// Picture 图片，实现UIComponent接口
type Picture struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (picture Picture) GetUIControlName() string {
	return "Picture"
}

// GetConcreteUIControlName 获取控件具体名字
func (picture Picture) GetConcreteUIControlName() string {
	return picture.Name
}

// PrintUIComponent 打印UI组件
func (picture Picture) PrintUIComponent() {
	client.printCurrentControl(picture)
}

// Button 按钮，实现UIComponent接口
type Button struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (button Button) GetUIControlName() string {
	return "Button"
}

// GetConcreteUIControlName 获取控件具体名字
func (button Button) GetConcreteUIControlName() string {
	return button.Name
}

// PrintUIComponent 打印UI组件
func (button Button) PrintUIComponent() {
	client.printCurrentControl(button)
}

// Label 标签，实现UIComponent接口
type Label struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (label Label) GetUIControlName() string {
	return "Label"
}

// GetConcreteUIControlName 获取控件具体名字
func (label Label) GetConcreteUIControlName() string {
	return label.Name
}

// PrintUIComponent 打印UI组件
func (label Label) PrintUIComponent() {
	client.printCurrentControl(label)
}

// TextBox 文本框，实现UIComponent接口
type TextBox struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (textBox TextBox) GetUIControlName() string {
	return "TextBox"
}

// GetConcreteUIControlName 获取控件具体名字
func (textBox TextBox) GetConcreteUIControlName() string {
	return textBox.Name
}

// PrintUIComponent 打印UI组件
func (textBox TextBox) PrintUIComponent() {
	client.printCurrentControl(textBox)
}

// PassWordBox 密码框，实现UIComponent接口
type PassWordBox struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (passWordBox PassWordBox) GetUIControlName() string {
	return "PassWordBox"
}

// GetConcreteUIControlName 获取控件具体名字
func (passWordBox PassWordBox) GetConcreteUIControlName() string {
	return passWordBox.Name
}

// PrintUIComponent 打印UI组件
func (passWordBox PassWordBox) PrintUIComponent() {
	client.printCurrentControl(passWordBox)
}

// CheckBox 复选框，实现UIComponent接口
type CheckBox struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (checkBox CheckBox) GetUIControlName() string {
	return "CheckBox"
}

// GetConcreteUIControlName 获取控件具体名字
func (checkBox CheckBox) GetConcreteUIControlName() string {
	return checkBox.Name
}

// PrintUIComponent 打印UI组件
func (checkBox CheckBox) PrintUIComponent() {
	client.printCurrentControl(checkBox)
}

// LinkLabel 关联的标签，实现UIComponent接口
type LinkLabel struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (linkLabel LinkLabel) GetUIControlName() string {
	return "LinkLabel"
}

// GetConcreteUIControlName 获取控件具体名字
func (linkLabel LinkLabel) GetConcreteUIControlName() string {
	return linkLabel.Name
}

// PrintUIComponent 打印UI组件
func (linkLabel LinkLabel) PrintUIComponent() {
	client.printCurrentControl(linkLabel)
}

// 具体的对象
//client 打印client
var client = &PrintClient{}

// WinForm 窗口，实现UIComponent、UIComponentAddtion接口
type WinForm struct {
	// UIAttr 嵌套UI属性
	UIAttr
	// Components 容器的组件列表
	Components []UIComponent
}

// GetUIControlName 获取控件名字
func (window *WinForm) GetUIControlName() string {
	return "WinForm"
}

// GetConcreteUIControlName 获取控件具体名字
func (window *WinForm) GetConcreteUIControlName() string {
	return window.Name
}

// PrintUIComponent 打印UI组件
func (window *WinForm) PrintUIComponent() {
	client.printContainer(window, window)
}

// AddUIComponent 添加UI组件
func (window *WinForm) AddUIComponent(component UIComponent) {
	window.Components = append(window.Components, component)
}

// AddUIComponents 添加UI组件列表
func (window *WinForm) AddUIComponents(components []UIComponent) {
	window.Components = append(window.Components, components...)
}

// GetUIComponentList 获取UI组件列表
func (window *WinForm) GetUIComponentList() []UIComponent {
	return window.Components
}

type Frame struct {
	// UIAttr 嵌套UI属性
	UIAttr
	// Components 容器的组件列表
	Components []UIComponent
}

// GetUIControlName 获取控件名字
func (frame *Frame) GetUIControlName() string {
	return "Frame"
}

// GetConcreteUIControlName 获取控件具体名字
func (frame *Frame) GetConcreteUIControlName() string {
	return frame.Name
}

// PrintUIComponent 打印UI组件
func (frame *Frame) PrintUIComponent() {
	client.printContainer(frame, frame)
}

// AddUIComponent 添加UI组件
func (frame *Frame) AddUIComponent(component UIComponent) {
	frame.Components = append(frame.Components, component)
}

// AddUIComponents 添加UI组件列表
func (frame *Frame) AddUIComponents(components []UIComponent) {
	frame.Components = append(frame.Components, components...)
}

// GetUIComponentList 获取UI组件列表
func (frame *Frame) GetUIComponentList() []UIComponent {
	return frame.Components
}

////////////////////
// client
////////////////////

// PrintClient 打印客户端
type PrintClient struct{}

// printContainer 打印容器控件
func (client PrintClient) printContainer(component UIComponent, componentAddtion UIComponentAddtion) {
	client.printCurrentControl(component)
	for _, v := range componentAddtion.GetUIComponentList() {
		v.PrintUIComponent()
	}
}

// printCurrentControl 打印当前控件
func (client PrintClient) printCurrentControl(component UIComponent) {
	fmt.Println(fmt.Sprintf("print %s(%s)", component.GetUIControlName(), component.GetConcreteUIControlName()))
}
