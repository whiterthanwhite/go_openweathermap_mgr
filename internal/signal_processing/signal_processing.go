package signal_processing

import (
	"context"
	"os"
	"os/signal"
)

func Processing(cancel context.CancelFunc) {
	sigInter := make(chan os.Signal)
	sigKill := make(chan os.Signal)

	signal.Notify(sigInter, os.Interrupt)
	signal.Notify(sigKill, os.Kill)

	select {
	case <-sigInter:
		cancel()
	case <-sigKill:
		cancel()
	}
}
