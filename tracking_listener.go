package httputil2

import (
	"net"
	"sync"
)

type trackingListener struct {
	net.Listener
	wg     sync.WaitGroup
	closed func()
}

// A net.Listener that tracks the livelyhood of the connections such that
// the wg internal counter will go back to it's initial value once the
// listener and all it's issued net.Conn are closed.
//
// This is useful for gracefully shutting down a server where first new
// connections are stoppped being accepted and then all the client connections
// are being shutdown as requests terminate.
//
// Note that net/http.Server only provides HTTP/2 when ListenAndServeTLS is
// called directly (whereas here you would use the Serve(l) function).
func NewTrackingListener(l net.Listener, wg sync.WaitGroup) net.Listener {
	var once sync.Once

	wg.Add(1)

	return &trackingListener{
		Listener: l,
		wg:       wg,
		closed: func() {
			once.Do(wg.Done)
		},
	}
}

func (l *trackingListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		l.closed()
		return nil, err
	}
	return newTrackedConn(conn, l.wg), err
}

func (l *trackingListener) Close() error {
	err := l.Listener.Close()
	l.closed()
	return err
}

type trackedConn struct {
	net.Conn
	closed func()
}

func newTrackedConn(c net.Conn, wg sync.WaitGroup) net.Conn {
	var once sync.Once

	wg.Add(1)

	return &trackedConn{
		Conn: c,
		closed: func() {
			once.Do(wg.Done)
		},
	}
}

func (c *trackedConn) Read(b []byte) (n int, err error) {
	n, err = c.Conn.Read(b)
	if err != nil {
		c.closed()
	}
	return n, err
}

func (c *trackedConn) Write(b []byte) (n int, err error) {
	n, err = c.Conn.Write(b)
	if err != nil {
		c.closed()
	}
	return n, err
}

func (c *trackedConn) Close() error {
	err := c.Conn.Close()
	c.closed()
	return err
}
