package model

type Status int

const (
	Ready Status = iota
	Offline
)

type User struct {
	Username string
	Status   Status
}
