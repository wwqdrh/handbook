package compiler

import (
	"fmt"
	"runtime"
)

func numFinizaler() {
	var i = 3
	// 这里的func 就是后面一直说的 finalizer
	runtime.SetFinalizer(&i, func(i *int) {
		fmt.Println(i, *i, "set finalizer")
	})
}

// // go run finalizer.go
// // go tool trace trace.out
// func main() {
// 	// f, _ := os.Create("trace.out")
// 	// defer f.Close()
// 	// trace.Start(f)
// 	// defer trace.Stop()

// 	numFinizaler()
// 	runtime.GC()
// 	time.Sleep(3 * time.Second)
// 	// runtime.GC()
// 	// time.Sleep(1 * time.Second)
// 	// 	runtime.GC()
// 	// 	time.Sleep(1 * time.Second)
// 	// 	runtime.GC()
// 	// 	time.Sleep(1 * time.Second)
// 	// 	runtime.GC()
// 	// 	time.Sleep(1 * time.Second)
// }
