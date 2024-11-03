package entity


type {{.StructName}} struct{ 
	{{- range $idx, $field := .FieldMap}}
	{{$field.Name}} {{$field.TypeName}} `json:"{{$field.JsonTag}}" orm:"{{$field.Orm}}" desc:"{{$field.Desc}}"`
	{{- end}}
}