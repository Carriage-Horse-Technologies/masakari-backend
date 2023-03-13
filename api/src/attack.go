package main

// 0.7以上で足切り
const threshold = 0.7

func attackServer(msg string) {
	score := 0.0
	//事前リストから計算させる
	for _, sentence := range sentenceList {
		score = NormalizeDistance([]rune(msg), []rune(sentence))
		if score > threshold {
			break
		}
	}
	Benchmark(score)
}
