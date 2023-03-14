package main

import (
	"math"

	"notchman.tech/masakari-backend/jsd"
)

const threshold = 0.7

func ScoreMsg(msg string) float64 {
	//事前リストから計算させる
	score := 0.01
	for _, sentence := range sentenceList {
		score = jsd.StringDistance(msg, sentence)
		// NormalizeDistance([]rune(msg), []rune(sentence))
		if score > threshold {
			break
		}
	}
	return score
}

func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

func NormalizeDistance(lhs, rhs []rune) float64 {
	return float64(Distance(lhs, rhs)) / math.Max(float64(len(lhs)), float64(len(rhs)))
}

func Distance(lhs, rhs []rune) int {
	rl1, rl2 := len(lhs), len(rhs)

	// make distance from the root
	costs := make([]int, rl1+1)
	for j := 1; j <= rl1; j++ {
		costs[j] = j
	}

	cost, last, prev := 0, 0, 0
	for i := 1; i <= rl2; i++ {
		costs[0] = i
		last = i - 1
		for j := 1; j <= rl1; j++ {
			prev = costs[j]
			cost = 0
			// previous letter was different
			if lhs[j-1] != rhs[i-1] {
				cost = 1
			}
			if costs[j]+1 < costs[j-1]+1 {
				// changes to lhs should be minimum
				if costs[j]+1 < last+cost {
					costs[j] = costs[j] + 1
				} else {
					costs[j] = last + cost
				}
			} else {
				// changes to rhs should be minimum
				if costs[j-1]+1 < last+cost {
					costs[j] = costs[j-1] + 1
				} else {
					costs[j] = last + cost
				}
			}
			last = prev
		}
	}
	return costs[rl1]
}
