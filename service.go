// Service
package service

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Handler interface {
	OnStart()
	OnStop()
}

type Service struct {
	handlers []Handler
	wg       *sync.WaitGroup
	sig      chan os.Signal
}

func New() *Service {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	return &Service{
		wg:  new(sync.WaitGroup),
		sig: sig,
	}
}

func (t *Service) Add(handler Handler) {
	t.handlers = append(t.handlers, handler)
}

func (t *Service) Run() {

	for i := range t.handlers {
		go t.start(t.handlers[i])
	}

	//kill signal
	<-t.sig

	for i := range t.handlers {
		go t.stop(t.handlers[i])
	}

	t.wg.Wait()
}

func (t *Service) start(h Handler) {
	h.OnStart()
	t.wg.Add(1)
}

func (t *Service) stop(h Handler) {
	h.OnStop()
	defer t.wg.Done()
}
