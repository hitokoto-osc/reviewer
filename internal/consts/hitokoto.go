package consts

import "time"

const (
	HitokotoV1SchemaCacheTime = time.Minute * 10 // HitokotoV1Schema 缓存时间
)

type HitokotoType string

// Ref: https://developer.hitokoto.cn/sentence/#%E5%8F%A5%E5%AD%90%E7%B1%BB%E5%9E%8B-%E5%8F%82%E6%95%B0
const (
	HitokotoTypeAnime      HitokotoType = "a"
	HitokotoTypeComic      HitokotoType = "b"
	HitokotoTypeGame       HitokotoType = "c"
	HitokotoTypeNovel      HitokotoType = "d"
	HitokotoTypeOriginal   HitokotoType = "e"
	HitokotoTypeInternet   HitokotoType = "f"
	HitokotoTypeOther      HitokotoType = "g"
	HitokotoTypeMovie      HitokotoType = "h"
	HitokotoTypePoem       HitokotoType = "i"
	HitokotoTypeNCM        HitokotoType = "j"
	HitokotoTypePhilosophy HitokotoType = "k"
	HitokotoTypeJoke       HitokotoType = "l"
)

type HitokotoStatus string

const (
	HitokotoStatusPending  HitokotoStatus = "pending"
	HitokotoStatusApproved HitokotoStatus = "approved"
	HitokotoStatusRejected HitokotoStatus = "rejected"
)
