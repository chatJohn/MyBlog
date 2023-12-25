package models

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%+v", err)))
	}
}
func InitTemplate(templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}
func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//访问首页的时候，由于存在多个模板的嵌套，所以需要解析所有的嵌套文件
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		pagination := templateDir + "layout/pagination.html"
		post := templateDir + "layout/post-list.html"
		viewHtml := templateDir + viewName

		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})

		files, err := t.ParseFiles(viewHtml, home, header, footer, personal, pagination, post)
		if err != nil {
			log.Println("Template ParseFiles Err: ", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = files
		tbs = append(tbs, tb)
	}
	return tbs, nil
}
