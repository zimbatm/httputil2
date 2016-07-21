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

// A net.Listener that tracks the livelyhood of the connections
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
	return NewTrackedConn(conn, l.wg), err
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

func NewTrackedConn(c net.Conn, wg sync.WaitGroup) net.Conn {
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
