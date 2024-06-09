package models


// User struct
type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
}

// Problem struct
type Problem struct {
	ID          int
	UserID      int
	Title       string
	Description string
}

// SolvedProblem struct
type SolvedProblem struct {
	ID        int
	ProblemID int
	UserID    int
	Solution  string
}

type Solution struct {
	ID int
	Username string
	ProplemTitle string
	ProblemDescription string
	ProblemSolution string
}