package watcher

import (
	"context"
	"database/sql"
	"math/rand"
	"os"
	"time"

	"github.com/rtfreedman/automator/cleanup"
)

var db *sql.DB

// Connect connects the watcher to postgres
func Connect() (err error) {
	user := os.Getenv("AUTOMATORPGUSER")
	password := os.Getenv("AUTOMATORPGPASS")
	dbname := os.Getenv("AUTOMATORPGDB")
	host := os.Getenv("AUTOMATORPGHOST")
	port := os.Getenv("AUTOMATORPGPORT")
	connStr := "user=" + user + " dbname=" + dbname + " host=" + host + " port=" + port + " password=" + password + " sslmode=disable"
	if db, err = sql.Open("postgres", connStr); err != nil {
		return
	}
	err = db.Ping()
	return
}

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
			break
		default:
		}
		getJobs()
	}
}

func getJobs() {
	// sleep between 0 seconds and 8 minutes
	time.Sleep(time.Duration(rand.Int31n(8*60)) * time.Second)

	return
}
