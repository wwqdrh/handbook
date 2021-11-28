package algorithm_test

import (
	"fmt"
	"sync"
	"testing"
	"wwqdrh/handbook/algorithm/gof"
)

func TestFactoryMode(t *testing.T) {
	var (
		factory gof.OperatorFactory
	)

	factory = gof.PlusOperatorFactory{}
	if gof.FactoryMode(factory, 1, 2) != 3 {
		t.Error("error with factory method pattern")
	}

	factory = gof.MinusOperatorFactory{}
	if gof.FactoryMode(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}

func ExampleAbstractFactoryMode() {
	var factory gof.FruitFactory
	var apple, banana, orange gof.Fruit

	//创建苹果工厂，生产苹果，吃苹果
	factory = &gof.AppleFactory{}
	apple = factory.CreateFruit()
	apple.Eat()

	//创建香蕉工厂，生产香蕉，吃香蕉
	factory = &gof.BananaFactory{}
	banana = factory.CreateFruit()
	banana.Eat()

	//创建橘子工厂，生产橘子，吃橘子
	factory = &gof.OrangeFactory{}
	orange = factory.CreateFruit()
	orange.Eat()

	// output: 吃苹果
	// 吃香蕉
	// 吃橘子
}

func ExampleBuilderMode() {
	pb := gof.NewPersonBuilder()
	pb.Lives().
		At("Bangalore").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)

	person := pb.Build()

	fmt.Println(person)

	// output: &{ Bangalore 560102 Bangalore IBM Software Engineer 150000}
}

func TestPrototypeMode(t *testing.T) {
	p := gof.NewConcretePrototype("hello world")
	newProto := p.Clone()
	if newProto.Name() != p.Name() {
		t.Error("原型对象复制失败")
	}
}

func TestSingletonMode(t *testing.T) {
	var (
		a, b, c *gof.SingletonType
	)

	wg := sync.WaitGroup{}
	wg.Add(3)

	func() {
		a = gof.GetInstance()
		wg.Done()
	}()

	func() {
		b = gof.GetInstance()
		wg.Done()
	}()

	func() {
		c = gof.GetInstance()
		wg.Done()
	}()

	wg.Wait()

	if a != b || b != c {
		t.Error("singleton失败")
	}

}

func ExampleAdapterMode() {
	adaptee := &gof.Volts220{}
	target := &gof.Adapter{Adaptee: adaptee}
	target.CovertTo5V()

	// output: 电源输出了220V电压
	// 通过手机电源适配器，转成了5V电压，可供手机充电
}

func ExampleCombine() {
	window := &gof.WinForm{UIAttr: gof.UIAttr{Name: "WINDOW窗口"}}
	picture := &gof.Picture{gof.UIAttr{Name: "LOGO图片"}}
	loginButton := &gof.Button{gof.UIAttr{Name: "登录"}}
	registerButton := &gof.Button{gof.UIAttr{Name: "注册"}}
	frame := &gof.Frame{UIAttr: gof.UIAttr{Name: "FRAME1"}}
	userLable := &gof.Label{gof.UIAttr{Name: "用户名"}}
	textBox := &gof.TextBox{gof.UIAttr{Name: "文本框"}}
	passwordLable := &gof.Label{gof.UIAttr{Name: "密码"}}
	passwordBox := &gof.PassWordBox{gof.UIAttr{Name: "密码框"}}
	checkBox := &gof.CheckBox{gof.UIAttr{Name: "复选框"}}
	rememberUserTextBox := &gof.TextBox{gof.UIAttr{Name: "记住用户名"}}
	linkLable := &gof.LinkLabel{gof.UIAttr{Name: "忘记密码"}}

	window.AddUIComponents([]gof.UIComponent{picture, loginButton, registerButton, frame})
	frame.AddUIComponents([]gof.UIComponent{userLable, textBox, passwordLable, passwordBox, checkBox, rememberUserTextBox, linkLable})

	window.PrintUIComponent()

	// output: print WinForm(WINDOW窗口)
	// print Picture(LOGO图片)
	// print Button(登录)
	// print Button(注册)
	// print Frame(FRAME1)
	// print Label(用户名)
	// print TextBox(文本框)
	// print Label(密码)
	// print PassWordBox(密码框)
	// print CheckBox(复选框)
	// print TextBox(记住用户名)
	// print LinkLabel(忘记密码)
}

