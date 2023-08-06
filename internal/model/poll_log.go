package model

import "github.com/hitokoto-osc/reviewer/internal/model/entity"

type PollLogWithSentence struct {
	entity.PollLog
	Sentence *HitokotoV1Schema // Nullable
}
