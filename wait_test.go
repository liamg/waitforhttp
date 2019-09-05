package waitforhttp

import (
	"net/http"
	"testing"
	"time"
)

func TestWaitingForHTTPServer(t *testing.T) {
	server := &http.Server{
		Addr: ":8080",
	}
	waitChan := make(chan error, 1)
	go func() {
		waitChan <- Wait(server, time.Second*10)
		_ = server.Close()
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		t.Fatalf("failed to start server: %s", err)
	}

	if err := <-waitChan; err != nil {
		t.Fail()
	}
}

func TestWaitingForHTTPServerFailsAfterTimeout(t *testing.T) {
	server := &http.Server{
		Addr: ":8080",
	}
	waitChan := make(chan error, 1)
	go func() {
		waitChan <- Wait(server, time.Second*3)
		_ = server.Close()
	}()

	time.Sleep(time.Second * 4)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		t.Fatalf("failed to start server: %s", err)
	}

	if err := <-waitChan; err == nil {
		t.Fail()
	}
}
