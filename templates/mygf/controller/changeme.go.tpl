package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"{{.GoModName}}/model"
	"{{.GoModName}}/apiv1"
	"{{.GoModName}}/service/{{.TableName}}"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

var (
	{{.StructName}} = c{{.StructName}}{}
)

type c{{.StructName}} struct{}

func (c *c{{.StructName}}) Add(ctx context.Context, req *apiv1.{{.StructName}}AddReq) (res *apiv1.{{.StructName}}AddRes, err error) {
	in := &model.{{.StructName}}AddInput{}
	in.{{.StructName}} = req.{{.StructName}}
	err = {{.TableName}}.{{.StructName}}().Add(ctx, in)
	return
}

func check{{.StructName}}UpdateReq(req *apiv1.{{.StructName}}UpdateReq) error {
	// 补充字段检查逻辑
	return nil
}

func (c *c{{.StructName}}) Update(ctx context.Context, req *apiv1.{{.StructName}}UpdateReq) (res *apiv1.{{.StructName}}UpdateRes, err error) {
	if err = check{{.StructName}}UpdateReq(req); err != nil {
		return nil, err
	}
	in := &model.{{.StructName}}UpdateInput{}
	in.{{.StructName}} = req.{{.StructName}}
	// 这里根据情况修改查询条件，注意在条件中控制权限
	// uid := user_context.Context.Get(ctx).User.Uid
	// in.{{.StructName}}.Uid = uid
	data, _ := {{.TableName}}.{{.StructName}}().GetById(ctx, int(req.Id))
	if data.Id != 0 {
		in.{{.StructName}}.Id = data.Id
	} else {
		in.{{.StructName}}.CreatedAt = gtime.TimestampMilli()
	}

	in.{{.StructName}}.UpdatedAt = gtime.TimestampMilli()

	err = {{.TableName}}.{{.StructName}}().Update(ctx, in)
	return
}

{{- range $idx, $logic := .UpdateLogics}}
func (c *c{{$.StructName}})Update{{$logic.Name}}(ctx context.Context, req *apiv1.{{$.StructName}}Update{{$logic.Name}}Req) (res *apiv1.{{$.StructName}}UpdateRes, err error){
	in := model.{{$.StructName}}Update{{$logic.Name}}Input{}
	gconv.Struct(req,&in.Data)
	gconv.Struct(req,&in.Where)
	_,err = {{$.TableName}}.{{$.StructName}}().Ctx(ctx).Data(in.Data).Where(in.Where).Update()
	return nil,err
}
{{- end}}

func (c *c{{.StructName}}) {{.StructName}}List(ctx context.Context, req *apiv1.{{.StructName}}ListReq) (res apiv1.{{.StructName}}ListRes, err error) {
	in := &model.{{.StructName}}ListInput{}
	in.{{.StructName}}ListInputQuery = req.{{.StructName}}ListInputQuery
	in.{{.StructName}}ListInputPage = req.{{.StructName}}ListInputPage
	list, err := {{$.TableName}}.{{.StructName}}().List(ctx, in)
	if err != nil {
		return res, err
	}
	res.List = list.List
	res.Total = list.Total
	return res, err
}

func (c *c{{.StructName}}) {{.StructName}}GetById(ctx context.Context, req *apiv1.{{.StructName}}GetByIdReq) (res apiv1.{{.StructName}}GetByIdRes, err error) {
	info, err := {{$.TableName}}.{{.StructName}}().GetById(ctx, req.Id)
	if err != nil {
		return res, err
	}
	res.{{.StructName}} = info.{{.StructName}}
	return
}
