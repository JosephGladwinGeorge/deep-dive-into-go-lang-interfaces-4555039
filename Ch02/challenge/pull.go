package main

type Record struct {
	Key  uint
	Data []byte
}

// Puller pull a single record from the database.
type Puller interface {
}
