package punchguard

import (
	"sync"
	"time"
)

type Async struct {
	service      *sync.WaitGroup
	async        *sync.WaitGroup
	asyncStopped chan Event
	asyncRunning bool
}

func (b *Async) InitAsync() {
	b.service = new(sync.WaitGroup)
	b.async = new(sync.WaitGroup)
	b.asyncStopped = make(chan Event, 1)
	b.asyncRunning = true
}

func (b *Async) closeAsync(mainClose func()) {
	mainClose()
	close(b.asyncStopped)
	b.async = nil
	b.service = nil
}

func (b *Async) MainLoop(mainOpen, main, mainClose func()) StoppedEvent {
	go func() {
		defer b.closeAsync(mainClose)
		mainOpen()
		for b.asyncRunning {
			main()
			<-time.After(50 * time.Millisecond)
		}
		b.async.Wait()
	}()

	return b.asyncStopped
}

func (b *Async) WhenRunningAsync(async func()) {
	if b.asyncRunning {
		go func() {
			defer b.async.Done()
			b.async.Add(1)
			async()
		}()
	}
}

func (b *Async) WhenRunningSync(sync func()) {
	if b.asyncRunning {
		sync()
	}
}

func (b *Async) StartService(service Service) {
	go func() {
		defer b.service.Done()
		b.service.Add(1)
		serviceStopped := service.Start()
		<-serviceStopped
	}()
}

func (b *Async) StopService(service Service) {
	go func() {
		defer b.async.Done()
		b.async.Add(1)
		service.Stop()
	}()
}

func (b *Async) StopAsync() {
	go func() {
		b.service.Wait()
		b.asyncRunning = false
	}()
}
