package watcher

import (
	"context"

	"github.com/rtfreedman/automator/cleanup"
)

// Watch watches the postgres database for scripts to trigger
// I know this is basically cron with extra steps, but I'm having fun damnit
func Watch(ctx context.Context) {
	var finished = make(chan bool, 1)
	// register ourselves with cleanup so we don't have a hard-kill
	cleanup.Register(finished)
	for {
		select {
		case <-ctx.Done():
			// indicate we're done with whatever script we were working on
			finished <- true
			return
		default:
		}
		return
	}
}
