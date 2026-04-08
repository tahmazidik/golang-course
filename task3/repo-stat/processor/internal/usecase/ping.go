package usecase

import "context"

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

func (u *Ping) Execute(_ context.Context) string {
	return "pong"
}
