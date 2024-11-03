package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CardFriendDao is the data access object for table report.
type CardFriendDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CardFriendColumns // columns contains all the column names of Table for convenient usage.
}

// CardFriendColumns defines and stores column names for table report.
type CardFriendColumns struct {
	Id          string // id
	Region      string // 大区##
	Level       string // 段位##
	State       string // 状态 已读/未读##
	CreatorUid  string // 受邀请人uid##
	CreatorName string // 受邀请人召唤师名##
	InviteeUid  string // 受邀请人uid##
	InviteeName string // 受邀请人召唤师名##
	CardInfo    string // 卡片快照信息##
	CreatedAt   string // 创建时间##
	UpdatedAt   string // 更新时间##
	DeletedAt   string // 删除时间##
}

// reportColumns holds the columns for table report.
var reportColumns = CardFriendColumns{
	Id:          "id",
	Region:      "region",
	Level:       "level",
	State:       "state",
	CreatorUid:  "creator_uid",
	CreatorName: "creator_name",
	InviteeUid:  "invitee_uid",
	InviteeName: "invitee_name",
	CardInfo:    "card_info",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

var (
	// CardFriend is globally public accessible object for table report operations.
	CardFriend = NewCardFriendDao()
)

// NewCardFriendDao creates and returns a new DAO object for table data access.
func NewCardFriendDao() *CardFriendDao {
	return &CardFriendDao{
		group:   "default",
		table:   "report",
		columns: reportColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CardFriendDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CardFriendDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CardFriendDao) Columns() CardFriendColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CardFriendDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CardFriendDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CardFriendDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
