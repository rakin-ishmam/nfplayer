package task

import (
	"context"
	"log"

	"github.com/rakin-ishmam/nfplayer/store"
)

type loadBalancer struct {
	id        chan int
	ctx       context.Context
	cancel    func()
	store     store.Team
	totWorker int

	ws []*worker
}

func (lb *loadBalancer) req(id int) {
	lb.id <- id
}

func (lb *loadBalancer) closeReq() {
	close(lb.id)
}

func (lb *loadBalancer) run() <-chan result {
	cs := []<-chan result{}

	for i := 0; i < lb.totWorker; i++ {
		lb.ws = append(lb.ws, newWorker(lb.ctx, lb.store))
		cs = append(cs, lb.ws[i].run())
	}

	go lb.listen()

	return merge(cs...)
}

func (lb *loadBalancer) listen() {

	log.Println("start to listen loadbalancer..")
	curw := 0

	for {
		select {
		case id, ok := <-lb.id:
			if !ok {
				lb.id = nil
				for _, w := range lb.ws {
					w.gracefulStop()
				}
				continue
			}
			lb.ws[curw].todo(id)
			curw = lb.nxtWorker(curw)

		case <-lb.ctx.Done():
			log.Println("stopping load balancer..")
			return
		}
	}
}

func (lb loadBalancer) nxtWorker(cur int) int {
	nxt := cur + 1
	if nxt >= len(lb.ws) {
		nxt = 0
	}

	return nxt
}

func (lb *loadBalancer) stop() {
	lb.cancel()
}

func newLoadBalan(ctx context.Context, store store.Team, totWorker int) *loadBalancer {
	nctx, cancel := context.WithCancel(ctx)
	return &loadBalancer{
		id:        make(chan int),
		ctx:       nctx,
		cancel:    cancel,
		store:     store,
		totWorker: totWorker,
	}
}
