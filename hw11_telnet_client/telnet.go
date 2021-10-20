package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

var ErrNoConnection = errors.New("no connection")

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type myTelnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &myTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (tc *myTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", tc.address, tc.timeout)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	tc.conn = conn
	return nil
}

func (tc *myTelnetClient) Close() error {
	if tc.conn != nil {
		if err := tc.conn.Close(); err != nil {
			return fmt.Errorf("close: %w", err)
		}
	}
	return nil
}

func (tc *myTelnetClient) Send() error {
	if tc.conn == nil {
		return ErrNoConnection
	}
	if _, err := io.Copy(tc.conn, tc.in); err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (tc *myTelnetClient) Receive() error {
	if tc.conn == nil {
		return ErrNoConnection
	}
	if _, err := io.Copy(tc.out, tc.conn); err != nil {
		return fmt.Errorf("receive: %w", err)
	}
	return nil
}
