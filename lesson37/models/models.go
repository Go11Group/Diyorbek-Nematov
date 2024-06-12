package models

type User struct {
	ID         string 
	FirstName  string
	LastName   string
	Age        int
	Gender     string
	Nation     string
	Field      string
	ParentName string
	City       string
}

type Filter struct {
	Age           int
	Gender        string
	Nation        string
	Field         string
	Limit, Offset int
}
