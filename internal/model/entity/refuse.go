// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Refuse is the golang structure for table refuse.
type Refuse struct {
	Id         int     `json:"id"         ` //
	Uuid       string  `json:"uuid"       ` //
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
