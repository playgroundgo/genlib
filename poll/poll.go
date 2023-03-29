package poll

import (
	"context"
	"time"
)

// PollWhile is used to continuously execute the given function whith a given interval until it is
// evaluating to false.
func PollWhile(ctx context.Context, interval time.Duration, f func() bool) error {
	doneCh := make(chan error)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	go func() {
		defer close(doneCh)
		for {
			select {
			case <-ctx.Done():
				doneCh <- ctx.Err()
				return
			case <-ticker.C:
				if !f() {
					return
				}
			}
		}
	}()
	return <-doneCh
}
