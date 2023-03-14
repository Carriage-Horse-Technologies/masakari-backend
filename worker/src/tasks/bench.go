package tasks

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sync"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

type TargeterType int

const (
	Login TargeterType = iota
	Schedule
)

type MetricsMutex struct {
	sync.Mutex
	vegeta.Metrics
}

type ProductTargeter struct {
	ttype    TargeterType
	targeter vegeta.Targeter
}
type ScheduleTargeter struct {
	Method string
	URL    string
	Header http.Header
}

// 予定取得リクエスト
func NewScheduleTargeter() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		target := ScheduleTargeter{
			Method: http.MethodGet,
			URL:    "http://legacy-stack.yukinissie.com/random.php",
			Header: map[string][]string{
				"Content-Type": {"application/json"},
			},
		}

		tgt.Method = target.Method
		tgt.URL = target.URL
		tgt.Header = target.Header

		return nil
	}
}

func Benchmark(count int) {
	// rate := flag.Int("rate", int(30.0*rateRatio), "Number of requests per time unit [0 = infinity] (default 10/1s)")
	rate := 100 + count
	// duration := flag.Int("duration", 3, "Duration of the test [0 = forever]")
	duration := 1
	flag.Parse()

	targeters := []*ProductTargeter{

		{
			// 予定取得エンドポイント
			Schedule,
			NewScheduleTargeter(),
		},
	}

	rt := vegeta.Rate{Freq: rate / len(targeters), Per: time.Second}
	dur := time.Duration(duration) * time.Second

	var metrics MetricsMutex

	var wg sync.WaitGroup
	for _, t := range targeters {
		wg.Add(1)
		t := t
		go func() {
			defer wg.Done()
			for res := range vegeta.NewAttacker().Attack(t.targeter, rt, dur, string(t.ttype)) {
				if res == nil || res.Error == vegeta.ErrNoTargets.Error() {
					continue
				}

				metrics.Lock()
				metrics.Add(res)
				metrics.Unlock()
			}
		}()
	}

	wg.Wait()
	metrics.Close()

	// 結果をJSONフォーマットにする
	_, err := json.Marshal(metrics)
	if err != nil {
		log.Println(err)
	}
	return
}
