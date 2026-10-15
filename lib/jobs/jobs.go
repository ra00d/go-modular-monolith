// Package jobs provides a minimal fx-lifecycle background job runner.
// Jobs run on fixed intervals; cron expressions are deliberately not
// supported yet.
// ponytail: interval tickers cover current sweeps; swap in a cron parser
// (e.g. robfig/cron) when schedules become irregular or per-env configurable.
package jobs

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	"go.uber.org/fx"
)

// Job is a named unit of recurring work.
type Job struct {
	Name  string
	Every time.Duration
	Run   func(ctx context.Context) error
}

type Runner struct {
	mu   sync.Mutex
	jobs []Job
	stop chan struct{}
	wg   sync.WaitGroup
	log  *glog.Logger
}

// NewRunner builds the runner; jobs register via Register before Start.
func NewRunner() *Runner {
	return &Runner{stop: make(chan struct{}), log: glog.New()}
}

// Module wires the runner into fx and hooks Start/Stop into the lifecycle.
var Module = fx.Module("jobs",
	fx.Provide(NewRunner),
	fx.Invoke(func(lc fx.Lifecycle, r *Runner) {
		lc.Append(fx.Hook{OnStart: r.Start, OnStop: r.Stop})
	}),
)

// Register adds jobs to the runner. Safe before Start only.
func (r *Runner) Register(jobs ...Job) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.jobs = append(r.jobs, jobs...)
}

// Start launches one goroutine per job (fx lifecycle hook).
func (r *Runner) Start(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, j := range r.jobs {
		job := j
		if job.Every <= 0 || job.Run == nil {
			continue
		}
		r.wg.Add(1)
		go func() {
			defer r.wg.Done()
			t := time.NewTicker(job.Every)
			defer t.Stop()
			for {
				select {
				case <-r.stop:
					return
				case <-t.C:
					runCtx, cancel := context.WithTimeout(context.Background(), job.Every)
					if err := job.Run(runCtx); err != nil {
						r.log.Error(runCtx, "job failed", "job", job.Name, err)
					}
					cancel()
				}
			}
		}()
	}
	return nil
}

// Stop halts all workers and waits for the in-flight run to finish.
func (r *Runner) Stop(_ context.Context) error {
	close(r.stop)
	r.wg.Wait()
	return nil
}
