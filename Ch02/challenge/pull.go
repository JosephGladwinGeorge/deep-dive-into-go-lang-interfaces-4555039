package main

import "fmt"

type Record struct {
	Key  uint
	Data []byte
}

// Puller pulls a single record from the database.
type Puller interface {
	pull(r *Record)
}

type DB struct{
	Records map[uint]Record
}

type query struct{
	Key uint
}

func (db *DB) pull(r *Record) {
	*r=db.Records[r.Key]
}

func (db *DB) push(r Record){
	db.Records[r.Key]=r
}

func NewDB() *DB{
	return &DB{
		Records: make(map[uint]Record),
	}
}

func main(){
	db:=NewDB()


	r:=Record{
		Key: 2,
		Data: []byte{byte(123),byte(21)},
	}

	db.push(r)


	p:=&Record{
		Key: 2,
	}

	fmt.Println(p)

	db.pull(p)

	fmt.Println(p)
}


