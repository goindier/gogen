package card_friend

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/goindier/example/dao"
	"github.com/goindier/example/model"
)

type (
	// sCardFriend is service struct of module CardFriend.
	sCardFriend struct{}
)

var (
	// insCardFriend is the instance of service User.
	insCardFriend = sCardFriend{}
)

// CardFriend returns the interface of CardFriend service.
func CardFriend() *sCardFriend {
	return &insCardFriend
}

func (s *sCardFriend) Add(ctx context.Context, input *model.CardFriendAddInput) error {
	input.Id = 0
	input.CreatedAt = gtime.TimestampMilli()
	input.UpdatedAt = gtime.TimestampMilli()
	_, err := dao.CardFriend.Ctx(ctx).Save(input)
	if err != nil {
	}
	return err
}

func (s *sCardFriend) Update(ctx context.Context, input *model.CardFriendUpdateInput) error {
	_, err := dao.CardFriend.Ctx(ctx).Save(input)
	if err != nil {
	}
	return err
}

func (s *sCardFriend) Ctx(ctx context.Context) *gdb.Model {
	return dao.CardFriend.Ctx(ctx)
}

// 根据id更新状态字段
// 根据id和状态更新name字段
// 更新state为常量
// 更新state字段和updated_at字段

func (s *sCardFriend) List(ctx context.Context, req *model.CardFriendListInput) (ret model.CardFriendListOutput, err error) {
	count, err := dao.CardFriend.Ctx(ctx).OmitEmpty().
		Where(req.CardFriendListInputQuery).Count()
	if err != nil {
		return ret, err
	}

	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 6
	}
	offset := (req.PageNum - 1) * req.PageSize
	err = dao.CardFriend.Ctx(ctx).OmitEmpty().
		Where(req.CardFriendListInputQuery).
		Offset(offset).Limit(req.PageSize).
		OrderDesc("updated_at").
		Scan(&ret.List)
	if err != nil {
		return ret, err
	}
	ret.Total = count
	return ret, nil
}

func (s *sCardFriend) GetById(ctx context.Context, id int) (ret model.CardFriendGetByIdOutput, err error) {
	err = dao.CardFriend.Ctx(ctx).Where(g.Map{"id": id}).Scan(&ret)
	return ret, err
}
