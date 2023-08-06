// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/hitokoto-osc/reviewer/internal/dao/internal"
)

// internalPollMarkRelationDao is internal type for wrapping internal DAO implements.
type internalPollMarkRelationDao = *internal.PollMarkRelationDao

// pollMarkRelationDao is the data access object for table hitokoto_poll_mark_relation.
// You can define custom methods on it to extend its functionality as you wish.
type pollMarkRelationDao struct {
	internalPollMarkRelationDao
}

var (
	// PollMarkRelation is globally public accessible object for table hitokoto_poll_mark_relation operations.
	PollMarkRelation = pollMarkRelationDao{
		internal.NewPollMarkRelationDao(),
	}
)

// Fill with you ideas below.