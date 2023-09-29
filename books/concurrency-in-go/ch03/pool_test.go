package ch03

import (
	"net"
	"testing"
)

func init() {
	wg := startNetworkDaemon()
	wg.Wait()
}

func Test_pool1(t *testing.T) {
	pool1()
}

func BenchmarkNetDaemon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", ":9999")
		if err != nil {
			b.Fatal(err)
		}
		conn.Close()
	}
}
