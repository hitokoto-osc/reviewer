// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/hitokoto-osc/reviewer/internal/dao/internal"
)

// internalPollDao is internal type for wrapping internal DAO implements.
type internalPollDao = *internal.PollDao

// pollDao is the data access object for table hitokoto_poll.
// You can define custom methods on it to extend its functionality as you wish.
type pollDao struct {
	internalPollDao
}

var (
	// Poll is globally public accessible object for table hitokoto_poll operations.
	Poll = pollDao{
		internal.NewPollDao(),
	}
)

// Fill with you ideas below.