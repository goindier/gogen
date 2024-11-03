package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main() {
	// sql := "update user set state = req.input, updated_at = 'curr', received_by = '$uid' where order_id = 1 AND state = 2"
	// statement, err := sqlparser.Parse(sql)
	// if err != nil {
	// 	fmt.Println("Error parsing SQL:", err)
	// 	return
	// }

	// updateStmt, ok := statement.(*sqlparser.Update)
	// if !ok {
	// 	fmt.Println("Not an UPDATE statement")
	// 	return
	// }

	// // 提取字段
	// for _, expr := range updateStmt.Exprs {
	// 	fmt.Printf("Column: %s, Value: %s\n", expr.Name.Name, expr.Expr)
	// }

	// // 提取 WHERE 条件
	// where := updateStmt.Where
	// if where != nil {
	// 	fmt.Printf("Where condition: %s\n", where)
	// }

	tmpl, err := template.New("test").ParseFiles("t.tpl")
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(nil)
	data := map[string]string{
		"Hello": "demo",
	}
	tmpl.ExecuteTemplate(buffer, "t.tpl", data)

	str := buffer.String()
	fmt.Println("str:", str)
}
