package models

type User struct {
	ID       int
	Login    string
	Password string
}

type Calculation struct {
	ID         int
	UserID     int
	Expression string
	Result     float64
}
