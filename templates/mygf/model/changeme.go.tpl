package model

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"{{.GoModName}}/model/entity"
)

type {{.StructName}}AddInput struct{ 
	entity.{{.StructName}}
}

type {{.StructName}}UpdateInput struct{ 
	entity.{{.StructName}}
}

type {{.StructName}}ListInputQuery struct {
	// 添加查询字段
}

type {{.StructName}}ListInputPage struct {
	PageNum  int
	PageSize int
}

type {{.StructName}}ListInput struct {
	{{.StructName}}ListInputQuery
	{{.StructName}}ListInputPage
}

type {{.StructName}}ListOutput struct {
	List  []entity.{{.StructName}}
	Total int
}

{{- range $idx, $logic := .UpdateLogics}}
type {{$.StructName}}Update{{$logic.Name}}Input struct{
	Data {{$.StructName}}Update{{$logic.Name}}InputUpdate
	Where {{$.StructName}}Update{{$logic.Name}}InputWhere
}

type {{$.StructName}}Update{{$logic.Name}}InputUpdate struct{
	{{- range $idx, $data := $logic.Data}}
		{{- formatFieldName $data.Name "CaseCamel"}} string // {{formatFieldName "Test" "CaseSnake"}}
	{{- end }}
}

type {{$.StructName}}Update{{$logic.Name}}InputWhere struct{
	{{- range $idx, $data := $logic.Where}}
		{{- formatFieldName $data.Name "CaseCamel"}} string // {{formatFieldName "Test" "CaseSnake"}}
	{{- end }}
}
{{- end}}


type {{.StructName}}GetByIdOutput struct {
	entity.{{.StructName}}
}