// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/hitokoto-osc/reviewer/internal/dao/internal"
)

// internalPollMarkDao is internal type for wrapping internal DAO implements.
type internalPollMarkDao = *internal.PollMarkDao

// pollMarkDao is the data access object for table hitokoto_poll_mark.
// You can define custom methods on it to extend its functionality as you wish.
type pollMarkDao struct {
	internalPollMarkDao
}

var (
	// PollMark is globally public accessible object for table hitokoto_poll_mark operations.
	PollMark = pollMarkDao{
		internal.NewPollMarkDao(),
	}
)

// Fill with you ideas below.