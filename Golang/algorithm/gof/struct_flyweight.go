package gof

import (
	"fmt"
	"sync"
)

type Color string

const (
	ColorBlack Color = "黑色"
	ColorRed   Color = "红色"
)

var (
	// once 保证只执行一次
	// once sync.Once

	// instance 单例对象
	PieceUnitFactory *pieceUnitFactory
)

// PieceUnit 棋子享元
type PieceUnit struct {
	// ID 唯一标识
	ID int
	// Text 棋子的文字
	Text string
	// Color 棋子的颜色
	Color Color
}

// pieceUnitFactory 棋子享元对象工厂
type pieceUnitFactory struct {
	//sync.RWMutex
	// PieceUnitMap 棋子id与享元对象的对应关系的map
	PieceUnitMap map[int]PieceUnit
}

// GetPieceUnitFactory 获取单例对象。 注：必须使用指针对象，否则会在返回时，拷贝一份，对象就不是同一个了。
func GetPieceUnitFactory() *pieceUnitFactory {
	once.Do(func() {
		PieceUnitFactory = &pieceUnitFactory{}
		PieceUnitFactory.PieceUnitMap = make(map[int]PieceUnit)
		PieceUnitFactory.init()
	})
	return PieceUnitFactory
}

// init 初始化工厂
func (factory pieceUnitFactory) init() {
	// 不需要锁，因为通过once.Do创建
	//factory.Lock()
	//defer factory.Unlock()

	factory.PieceUnitMap[1] = PieceUnit{1, "將", ColorBlack}
	factory.PieceUnitMap[2] = PieceUnit{2, "帥", ColorRed}
	factory.PieceUnitMap[3] = PieceUnit{3, "馬", ColorBlack}
	factory.PieceUnitMap[4] = PieceUnit{4, "傌", ColorRed}
	factory.PieceUnitMap[5] = PieceUnit{5, "馬", ColorBlack}
	factory.PieceUnitMap[6] = PieceUnit{6, "傌", ColorRed}
	factory.PieceUnitMap[7] = PieceUnit{7, "砲", ColorBlack}
	factory.PieceUnitMap[8] = PieceUnit{8, "炮", ColorRed}
	factory.PieceUnitMap[9] = PieceUnit{9, "砲", ColorBlack}
	factory.PieceUnitMap[10] = PieceUnit{10, "炮", ColorRed}
	factory.PieceUnitMap[11] = PieceUnit{11, "車", ColorBlack}
	factory.PieceUnitMap[12] = PieceUnit{12, "俥", ColorRed}
	factory.PieceUnitMap[13] = PieceUnit{13, "車", ColorBlack}
	factory.PieceUnitMap[14] = PieceUnit{14, "俥", ColorRed}
	factory.PieceUnitMap[15] = PieceUnit{15, "士", ColorBlack}
	factory.PieceUnitMap[16] = PieceUnit{16, "仕", ColorRed}
	factory.PieceUnitMap[17] = PieceUnit{17, "士", ColorBlack}
	factory.PieceUnitMap[18] = PieceUnit{18, "仕", ColorRed}
	factory.PieceUnitMap[19] = PieceUnit{19, "象", ColorBlack}
	factory.PieceUnitMap[20] = PieceUnit{20, "相", ColorRed}
	factory.PieceUnitMap[21] = PieceUnit{21, "象", ColorBlack}
	factory.PieceUnitMap[22] = PieceUnit{22, "相", ColorRed}
	factory.PieceUnitMap[23] = PieceUnit{23, "卒", ColorBlack}
	factory.PieceUnitMap[24] = PieceUnit{24, "兵", ColorRed}
	factory.PieceUnitMap[25] = PieceUnit{25, "卒", ColorBlack}
	factory.PieceUnitMap[26] = PieceUnit{26, "兵", ColorRed}
	factory.PieceUnitMap[27] = PieceUnit{27, "卒", ColorBlack}
	factory.PieceUnitMap[28] = PieceUnit{28, "兵", ColorRed}
	factory.PieceUnitMap[29] = PieceUnit{29, "卒", ColorBlack}
	factory.PieceUnitMap[30] = PieceUnit{30, "兵", ColorRed}
	factory.PieceUnitMap[31] = PieceUnit{31, "卒", ColorBlack}
	factory.PieceUnitMap[32] = PieceUnit{32, "兵", ColorRed}

}

// GetChessPieceUnit 根据id获取棋子享元对象
func (factory pieceUnitFactory) GetChessPieceUnit(id int) PieceUnit {
	return factory.PieceUnitMap[id]
}

// Piece 棋子
type Piece struct {
	// PieceUnit
	PieceUnit
	// PositionX
	PositionX int
	// PositionY
	PositionY int
}

// ChessBoard 棋盘
type ChessBoard struct {
	// sync.RWMutex 嵌套锁
	sync.RWMutex
	// ChessPieceMap 棋子id与棋子对象的对应关系的map
	ChessPieceMap map[int]*Piece
}

//init 初始化棋局。 You must call this。
func (board *ChessBoard) Init() {
	// 加锁
	board.Lock()
	defer board.Unlock()
	// 获取棋子享元工厂
	unitFactory := GetPieceUnitFactory()
	board.ChessPieceMap = make(map[int]*Piece)
	board.ChessPieceMap[1] = &Piece{
		PieceUnit: unitFactory.GetChessPieceUnit(1),
		PositionX: 0,
		PositionY: 1,
	}
	board.ChessPieceMap[2] = &Piece{
		PieceUnit: unitFactory.GetChessPieceUnit(2),
		PositionX: 5,
		PositionY: 1,
	}
	board.ChessPieceMap[3] = &Piece{
		PieceUnit: unitFactory.GetChessPieceUnit(3),
		PositionX: 2,
		PositionY: 8,
	}
	//......
}

// Move 棋盘上移动棋子
func (board *ChessBoard) Move(id int, positionX, positionY int) {
	board.ChessPieceMap[id].PositionX = positionX
	board.ChessPieceMap[id].PositionY = positionY
	fmt.Println(fmt.Sprintf("%s的%s被移动到了(%d，%d)的位置", board.ChessPieceMap[id].Color, board.ChessPieceMap[id].Text, positionX, positionY))
}
