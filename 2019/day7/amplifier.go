package day7

import (
	"fmt"
	"sync"

	"github.com/ianmcmahon/advent/2019/intcode"
	"github.com/mum4k/termdash/widgets/text"
)

type Amps []*Amp

type Amp struct {
	Text        *text.Text
	ID          string
	inch, outch chan int
	lastOutput  int
	quiet       bool
}

func (amps Amps) Run(wg *sync.WaitGroup, phases []int) {
	wg.Add(1)
	for i, amp := range amps {
		// first, wire them together
		go func(i int) {
			wg.Add(1)
			this := amps[i]
			next := amps[(i+1)%len(amps)]
			if !this.quiet {
				this.Text.Write("Waiting for output\n")
			}
			for {
				o, more := <-this.outch
				if !more {
					break
				}
				this.lastOutput = o
				if !this.quiet {
					this.Text.Write(fmt.Sprintf("Got %d, forwarding to %s\n", o, next.ID))
				}
				next.inch <- o
			}
			if !this.quiet {
				this.Text.Write("exiting wire\n")
			}
			wg.Done()
		}(i)

		amp.inch <- phases[i]

		go intcode.Run(program(), amp.inch, amp.outch, amp)
	}
	amps[0].inch <- 0
	wg.Done()
}

func newQuietAmp(id string) (*Amp, error) {
	amp, err := newAmp(id)
	if err != nil {
		return nil, err
	}
	amp.quiet = true
	return amp, nil
}

func newAmp(id string) (*Amp, error) {
	text, err := text.New()
	if err != nil {
		return nil, err
	}
	if err := text.Write(fmt.Sprintf("Amplifier %s\n", id)); err != nil {
		return nil, err
	}
	return &Amp{
		Text:  text,
		ID:    id,
		inch:  make(chan int, 1),
		outch: make(chan int, 1),
	}, nil
}

func AmplifiersHeadless(phases []int) (int, error) {
	ampIDs := []string{"A", "B", "C", "D", "E"}
	amps := Amps(make([]*Amp, len(ampIDs)))

	for i, id := range ampIDs {
		//fmt.Printf("Creating amp: %s\n", id)
		amp, err := newQuietAmp(id)
		if err != nil {
			return -1, err
		}
		amps[i] = amp
	}
	//fmt.Printf("about to run\n")
	wg := &sync.WaitGroup{}
	amps.Run(wg, phases)
	//fmt.Printf("about to wait...\n")
	wg.Wait()
	//fmt.Printf("done waiting\n")

	return amps[len(amps)-1].lastOutput, nil
}
