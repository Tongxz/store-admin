package main

import (
	"math/rand"
	"os"
	"runtime"
	"store-admin/internal/apiserver"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	apiserver.NewApp()
}
