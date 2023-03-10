package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]string
	Data      map[string]*interface{}
	CSRFToken string
	flash     string
	warning   string
}
