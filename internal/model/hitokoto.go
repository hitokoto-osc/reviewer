package model

type ModifySentenceParams struct {
	Hitokoto string `json:"hitokoto" dc:"句子内容" v:"length:1,1000#句子长度应该在 1-1000 之间"`
	Type     string `json:"type" dc:"句子类型" v:"length:1,100#句子类型长度应该在 1-100 之间"`
	From     string `json:"from" dc:"句子来源" v:"length:1,100#来源长度应该在 1-100 之间"`
	FromWho  string `json:"from_who" dc:"句子来源者" v:"length:1,100#来源者长度应该在 1-100 之间"`
	Creator  string `json:"creator" dc:"句子创建者" v:"length:1,100#创建者长度应该在 1-100 之间"`
}
