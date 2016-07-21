package httputil2_test

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/zimbatm/httputil2"
)

func ExampleNewTrackingListener() {
	// This WaitGroup is passed around and can be used to make sure everything
	// is finished.
	var done sync.WaitGroup

	// Setup the server
	s := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		}),
	}

	// Listen to connections
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// And track them
	l = httputil2.NewTrackingListener(l, done)

	// Setup signal handling
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	// Close the listener on signal
	go func() {
		<-c
		l.Close()
	}()

	err = s.Serve(l)
	if err != nil {
		panic(err)
	}

	done.Wait()
}
