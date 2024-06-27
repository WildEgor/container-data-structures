package main

import (
	"context"
	"fmt"
	"github.com/WildEgor/container-data-structures/pkg/queue"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Job struct {
	Msg []byte
}

type IWorker interface {
	Do(job *Job)
}

type Dispatcher struct {
	sem    chan struct{} // limit workers
	queue  *queue.SyncQueue[*Job]
	worker IWorker
	wg     sync.WaitGroup
}

func NewDispatcher(worker IWorker, mw int) *Dispatcher {
	return &Dispatcher{
		sem:    make(chan struct{}, mw),
		queue:  queue.NewSyncQueue[*Job](),
		worker: worker,
	}
}

func (d *Dispatcher) Start(ctx context.Context) {
	d.wg.Add(1)
	go d.loop(ctx)
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

func (d *Dispatcher) Add(job *Job) {
	d.queue.Enqueue(job)
}

func (d *Dispatcher) stop() {
	d.wg.Done()
}

func (d *Dispatcher) loop(ctx context.Context) {
	var wg sync.WaitGroup

Loop:
	for {
		select {
		default:
			if d.queue.Size() == 0 {
				break Loop
			}

			job, ok := d.queue.Dequeue()
			if !ok {
				break
			}

			wg.Add(1)
			d.sem <- struct{}{}

			go func(job *Job) {
				defer func() { <-d.sem }()
				d.worker.Do(job)
			}(job)
		case <-ctx.Done():
			// done do jobs before break
			wg.Wait()
			break Loop
		}
	}
	d.stop()
}

var _ IWorker = (*Printer)(nil)

// Printer is a dummy worker that just prints msg
type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

// Do waits for a few seconds and print a data
func (p *Printer) Do(j *Job) {
	t := time.NewTimer(time.Duration(rand.Intn(5)) * time.Second)
	defer t.Stop()
	<-t.C
	fmt.Println(j.Msg)
}

func main() {
	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		<-sigCh
		cancel()
	}()

	p := NewPrinter()
	d := NewDispatcher(p, 10)

	d.Start(ctx)
	for i := 0; i < 20; i++ {
		var msg = []byte{byte(rand.Intn(10))}

		d.Add(&Job{
			Msg: msg,
		})
	}
	d.Wait()
}
