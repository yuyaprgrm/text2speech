package guild

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yuyaprgrm/text2speech/internal/pkg/dgoworker"
	"github.com/yuyaprgrm/text2speech/pkg/collection"
)

type Guild struct {
	dgoPool         *dgoworker.Pool
	readingChannels collection.Set[string]
}

func NewGuild(workers []*discordgo.Session) *Guild {
	dgoPool := dgoworker.NewPool()
	for _, worker := range workers {
		dgoPool.AddWorker(worker)
	}
	return &Guild{
		dgoPool:         dgoPool,
		readingChannels: collection.NewSet[string](),
	}
}
