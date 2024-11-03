package {{.TableName}}

import (
	"{{.GoModName}}/model"
	"{{.GoModName}}/dao"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	// s{{.StructName}} is service struct of module {{.StructName}}.
	s{{.StructName}} struct{}
)

var (
	// ins{{.StructName}} is the instance of service User.
	ins{{.StructName}} = s{{.StructName}}{}
)

// {{.StructName}} returns the interface of {{.StructName}} service.
func {{.StructName}}() *s{{.StructName}} {
	return &ins{{.StructName}}
}

func (s *s{{.StructName}}) Add(ctx context.Context, input *model.{{.StructName}}AddInput) error {
	input.Id = 0
	input.CreatedAt = gtime.TimestampMilli()
	input.UpdatedAt = gtime.TimestampMilli()
	_, err := dao.{{.StructName}}.Ctx(ctx).Save(input)
	if err != nil {
	}
	return err
}

func (s *s{{.StructName}}) Update(ctx context.Context, input *model.{{.StructName}}UpdateInput) error {
	_, err := dao.{{.StructName}}.Ctx(ctx).Save(input)
	if err != nil {
	}
	return err
}


func (s *s{{.StructName}}) Ctx(ctx context.Context) *gdb.Model {
	return dao.{{.StructName}}.Ctx(ctx)
}

// 根据id更新状态字段
// 根据id和状态更新name字段
// 更新state为常量
// 更新state字段和updated_at字段



func (s *s{{.StructName}}) List(ctx context.Context, req *model.{{.StructName}}ListInput) (ret model.{{.StructName}}ListOutput, err error) {
	count, err := dao.{{.StructName}}.Ctx(ctx).OmitEmpty().
		Where(req.{{.StructName}}ListInputQuery).Count()
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
	err = dao.{{.StructName}}.Ctx(ctx).OmitEmpty().
		Where(req.{{.StructName}}ListInputQuery).
		Offset(offset).Limit(req.PageSize).
		OrderDesc("updated_at").
		Scan(&ret.List)
	if err != nil {
		return ret, err
	}
	ret.Total = count
	return ret, nil
}


func (s *s{{.StructName}}) GetById(ctx context.Context, id int) (ret model.{{.StructName}}GetByIdOutput, err error) {
	err = dao.{{.StructName}}.Ctx(ctx).Where(g.Map{"id": id}).Scan(&ret)
	return ret, err
}