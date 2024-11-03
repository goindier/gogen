package controller

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/goindier/example/apiv1"
	"github.com/goindier/example/model"
	"github.com/goindier/example/service/card_friend"
)

var (
	CardFriend = cCardFriend{}
)

type cCardFriend struct{}

func (c *cCardFriend) Add(ctx context.Context, req *apiv1.CardFriendAddReq) (res *apiv1.CardFriendAddRes, err error) {
	in := &model.CardFriendAddInput{}
	in.CardFriend = req.CardFriend
	err = card_friend.CardFriend().Add(ctx, in)
	return
}

func checkCardFriendUpdateReq(req *apiv1.CardFriendUpdateReq) error {
	// 补充字段检查逻辑
	return nil
}

func (c *cCardFriend) Update(ctx context.Context, req *apiv1.CardFriendUpdateReq) (res *apiv1.CardFriendUpdateRes, err error) {
	if err = checkCardFriendUpdateReq(req); err != nil {
		return nil, err
	}
	in := &model.CardFriendUpdateInput{}
	in.CardFriend = req.CardFriend
	// 这里根据情况修改查询条件，注意在条件中控制权限
	// uid := user_context.Context.Get(ctx).User.Uid
	// in.CardFriend.Uid = uid
	data, _ := card_friend.CardFriend().GetById(ctx, int(req.Id))
	if data.Id != 0 {
		in.CardFriend.Id = data.Id
	} else {
		in.CardFriend.CreatedAt = gtime.TimestampMilli()
	}

	in.CardFriend.UpdatedAt = gtime.TimestampMilli()

	err = card_friend.CardFriend().Update(ctx, in)
	return
}
func (c *cCardFriend) UpdateState(ctx context.Context, req *apiv1.CardFriendUpdateStateReq) (res *apiv1.CardFriendUpdateRes, err error) {
	in := model.CardFriendUpdateStateInput{}
	gconv.Struct(req, &in.Data)
	gconv.Struct(req, &in.Where)
	_, err = card_friend.CardFriend().Ctx(ctx).Data(in.Data).Where(in.Where).Update()
	return nil, err
}

func (c *cCardFriend) CardFriendList(ctx context.Context, req *apiv1.CardFriendListReq) (res apiv1.CardFriendListRes, err error) {
	in := &model.CardFriendListInput{}
	in.CardFriendListInputQuery = req.CardFriendListInputQuery
	in.CardFriendListInputPage = req.CardFriendListInputPage
	list, err := card_friend.CardFriend().List(ctx, in)
	if err != nil {
		return res, err
	}
	res.List = list.List
	res.Total = list.Total
	return res, err
}

func (c *cCardFriend) CardFriendGetById(ctx context.Context, req *apiv1.CardFriendGetByIdReq) (res apiv1.CardFriendGetByIdRes, err error) {
	info, err := card_friend.CardFriend().GetById(ctx, req.Id)
	if err != nil {
		return res, err
	}
	res.CardFriend = info.CardFriend
	return
}