func ExampleBridgeMode() {
	// 大杯咖啡+奶
	var (
		coffeeAddtion gof.ICoffeeAddtion
		coffee        gof.ICoffee
	)

	coffeeAddtion = gof.NewCoffeeAddtion(gof.CoffeeAddtionTypeMilk)
	coffee = gof.NewCoffee(gof.CoffeeCupTypeLarge, coffeeAddtion)
	coffee.OrderCoffee()

	// 大杯咖啡+糖
	coffeeAddtion = gof.NewCoffeeAddtion(gof.CoffeeAddtionTypeSugar)
	coffee = gof.NewCoffee(gof.CoffeeCupTypeLarge, coffeeAddtion)
	coffee.OrderCoffee()

	// output: 订购了大杯咖啡
	// 加奶
	// 订购了大杯咖啡
	// 加糖
}

func ExampleDecoratorMode() {
	var book gof.Booker
	book = &gof.Book{}
	book.Reading()
	fmt.Println("============")

	var notesTake gof.NotesTaker
	notesTake = &gof.ConcreteNotesTake{Booker: book}
	notesTake.Reading()
	notesTake.TakeNotes()
	fmt.Println("============")

	var Underline gof.Underliner
	Underline = &gof.ConcreteUnderline{Booker: book}
	Underline.Reading()
	Underline.Underline()

	// output: 我正在读书
	// ============
	// 我正在读书
	// 我正在记笔记
	// ============
	// 我正在读书
	// 我正在划线
}

func ExampleFacade() {
	facade := gof.RegisterFacade{}
	facade.Register("18618193858")

	// output: 创建了手机号码为18618193858的账户
	// 为手机号码为18618193858的用户创建虚拟钱包账户
	// 为手机号码为18618193858的用户发放优惠券
}

func ExampleFlyWeightMode() {
	board1 := gof.ChessBoard{}
	board1.Init()
	board1.Move(1, 10, 20)

	board2 := gof.ChessBoard{}
	board2.Init()
	board2.Move(2, 6, 6)

	// output: 黑色的將被移动到了(10，20)的位置
	// 红色的帥被移动到了(6，6)的位置
}

func ExampleProxyMode() {
	url := "http://baidu.com"
	download := &gof.ConcreteDownload{Url: url}
	proxy := &gof.DownloadProxy{
		Url:        url,
		Downloader: download,
	}
	proxy.Download()

	// output: 准备开始下载http://baidu.com
	// http://baidu.com 在下载中
	// 下载http://baidu.com完成
}

func ExampleResponsChain() {
	var leader, director, cfo gof.Manager
	leader = &gof.Leader{}
	director = &gof.Director{}
	cfo = &gof.CFO{}

	handlerChain := gof.HandlerChain{}
	handlerChain.AddHandler(leader)
	handlerChain.AddHandler(director)
	handlerChain.AddHandler(cfo)

	request := gof.Request{
		Name:   "kay",
		Amount: 200,
	}
	handlerChain.HandleRequest(request)

	// output: leader 处理了kay的200元报销
}

func ExampleMemorandum() {
	// 发起者初始状态
	originator := gof.Originator{}
	originator.State = "start"
	originator.Print()

	// 设置备忘录
	caretaker := gof.NewCaretaker(originator.CreateMemento())
	originator.State = "Stage Two"
	originator.Print()

	// 恢复为备忘录
	originator.RestoreMemento(caretaker)
	originator.Print()

	// output: originator的状态为start
	// originator的状态为Stage Two
	// originator的状态为start
}

