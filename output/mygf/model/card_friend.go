package model

import (
	"github.com/goindier/example/model/entity"
)

type CardFriendAddInput struct {
	entity.CardFriend
}

type CardFriendUpdateInput struct {
	entity.CardFriend
}

type CardFriendListInputQuery struct {
	// 添加查询字段
}

type CardFriendListInputPage struct {
	PageNum  int
	PageSize int
}

type CardFriendListInput struct {
	CardFriendListInputQuery
	CardFriendListInputPage
}

type CardFriendListOutput struct {
	List  []entity.CardFriend
	Total int
}
type CardFriendUpdateStateInput struct {
	Data  CardFriendUpdateStateInputUpdate
	Where CardFriendUpdateStateInputWhere
}

type CardFriendUpdateStateInputUpdate struct {
	State string // testUpdatedAt string // test
}

type CardFriendUpdateStateInputWhere struct {
	Id string // testState string // test
}

type CardFriendGetByIdOutput struct {
	entity.CardFriend
}
