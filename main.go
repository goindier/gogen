package main

import (
	// "github.com/gogf/gf/v2/database/gdb"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/imports"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/goindier/gogen/pkg"
	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
)

const cfgFilePath = "cli"

func main() {
	var (
		// err error
		db  gdb.DB
		in  genInput
		ctx = context.Background()
	)

	// 连接数据库
	err := g.Cfg().MustGet(
		ctx, cfgFilePath,
	).Scan(&in)
	if in.JsonCase == "" {
		in.JsonCase = "CamelLower"
	}
	// fieldMap, err := in.DB.TableFields(ctx, tableName)

	// 读取数据库字段信息
	var tempGroup = gtime.TimestampNanoStr()
	gdb.AddConfigNode(tempGroup, gdb.ConfigNode{
		Link: in.Link,
	})
	if db, err = gdb.Instance(tempGroup); err != nil {
		glog.Fatalf(`database initialization failed: %+v`, err)
	}
	in.DB = db
	in.TypeMapping = pkg.DefaultTypeMapping

	// var tableNames []string
	// if in.Tables != "" {
	// 	tableNames = gstr.SplitAndTrim(in.Tables, ",")
	// }

	fieldMap, err := in.DB.TableFields(ctx, in.Tables)
	in.FieldMap = fieldMap

	ps, err := pkg.FindTplFiles("templates")
	if err != nil {
		glog.Warning("find tpl file failed")
		return
	}
	fmt.Println("ps", ps)

	for _, tplFilePath := range ps {
		// 生成结构体文件
		generateByFile(ctx, in, tplFilePath)
	}

	// 替换模版
}

func readFileOS(file string) (name string, b []byte, err error) {
	name = filepath.Base(file)
	b, err = os.ReadFile(file)
	return
}

func generateByFile(ctx context.Context, in genInput, tplFilePath string) {
	// tmpl, err := template.ParseFiles(tplFilePath)
	// if err != nil {
	// 	glog.Error("tpl invalid", err)
	// 	return
	// }
	// template.ParseFiles()
	tmpl, err := template.New("example").Funcs(template.FuncMap{
		"formatFieldName": formatFieldName, // 注册自定义函数
	}).ParseFiles(tplFilePath)
	if err != nil {
		glog.Error("tpl invalid", err)
		return
	}

	// tmpl = tmpl.Funcs(template.FuncMap{
	// 	"formatFieldName": formatFieldName, // 注册自定义函数
	// })

	tplInput := tplInput{
		TableName:    in.Tables,
		GoModName:    in.GoModName,
		UpdateLogics: in.UpdateLogics,
		// FieldMap:  in.FieldMap,
	}
	fList := []StructFieldInfo{}
	sortedNames := sortFieldKeyForDao(in.FieldMap)
	for _, name := range sortedNames {
		v := in.FieldMap[name]
		structInfo := DBFieldToStructFieldInfo(ctx, in, v)
		fList = append(fList, structInfo)
	}
	tplInput.FieldMap = fList
	tplInput.StructName = formatFieldName(in.Tables, FieldNameCaseCamel)

	buffer := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(buffer, filepath.Base(tplFilePath), tplInput)
	if err != nil {
		glog.Error("tpl execute failed", err)
		return
	}
	str := buffer.String()
	glog.Info(str)

	outPath := gstr.ReplaceByMap(tplFilePath, map[string]string{
		"changeme":  in.Tables,
		"templates": "output",
		".tpl":      "",
	})
	outputFile, err := gfile.Create(outPath)
	if err != nil {
		glog.Warning("create file failed", outPath)
		return
	}
	outputFile.Write(buffer.Bytes())
	GoFmt(outPath)
}

type tplInput struct {
	TableName    string
	FieldMap     []StructFieldInfo
	StructName   string
	GoModName    string
	UpdateLogics []UpdateLogic
}

type StructFieldInfo struct {
	Name     string
	TypeName string
	JsonTag  string
	Orm      string
	Desc     string
}

type genInput struct {
	Link              string
	Tables            string
	DB                gdb.DB
	FieldMap          map[string]*gdb.TableField // Table field map.
	StructName        string
	JsonCase          string
	FieldMapping      map[pkg.DBTableFieldName]pkg.CustomAttributeType
	TypeMapping       map[pkg.DBFieldTypeName]pkg.CustomAttributeType
	StdTime           bool
	GJsonSupport      bool
	RemoveFieldPrefix string
	TableName         string
	GoModName         string
	UpdateLogics      []UpdateLogic
}

