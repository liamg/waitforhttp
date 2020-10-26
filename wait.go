package waitforhttp

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

// Wait for the supplied server to start listening and return nil, or return error if timeout is reached
func Wait(server *http.Server, timeout time.Duration) error {

	delay := time.Millisecond * 50
	var attempts uint

	if timeout < delay {
		attempts = 1
	} else {
		attempts = uint((timeout / delay) + 1)
	}

	if server == nil {
		return fmt.Errorf("server cannot be nil")
	}

	if err := retry.Do(
		func() error {
			conn, err := net.Dial("tcp", server.Addr)
			if err != nil {
				return err
			}
			_ = conn.Close()
			return nil
		},
		retry.Attempts(attempts),
		retry.DelayType(
			func(n uint, err error, config *retry.Config) time.Duration {
				return delay
			},
		),
	); err != nil {
		return fmt.Errorf("timed out waiting for server to be ready at %s", server.Addr)
	}

	return nil
}
