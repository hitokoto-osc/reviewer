// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/hitokoto-osc/reviewer/internal/dao/internal"
)

// internalPollUsersDao is internal type for wrapping internal DAO implements.
type internalPollUsersDao = *internal.PollUsersDao

// pollUsersDao is the data access object for table hitokoto_poll_users.
// You can define custom methods on it to extend its functionality as you wish.
type pollUsersDao struct {
	internalPollUsersDao
}

var (
	// PollUsers is globally public accessible object for table hitokoto_poll_users operations.
	PollUsers = pollUsersDao{
		internal.NewPollUsersDao(),
	}
)

// Fill with you ideas below.
