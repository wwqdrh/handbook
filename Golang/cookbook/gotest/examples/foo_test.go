package main

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Live(t *testing.T) {
	ctrl := gomock.NewController(t)
	life := NewMockLife(ctrl)
	handler := func(money int64) error {
		if money <= 0 {
			return errors.New("error")
		}
		return nil
	}
	life.EXPECT().GoodGoodStudy(gomock.Any()).AnyTimes().DoAndReturn(handler)
	life.EXPECT().BuyHouse(gomock.Any()).AnyTimes().DoAndReturn(handler)
	life.EXPECT().Marry(gomock.Any()).AnyTimes().DoAndReturn(handler)
	Convey("Live", t, func() {
		person := &Person{
			life: life,
		}
		Convey("GoodGoodStudy  error", func() {
			So(person.Live(0, 100, 100), ShouldBeError)
		})
		Convey("GoodGoodStudy  ok", func() {
			Convey("BuyHouse  error", func() {
				So(person.Live(100, 0, 100), ShouldBeError)
			})
			Convey("BuyHouse  ok", func() {
				Convey("Marry  error", func() {
					So(person.Live(100, 100, 0), ShouldBeError)
				})
				Convey("Marry  ok", func() {
					So(person.Live(100, 100, 100), ShouldBeNil)
				})
			})
		})
	})
}