type UpdateLogic struct {
	Name  string
	Data  []WhereField
	Where []WhereField
}

type WhereField struct {
	Name  string
	Verb  string
	Value string
}

func DBFieldToStructFieldInfo(ctx context.Context, in genInput, field *gdb.TableField) StructFieldInfo {
	localTypeName, err := in.DB.CheckLocalTypeForField(ctx, field.Type, nil)
	if err != nil {
		panic(err)
	}
	localTypeNameStr := string(localTypeName)
	switch localTypeName {
	case gdb.LocalTypeDate, gdb.LocalTypeTime, gdb.LocalTypeDatetime:
		if in.StdTime {
			localTypeNameStr = "time.Time"
		} else {
			localTypeNameStr = "*gtime.Time"
		}

	case gdb.LocalTypeInt64Bytes:
		localTypeNameStr = "int64"

	case gdb.LocalTypeUint64Bytes:
		localTypeNameStr = "uint64"

	// Special type handle.
	case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
		if in.GJsonSupport {
			localTypeNameStr = "*gjson.Json"
		} else {
			localTypeNameStr = "string"
		}
	}
	ret := StructFieldInfo{}
	ret.Name = formatFieldName(field.Name, FieldNameCaseCamel)
	ret.TypeName = localTypeNameStr
	ret.JsonTag = gstr.CaseConvert(field.Name, gstr.CaseTypeMatch(in.JsonCase))
	ret.Orm = field.Name
	ret.Desc = field.Comment
	return ret
}

func generateStructDefinition(ctx context.Context, in genInput) (string, []string) {
	var appendImports []string
	buffer := bytes.NewBuffer(nil)
	array := make([][]string, len(in.FieldMap))
	names := sortFieldKeyForDao(in.FieldMap)
	for index, name := range names {
		var imports string
		field := in.FieldMap[name]
		array[index], imports = generateStructFieldDefinition(ctx, field, in)
		if imports != "" {
			appendImports = append(appendImports, imports)
		}
	}
	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(array)
	tw.Render()
	stContent := buffer.String()
	// Let's do this hack of table writer for indent!
	stContent = gstr.Replace(stContent, "  #", "")
	stContent = gstr.Replace(stContent, "` ", "`")
	stContent = gstr.Replace(stContent, "``", "")
	buffer.Reset()
	buffer.WriteString(fmt.Sprintf("type %s struct {\n", in.StructName))
	buffer.WriteString(stContent)
	buffer.WriteString("}")
	return buffer.String(), appendImports
}

func sortFieldKeyForDao(fieldMap map[string]*gdb.TableField) []string {
	names := make(map[int]string)
	for _, field := range fieldMap {
		names[field.Index] = field.Name
	}
	var (
		i      = 0
		j      = 0
		result = make([]string, len(names))
	)
	for {
		if len(names) == 0 {
			break
		}
		if val, ok := names[i]; ok {
			result[j] = val
			j++
			delete(names, i)
		}
		i++
	}
	return result
}

func formatFieldName(fieldName string, nameCase FieldNameCase) string {
	// For normal databases like mysql, pgsql, sqlite,
	// field/table names of that are in normal case.
	var newFieldName = fieldName
	if isAllUpper(fieldName) {
		// For special databases like dm, oracle,
		// field/table names of that are in upper case.
		newFieldName = strings.ToLower(fieldName)
	}
	switch nameCase {
	case FieldNameCaseCamel:
		return gstr.CaseCamel(newFieldName)
	case FieldNameCaseCamelLower:
		return gstr.CaseCamelLower(newFieldName)
	case FieldNameCaseSnake:
		return gstr.CaseSnake(newFieldName)
	default:
		return ""
	}
}

// isAllUpper checks and returns whether given `fieldName` all letters are upper case.
func isAllUpper(fieldName string) bool {
	for _, b := range fieldName {
		if b >= 'a' && b <= 'z' {
			return false
		}
	}
	return true
}

type FieldNameCase string

const (
	FieldNameCaseCamel      FieldNameCase = "CaseCamel"
	FieldNameCaseCamelLower FieldNameCase = "CaseCamelLower"
	FieldNameCaseSnake      FieldNameCase = "CaseSnake"
)

