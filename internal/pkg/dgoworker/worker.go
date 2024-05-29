package dgoworker

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yuyaprgrm/text2speech/pkg/worker"
)

type Pool = worker.Pool[discordgo.Session]
type Worker = worker.Worker[discordgo.Session]

func NewPool() *Pool {
	return worker.NewPool[discordgo.Session]()
}
