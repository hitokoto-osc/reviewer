// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Sentence is the golang structure for table sentence.
type Sentence struct {
	Id         int     `json:"id"         ` //
	Uuid       string  `json:"uuid"       ` //
	Hitokoto   string  `json:"hitokoto"   ` //
	Type       string  `json:"type"       ` //
	From       string  `json:"from"       ` //
	FromWho    *string `json:"fromWho"    ` //
	Creator    string  `json:"creator"    ` //
	CreatorUid int     `json:"creatorUid" ` //
	Reviewer   int     `json:"reviewer"   ` //
	CommitFrom string  `json:"commitFrom" ` //
	Assessor   string  `json:"assessor"   ` //
	Owner      string  `json:"owner"      ` //
	CreatedAt  string  `json:"createdAt"  ` //
}
