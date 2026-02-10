package main

// Functional Optional patterns
// For structs with many optional parameters:


type Server struct {
    host     string
    port     int
    timeout  time.Duration
    maxConns int
    tls      *tls.Config
}

// Option - function that modifies Server
type Option func(*Server)

// Factory functions for options
func WithTimeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.timeout = timeout
    }
}

func WithTLS(config *tls.Config) Option {
    return func(s *Server) {
        s.tls = config
    }
}

func WithMaxConnections(max int) Option {
    return func(s *Server) {
        s.maxConns = max
    }
}

// Constructor accepts required parameters and options
func NewServer(host string, port int, opts ...Option) *Server {
    server := &Server{
        host:     host,
        port:     port,
        timeout:  30 * time.Second, // defaults
        maxConns: 100,
    }
    
    // Apply options
    for _, opt := range opts {
        opt(server)
    }
    
    return server
}

// Usage - reads like prose
server := NewServer("localhost", 8080,
    WithTimeout(60*time.Second),
    WithMaxConnections(1000),
    WithTLS(tlsConfig),
)
