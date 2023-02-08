package models

type TemplateData struct {
	StringMap map[string]string
	Data      map[string]interface{}
	Error     string
}
