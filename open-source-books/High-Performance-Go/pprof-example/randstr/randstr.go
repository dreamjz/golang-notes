package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
	"time"
	"unsafe"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	N       = 1000
)

var (
	profileType = flag.Int("prfType", 0, "profile type: 0 - runtime/pprof, 1 - net/http/pprof, 2 - github.com/pkg/profile")
)

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

func concatString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func main() {
	flag.Parse()
	switch *profileType {
	case 0:
		profileWithRuntime()
	case 1:
		profileWithHttp()
	case 2:
		profileWithPkgProfile()
	}
}

func profileWithRuntime() {
	concatString(N)

	f, err := os.OpenFile("./mem.pprof", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln("create profile error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("close file error:", err)
		}
	}()

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatalln("write profile error:", err)
	}
}

func profileWithHttp() {
	concatString(N)
	go func() {
		fmt.Println("Listening :6060")
		log.Fatalln(http.ListenAndServe(":6060", nil))
	}()
	var wg sync.WaitGroup
	// Block main goroutine
	wg.Add(1)
	wg.Wait()
}

func profileWithPkgProfile() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concatString(N)
}
