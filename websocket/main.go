package main

import (
	"crypto/sha1"
	"encoding/base32"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
)

// Websocket connection
type Websocket struct {
	conn   net.Conn
	mu     sync.Mutex
	closed bool
}

// Opcodes
const (
	OpcodeContinuation = 0x0
	OpcodeText         = 0x1
	OpcodeBinary       = 0x2
	OpcodeClose        = 0x8
	OpcodePing         = 0x9
	OpcodePong         = 0xA
)

// upgrade an HTTP connection to Websocket
func UpgradeToWebSocket(w http.ResponseWriter, r *http.Request) (*Websocket, error) {
	// Validate websocker upgrade request
	if r.Method != http.MethodGet {
		return nil, errors.New("method must be GET")
	}

	if r.Header.Get("Upgrade") != "websocket" {
		return nil, errors.New("upgrade header must be websocket")
	}
	if !strings.Contains(r.Header.Get("Connection"), "Upgrade") {
		return nil, errors.New("connection header must contain Upgrade")
	}

	wsVersion := r.Header.Get("Sec-Websocket-Vesion")
	if wsVersion != "13" {
		return nil, errors.New("websocket version must be 13")
	}

	key := r.Header.Get("Sec-WebSocket-Key")
	if key == "" {
		return nil, errors.New("missing Sec-Websocket-Key")
	}
	acceptKey := computAcceptKey(key)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		return nil, errors.New("http.ResponseWriter doesn't support hijacking")
	}

	conn, bufrw, err := hijacker.Hijack()
	if err != nil {
		return nil, err
	}

	// Send upgrade response
	response := fmt.Sprintf(
		"HTTP/1.1 101 Switching Protocols\r\n"+
			"Upgrade: websocket\r\n"+
			"Connection: Upgrade\r\n"+
			"Sec-WebSocket-Accept: %s\r\n\r\n",
		acceptKey,
	)
	if _, err := bufrw.Write([]byte(response)); err != nil {
		conn.Close()
		return nil, err
	}

	if err := bufrw.Flush(); err != nil {
		conn.Close()
		return nil, err
	}

	return &Websocket{conn: conn}, nil
}

func computAcceptKey(key string) string {
	const magicString = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	h := sha1.New()
	h.Write([]byte(key + magicString))
	return base32.StdEncoding.EncodeToString(h.Sum(nil))
}
