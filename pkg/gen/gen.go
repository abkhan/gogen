package gen

import (
	"errors"
	"fmt"
	"time"
)

// how it should be;
// - use of a service + a goroutine
// - use of channels
// - an exec to run to yeild
// - a run that returns a chennel and calls the generator over a over
//		to allow range over the channel, but the channel should not be buffered

type generator struct {
	started bool
	datach  chan any
	indch   chan struct{}
	gf      func(<-chan struct{}, chan<- any, any) error
	funcdet any
}

func New(name string, details any, f func(<-chan struct{}, chan<- any, any) error) *generator {
	return &generator{
		datach:  make(chan any),
		indch:   make(chan struct{}),
		funcdet: details,
		gf:      f,
	}
}

func (g *generator) Exec() (any, error) {
	if g.indch == nil || g.datach == nil {
		return nil, errors.New("channel closed")
	}

	// start the go-routine that runs the generator
	if !g.started {
		go g.gf(g.indch, g.datach, g.funcdet)
		time.Sleep(1 * time.Millisecond)
		g.started = true
		g.indch <- struct{}{}
	}

	g.indch <- struct{}{}
	return <-g.datach, nil
}

func (g *generator) Run() (chan any, error) {
	if g.indch == nil || g.datach == nil {
		return nil, errors.New("channel closed")
	}

	// start the go-routine that runs the generator
	if !g.started {
		fmt.Println("running go-routine inn Run")
		go g.gf(g.indch, g.datach, g.funcdet)
		time.Sleep(1 * time.Millisecond)
		g.started = true
		g.indch <- struct{}{}
	}

	retch := make(chan any)
	go func() {
		defer close(retch)

		for {

			// need to see if the channel is alive
			select {
			case _, alive := <-g.datach:
				if !alive {
					fmt.Println("failed check for channel alive")
					return
				}
			default:
				fmt.Println("channel is alive")
			}

			g.indch <- struct{}{}
			fmt.Println("gen: case of sending indicator")

			val, ok := <-g.datach
			if !ok {
				break
			}
			fmt.Println("gen: case of data")
			retch <- val
		}

	}()

	return retch, nil
}