func generateStructFieldDefinition(
	ctx context.Context, field *gdb.TableField, in genInput,
) (attrLines []string, appendImport string) {
	var (
		err              error
		localTypeName    gdb.LocalType
		localTypeNameStr string
		jsonTag          = gstr.CaseConvert(field.Name, gstr.CaseTypeMatch(in.JsonCase))
	)

	if in.TypeMapping != nil && len(in.TypeMapping) > 0 {
		var (
			tryTypeName string
		)
		tryTypeMatch, _ := gregex.MatchString(`(.+?)\((.+)\)`, field.Type)
		if len(tryTypeMatch) == 3 {
			tryTypeName = gstr.Trim(tryTypeMatch[1])
		} else {
			tryTypeName = gstr.Split(field.Type, " ")[0]
		}
		if tryTypeName != "" {
			if typeMapping, ok := in.TypeMapping[strings.ToLower(tryTypeName)]; ok {
				localTypeNameStr = typeMapping.Type
				appendImport = typeMapping.Import
			}
		}
	}

	if localTypeNameStr == "" {
		localTypeName, err = in.DB.CheckLocalTypeForField(ctx, field.Type, nil)
		if err != nil {
			panic(err)
		}
		localTypeNameStr = string(localTypeName)
		switch localTypeName {
		case gdb.LocalTypeDate, gdb.LocalTypeTime, gdb.LocalTypeDatetime:
			if in.StdTime {
				localTypeNameStr = "time.Time"
			} else {
				localTypeNameStr = "*gtime.Time"
			}

		case gdb.LocalTypeInt64Bytes:
			localTypeNameStr = "int64"

		case gdb.LocalTypeUint64Bytes:
			localTypeNameStr = "uint64"

		// Special type handle.
		case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
			if in.GJsonSupport {
				localTypeNameStr = "*gjson.Json"
			} else {
				localTypeNameStr = "string"
			}
		}
	}

	var (
		tagKey         = "`"
		descriptionTag = gstr.Replace(formatComment(field.Comment), `"`, `\"`)
	)
	removeFieldPrefixArray := gstr.SplitAndTrim(in.RemoveFieldPrefix, ",")
	newFiledName := field.Name
	for _, v := range removeFieldPrefixArray {
		newFiledName = gstr.TrimLeftStr(newFiledName, v, 1)
	}

	if in.FieldMapping != nil && len(in.FieldMapping) > 0 {
		if typeMapping, ok := in.FieldMapping[fmt.Sprintf("%s.%s", in.TableName, newFiledName)]; ok {
			localTypeNameStr = typeMapping.Type
			appendImport = typeMapping.Import
		}
	}

	attrLines = []string{
		"    #" + formatFieldName(newFiledName, FieldNameCaseCamel),
		" #" + localTypeNameStr,
	}
	attrLines = append(attrLines, fmt.Sprintf(` #%sjson:"%s"`, tagKey, jsonTag))
	// orm tag
	// if !in.IsDo {
	// 	// entity
	// 	attrLines = append(attrLines, fmt.Sprintf(` #orm:"%s"`, field.Name))
	// }
	attrLines = append(attrLines, fmt.Sprintf(` #description:"%s"%s`, descriptionTag, tagKey))
	attrLines = append(attrLines, fmt.Sprintf(` #// %s`, formatComment(field.Comment)))

	for k, v := range attrLines {
		// if in.NoJsonTag {
		// 	v, _ = gregex.ReplaceString(`json:".+"`, ``, v)
		// }
		// if !in.DescriptionTag {
		// 	v, _ = gregex.ReplaceString(`description:".*"`, ``, v)
		// }
		// if in.NoModelComment {
		// 	v, _ = gregex.ReplaceString(`//.+`, ``, v)
		// }
		attrLines[k] = v
	}
	return attrLines, appendImport
}

// formatComment formats the comment string to fit the golang code without any lines.
func formatComment(comment string) string {
	comment = gstr.ReplaceByArray(comment, g.SliceStr{
		"\n", " ",
		"\r", " ",
	})
	comment = gstr.Replace(comment, `\n`, " ")
	comment = gstr.Trim(comment)
	return comment
}

// GoFmt formats the source file and adds or removes import statements as necessary.
func GoFmt(path string) {
	replaceFunc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			glog.Infof(`error format "%s" go files: %v`, path, err)
			return content
		}
		return string(res)
	}

	var err error
	if gfile.IsFile(path) {
		// File format.
		if gfile.ExtName(path) != "go" {
			return
		}
		err = gfile.ReplaceFileFunc(replaceFunc, path)
	} else {
		// Folder format.
		err = gfile.ReplaceDirFunc(replaceFunc, path, "*.go", true)
	}
	if err != nil {
		glog.Infof(`error format "%s" go files: %v`, path, err)
	}
}
