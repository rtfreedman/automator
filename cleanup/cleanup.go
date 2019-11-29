package cleanup

import "time"

var finished = make(chan chan bool, 1000)

// Done iterates over finished and checks whether everything is finished (channels only populated when they receive the cancel signal)
func Done() {
	close(finished)
	for c := range finished {
		select {
		case <-c:
			// we sleep here because we don't want to take up all cycles by continuously polling for done
			// we need to give some breathing room
			time.Sleep(1 * time.Millisecond)
		default:
			// in the default case we have yet to finish
			finished <- c
		}
	}
}

// Register a channel to indicate when done
func Register(c chan bool) {
	select {
	case finished <- c:
	default:
		panic("not enough space in finished for all services")
	}
}
