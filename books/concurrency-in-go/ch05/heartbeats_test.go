package ch05

import (
	"testing"
	"time"
)

func Test_heartbeat1(t *testing.T) {
	heartbeat1()
}

func Test_heartbeat2(t *testing.T) {
	heartbeat2()
}

func Test_heartbeat3(t *testing.T) {
	heartbeat3()
}

func Test_doWork1(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	nums := []int{1, 2, 3, 4}
	res := doWork1(done, nums)

	for _, v := range nums {
		select {
		case n, ok := <-res:
			if !ok {
				return
			}
			if n != v {
				t.Errorf("get: %v, want: %v", n, v)
			}
		case <-time.After(time.Second):
			t.Fatal("test time out")
		}
	}
}

func Test_doWork2(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	nums := []int{1, 2, 3, 4}
	hb, res := doWork2(done, nums)

	<-hb

	for _, v := range nums {
		select {
		case n, ok := <-res:
			if !ok {
				return
			}
			if n != v {
				t.Errorf("get: %v, want: %v", n, v)
			}
		}
	}
}

func Test_doWork3(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	nums := []int{1, 2, 3, 4}
	timeout := 2 * time.Second
	hb, res := doWork3(done, timeout/2, nums)

	<-hb // 进入处理循环

	for _, v := range nums {
		select {
		case n, ok := <-res:
			if !ok {
				return
			}
			if n != v {
				t.Errorf("get: %v, want: %v", n, v)
			}
		case <-hb:
		case <-time.After(timeout):
			t.Fatal("text time out")
		}
	}
}
