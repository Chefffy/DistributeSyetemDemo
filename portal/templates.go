package portal

import (
	"html/template"
)

var rootTemplate *template.Template

// ImportTemplates 加载页面
func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"../../portal/students.html",
		"../../portal/student.html")

	if err != nil {
		return err
	}

	return nil
}