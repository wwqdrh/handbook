package cookbook

import "fmt"

func asteroidCollision(asteroids []int) []int {
	st := []int{}
	for _, aster := range asteroids {
		alive := true
		for alive && aster < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			alive = st[len(st)-1] < -aster // aster 是否存在
			if st[len(st)-1] <= -aster {   // 栈顶行星爆炸
				st = st[:len(st)-1]
			}
		}
		if alive {
			st = append(st, aster)
		}
	}
	return st
}

func loopslice() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println(r)
}

func loopMap() {
	var a = map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	flag := 0 // 保证将元素加到的桶中是没有被访问过的，这样才能保证被遍历到
	for i, v := range a {
		if flag == 0 {
			a[4] = 4
			a[5] = 5
			flag = 1
		}

		fmt.Println(i, v)
		delete(a, i)
	}
	fmt.Printf("%v\n", a)
}

func switchCase() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			// fallthrough
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}

func loopVariable() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k)

	for k = range (*[3]int)(nil) {
		fmt.Println(".")
	}
	fmt.Println(k)
}
