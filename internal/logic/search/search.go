package search

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/meilisearch/meilisearch-go"
)

type sSearch struct {
	meilisearchClient meilisearch.ServiceManager
}

func init() {
	ctx := gctx.GetInitCtx()
	data := g.Cfg().MustData(ctx)
	searchConfig, ok := data["search"].(g.Map)
	if !ok {
		panic("Search 配置不存在！")
	}
	meilisearchConfig, ok := searchConfig["meilisearch"].(g.Map)
	if !ok {
		panic("meilisearch 配置不存在！")
	}
	endpoint, ok := meilisearchConfig["endpoint"].(string)
	if !ok {
		panic("meilisearch 配置不存在！")
	}
	apiKey, ok := meilisearchConfig["apiKey"].(string)
	if !ok {
		panic("meilisearch 配置不存在！")
	}
	client := meilisearch.New(endpoint, meilisearch.WithAPIKey(apiKey))
	service.RegisterSearch(New(client))
}

func New(client meilisearch.ServiceManager) service.ISearch {
	return &sSearch{meilisearchClient: client}
}

func (s *sSearch) AddSentenceToSearch(ctx context.Context, schema *entity.Pending) error {
	_, err := s.meilisearchClient.Index("sentences").AddDocuments(g.Slice{g.Map{
		"uuid":        schema.Uuid,
		"hitokoto":    schema.Hitokoto,
		"from":        schema.From,
		"from_who":    schema.FromWho,
		"creator":     schema.Creator,
		"creator_uid": schema.CreatorUid,
		"reviewer":    schema.Reviewer,
		"commit_from": schema.CommitFrom,
		"created_at":  schema.CreatedAt,
	}})
	return err
}
