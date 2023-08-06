// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Pending is the golang structure for table pending.
type Pending struct {
	Id         int     `json:"id"         ` //
	Uuid       string  `json:"uuid"       ` //
	PollStatus int     `json:"pollStatus" ` // 投票状态，默认： 0，未启用投票
	Hitokoto   string  `json:"hitokoto"   ` //
	Type       string  `json:"type"       ` //
	From       string  `json:"from"       ` //
	FromWho    *string `json:"fromWho"    ` //
	Creator    string  `json:"creator"    ` //
	CreatorUid int     `json:"creatorUid" ` //
	Owner      string  `json:"owner"      ` //
	Reviewer   int     `json:"reviewer"   ` //
	CommitFrom string  `json:"commitFrom" ` //
	CreatedAt  string  `json:"createdAt"  ` //
}
