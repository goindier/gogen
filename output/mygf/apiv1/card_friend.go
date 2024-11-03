package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/goindier/example/model"
	"github.com/goindier/example/model/entity"
)

// 新增
type CardFriendAddReq struct {
	g.Meta `path:"/card_friend" tags:"card_friend" method:"post" summary:"创建card_friend"`
	entity.CardFriend
}

type CardFriendAddRes struct {
}

// 更新
type CardFriendUpdateReq struct {
	g.Meta `path:"/card_friend" tags:"card_friend" method:"put" summary:"更新card_friend"`
	entity.CardFriend
}

type CardFriendUpdateRes struct {
}
type CardFriendUpdateStateReq struct {
	g.Meta `path:"/card_friend/update_state" method:"put"`
	// update fields
	State string // test

	// where fields
	Id string // test
}

// 列表查询
type CardFriendListReq struct {
	g.Meta `path:"/card_friend/list" tags:"card_friend" method:"get" summary:"查询card_friend列表"`
	model.CardFriendListInputQuery
	model.CardFriendListInputPage
}

type CardFriendListRes struct {
	List  []entity.CardFriend
	Total int
}

// 根据id查询
type CardFriendGetByIdReq struct {
	g.Meta `path:"/card_friend/{id}" tags:"card_friend" method:"get" summary:"根据id查询card_friend"`
	Id     int
}

type CardFriendGetByIdRes struct {
	entity.CardFriend
}
