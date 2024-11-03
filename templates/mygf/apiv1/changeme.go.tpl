package apiv1


import (
	"github.com/gogf/gf/v2/frame/g"
	"{{.GoModName}}/model"
	"{{.GoModName}}/model/entity"
)

// 新增
type {{.StructName}}AddReq struct {
	g.Meta `path:"/{{.TableName}}" tags:"{{.TableName}}" method:"post" summary:"创建{{.TableName}}"`
	entity.{{.StructName}}
}

type {{.StructName}}AddRes struct {
}

// 更新
type {{.StructName}}UpdateReq struct {
	g.Meta `path:"/{{.TableName}}" tags:"{{.TableName}}" method:"put" summary:"更新{{.TableName}}"`
	entity.{{.StructName}}
}

type {{.StructName}}UpdateRes struct {
}

{{- range $idx, $logic := .UpdateLogics}}
type {{$.StructName}}Update{{$logic.Name}}Req struct{
	g.Meta `path:"/{{$.TableName}}/update_{{formatFieldName $logic.Name "CaseSnake"}}" method:"put"`
	// update fields
	{{- range $idx, $data := $logic.Data}}
		{{if eq $data.Value "$input"}}
			{{- formatFieldName $data.Name "CaseCamel"}} string // {{formatFieldName "Test" "CaseSnake"}}
		{{- end}}
	{{- end }}
	// where fields
	{{- range $idx, $data := $logic.Where}}
		{{if eq $data.Value "$input"}}
			{{- formatFieldName $data.Name "CaseCamel"}} string // {{formatFieldName "Test" "CaseSnake" -}}
		{{- end -}}
	{{- end -}}
}
{{- end}}


// 列表查询
type {{.StructName}}ListReq struct {
	g.Meta `path:"/{{.TableName}}/list" tags:"{{.TableName}}" method:"get" summary:"查询{{.TableName}}列表"`
	model.{{.StructName}}ListInputQuery
	model.{{.StructName}}ListInputPage
}

type {{.StructName}}ListRes struct {
	List  []entity.{{.StructName}}
	Total int
}

// 根据id查询
type {{.StructName}}GetByIdReq struct {
	g.Meta `path:"/{{.TableName}}/{id}" tags:"{{.TableName}}" method:"get" summary:"根据id查询{{.TableName}}"`
	Id     int
}

type {{.StructName}}GetByIdRes struct {
	entity.{{.StructName}}
}