func ExampleNotifyMode() {
	var subject gof.Subject
	subject = &gof.StudentSubject{}

	var mathTeacherObserver, englishTeacherObserver, motherObserver, fatherObserver gof.Observer
	mathTeacherObserver = &gof.TeacherObserver{gof.BaseObserver{"数学"}}
	englishTeacherObserver = &gof.TeacherObserver{gof.BaseObserver{"英语"}}
	motherObserver = &gof.ParentObserver{gof.BaseObserver{"妈妈"}}
	fatherObserver = &gof.ParentObserver{gof.BaseObserver{"爸爸"}}

	//只添加数学老师观察者，那么只有数学老师会收到作业完成通知
	fmt.Println("******只添加数学老师观察者，那么只有数学老师会收到作业完成通知******")
	subject.AddObserver(mathTeacherObserver)
	subject.NotifyObservers()
	fmt.Println()

	//又批量添加了英语老师、妈妈、爸爸，那么数学老师、英语老师、妈妈、爸爸，都会收到作业完成通知
	fmt.Println("******又批量添加了英语老师、妈妈、爸爸，那么数学老师、英语老师、妈妈、爸爸，都会收到作业完成通知******")
	subject.AddObservers(englishTeacherObserver, motherObserver, fatherObserver)
	subject.NotifyObservers()
	fmt.Println()

	//移除了妈妈观察者，那么只有数学老师、英语老师、爸爸，会收到作业完成通知
	fmt.Println("******移除了妈妈观察者，那么只有数学老师、英语老师、爸爸，会收到作业完成通知******")
	subject.RemoveObserver(motherObserver)
	subject.NotifyObservers()
	fmt.Println()

	//移除了所有观察者，那么不会有人收到作业完成通知
	fmt.Println("******移除了所有观察者，那么不会有人收到作业完成通知******")
	subject.RemoveAllObservers()
	subject.NotifyObservers()
	fmt.Println()

	// output: ******只添加数学老师观察者，那么只有数学老师会收到作业完成通知******
	// 学生写完了作业，通知观察者们
	// 数学老师收到了作业

	// ******又批量添加了英语老师、妈妈、爸爸，那么数学老师、英语老师、妈妈、爸爸，都会收到作业完成通知******
	// 学生写完了作业，通知观察者们
	// 数学老师收到了作业
	// 英语老师收到了作业
	// 妈妈(家长)收到了作业
	// 爸爸(家长)收到了作业

	// ******移除了妈妈观察者，那么只有数学老师、英语老师、爸爸，会收到作业完成通知******
	// 学生写完了作业，通知观察者们
	// 数学老师收到了作业
	// 英语老师收到了作业
	// 爸爸(家长)收到了作业

	// ******移除了所有观察者，那么不会有人收到作业完成通知******
	// 学生写完了作业，通知观察者们
}

func ExampleState() {
	user := gof.StateUser{}
	user.WatchVideo()

	user.PurchaseVip()
	user.WatchVideo()

	user.Expire()
	user.WatchVideo()

	// output: 看广告中...
	// 您是尊敬的vip用户，已为您跳过120s广告
	// 看广告中...
}

func ExampleTemplate() {
	var company *gof.Company

	company = gof.NewCompany(&gof.MyAskForLeaveRequest{})
	company.AskLeave()

	company = gof.NewCompany(&gof.TomAskForLeaveRequest{})
	company.AskLeave()

	// output: kaysun 因为 给娃打疫苗，请假 0.5 天
	// tom 因为 回家探亲，请假 5.0 天
}

func ExampleVisitor() {
	e := new(gof.Element)
	e.Accept(new(gof.ProductionVisitor))
	e.Accept(new(gof.TestingVisitor))
	m := new(gof.EnvExample)
	m.Print(new(gof.ProductionVisitor))
	m.Print(new(gof.TestingVisitor))

	// output: 这是生产环境
	// 这是测试环境
	// 这是生产环境
	// 这是测试环境
}

func ExampleMediator() {
	concreteMediator := &gof.ConcreteMediator{}
	c1 := &gof.ConcreteColleague1{}
	c2 := &gof.ConcreteColleague2{}
	c1.SetMediator(concreteMediator)
	c2.SetMediator(concreteMediator)
	concreteMediator.C1 = c1
	concreteMediator.C2 = c2
	c1.Send("吃饭了吗c2")
	c2.Send("我吃过了c1")

	// output: ConcreteColleague2 recv msg: 吃饭了吗c2
	// ConcreteColleague1 recv msg: 我吃过了c1
}

func ExampleIterator() {
	array := []interface{}{6, 8, 7, 2, 5, 0, 3, 2}
	a := 0
	iterator := gof.NewArrayIterator(array, a)
	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(int)
		for value != array[index].(int) {
			fmt.Println("error")
		}
		fmt.Println(index, value)
	}

	// output: 0 6
	// 1 8
	// 2 7
	// 3 2
	// 4 5
	// 5 0
	// 6 3
	// 7 2
}
