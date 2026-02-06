package main

import "fmt"

type Status int

const (
	Pending    Status = iota // 0
	Processing               // 1
	Completed                // 2
	Failed                   // 3
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "PENDING"
	case Processing:
		return "PROCESSING"
	case Completed:
		return "COMPLETED"
	case Failed:
		return "FAILED"
	default:
		return "UNKNOWN"
	}
}

func main() {
	s := Processing
	fmt.Printf("Status: %d (%s)\n", s, s)
	// Status: 1 (PROCESSING)
}
