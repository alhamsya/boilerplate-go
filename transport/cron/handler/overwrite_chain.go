package cronHandler

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/alhamsya/boilerplate-go/lib/managers/custom_log"
	"github.com/robfig/cron/v3"
)

// DelayIfStillRunning serializes jobs, delaying subsequent runs until the
// previous one is complete. Jobs running after a delay of more than a minute
// have the delay logged at Info.
func DelayIfStillRunning(ctx context.Context, name string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var mu sync.Mutex
		return cron.FuncJob(func() {
			start := time.Now()
			mu.Lock()
			defer mu.Unlock()
			if dur := time.Since(start); dur > time.Minute {
				customLog.WarnF("[CRON] [DELAY] scheduler name %s: %v", name, dur)
			}
			j.Run()
		})
	}
}

// SkipIfStillRunning skips an invocation of the Job if a previous invocation is
// still running. It logs skips to the given logger at Info level.
func SkipIfStillRunning(ctx context.Context, name string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var ch = make(chan struct{}, 1)
		ch <- struct{}{}
		return cron.FuncJob(func() {
			select {
			case v := <-ch:
				j.Run()
				ch <- v
			default:
				customLog.WarnF("[CRON] [SKIP] scheduler name %s", name)
			}
		})
	}
}

// Recover panics in wrapped jobs and log them with the provided logger.
func Recover(ctx context.Context, name string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer func() {
				if r := recover(); r != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					customLog.ErrorF("[CRON] [PANIC] scheduler name %s: %v | stack: %s", name, err, "...\n"+string(buf))
				}
			}()
			j.Run()
		})
	}
}
