package waitforhttp

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

func Wait(server *http.Server, timeout time.Duration) error {

	delay := timeout / 100

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
		retry.Attempts(100),
		retry.DelayType(
			func(n uint, config *retry.Config) time.Duration {
				return delay
			},
		),
	); err != nil {
		return fmt.Errorf("timed out waiting for server to be ready at %s", server.Addr)
	}

	return nil
}
