package ch03

import "testing"

func Test_once(t *testing.T) {
	once()
}

func Test_onceDeadlock(t *testing.T) {
	onceDeadlock()
}
