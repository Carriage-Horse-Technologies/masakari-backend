package main

// 0.7以上で足切り
const threshold = 0.7

func attackServer(msg string) {
	score := ScoreMsg(msg)
	Benchmark(score)
	return
}
