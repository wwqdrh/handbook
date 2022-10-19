package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"pprofx/animal"
)

var mode = flag.String("mode", "0", "类型")

func main() {
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		switch *mode {
		case "0":
			for _, v := range animal.AllAnimals {
				v.Live()
			}
		case "1":
			for _, v := range animal.CPUAnimal {
				v.Live()
			}
		case "2":
			for _, v := range animal.MemoryAnimal {
				v.Live()
			}
		case "3":
			for _, v := range animal.AllocsAnimal {
				v.Live()
			}
		case "4":
			for _, v := range animal.GroutineAnimal {
				v.Live()
			}
		case "5":
			for _, v := range animal.BlockAnimal {
				v.Live()
			}
		}

		time.Sleep(10 * time.Second)
	}
}
