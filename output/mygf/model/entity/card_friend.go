package entity

type CardFriend struct {
	Id          uint   `json:"id" orm:"id" desc:"id"`
	Region      string `json:"region" orm:"region" desc:"大区##"`
	Level       string `json:"level" orm:"level" desc:"段位##"`
	State       string `json:"state" orm:"state" desc:"状态 已读/未读##"`
	CreatorUid  string `json:"creatorUid" orm:"creator_uid" desc:"受邀请人uid##"`
	CreatorName string `json:"creatorName" orm:"creator_name" desc:"受邀请人召唤师名##"`
	InviteeUid  string `json:"inviteeUid" orm:"invitee_uid" desc:"受邀请人uid##"`
	InviteeName string `json:"inviteeName" orm:"invitee_name" desc:"受邀请人召唤师名##"`
	CardInfo    string `json:"cardInfo" orm:"card_info" desc:"卡片快照信息##"`
	CreatedAt   int64  `json:"createdAt" orm:"created_at" desc:"创建时间##"`
	UpdatedAt   int64  `json:"updatedAt" orm:"updated_at" desc:"更新时间##"`
	DeletedAt   int64  `json:"deletedAt" orm:"deleted_at" desc:"删除时间##"`
}
