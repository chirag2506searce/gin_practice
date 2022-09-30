package main

const PORT string = ":8080"

type student struct {
	RollNum   int
	FirstName string
	LastName  string
	Class     int
	Marks     float64
	Contact   int
}

type multipleStudent struct {
	Data []map[string]student
}

var std1 = student{
	RollNum:   1,
	FirstName: "Chirag",
	LastName:  "Gupta",
	Class:     12,
	Marks:     95.0,
	Contact:   9020105494,
}

var std2 = student{
	RollNum:   2,
	FirstName: "John",
	LastName:  "Doe",
	Class:     12,
	Marks:     92.3,
	Contact:   9020410294,
}

var students = map[int]student{
	std1.RollNum: std1,
	std2.RollNum: std2,
}
