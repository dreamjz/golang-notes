package chapter_14

import "testing"

func BenchmarkChannelSync(b *testing.B) {
	channelSync(b.N)
}

func BenchmarkChannelBuffered(b *testing.B) {
	channelBuffered(b.N)
}
