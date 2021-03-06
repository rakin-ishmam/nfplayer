package task

import (
	"context"
	"log"

	"github.com/rakin-ishmam/nfplayer/store"
)

type worker struct {
	wid   int
	id    chan int
	ctx   context.Context
	store store.Team
}

func (w *worker) todo(id int) {
	w.id <- id
}

func (w *worker) gracefulStop() {
	close(w.id)
}

func (w *worker) run() <-chan result {
	c := make(chan result, 200)

	go func() {
		defer close(c)
		defer func() {
			log.Println("stop worker", w.wid)
		}()

		for w.id != nil {
			select {
			case id, ok := <-w.id:
				if !ok {
					w.id = nil
					continue
				}

				log.Println("fetching id", id)
				mt, err := w.store.ByID(id)
				if err == store.ErrNotFound {
					continue
				}

				res := result{}

				if err != nil {
					res.err = err
					c <- res
					continue
				}

				log.Printf("found team=%s, id=%d", mt.Name, mt.ID)
				res.team = *mt

				c <- res
			case <-w.ctx.Done():
				log.Println("stopping worker...")
				return
			}
		}
	}()

	return c
}

func newWorker(ctx context.Context, store store.Team, id int) *worker {
	return &worker{
		wid:   id,
		id:    make(chan int, 500),
		ctx:   ctx,
		store: store,
	}
}
