package day7

import (
	"context"
	"sync"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

const redrawInterval = 250 * time.Millisecond

func (a *Amp) Write(b []byte) (n int, err error) {
	if a.quiet {
		return 0, nil
	}
	if err := a.Text.Write(string(b)); err != nil {
		return 0, err
	}
	return len(b), nil
}

func AmplifierControlPanel(phases []int) (int, error) {
	tb, err := termbox.New(termbox.ColorMode(terminalapi.ColorMode256))
	if err != nil {
		return -1, err
	}

	ampIDs := []string{"A", "B", "C", "D", "E"}
	amps := Amps(make([]*Amp, len(ampIDs)))

	builder := grid.New()
	for i, id := range ampIDs {
		amp, err := newAmp(id)
		if err != nil {
			return -1, err
		}
		amps[i] = amp
		builder.Add(grid.ColWidthPerc(100/len(ampIDs), grid.Widget(amp.Text, container.Border(linestyle.Light))))
	}
	gridOpts, err := builder.Build()
	if err != nil {
		return -1, err
	}

	c, err := container.New(tb, gridOpts...)
	if err != nil {
		return -1, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}
	go func() {
		wg := &sync.WaitGroup{}
		amps.Run(wg, phases)
		wg.Wait()
		cancel()
	}()

	termdash.Run(ctx, tb, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval))

	return amps[len(amps)-1].lastOutput, nil
}
