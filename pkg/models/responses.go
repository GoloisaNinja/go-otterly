package models

type build interface {
	build(int, string, interface{})
}
type StandardResponse struct {
	Status  int
	Message string
	Data    interface{}
}

func (r *StandardResponse) Build(s int, m string, d interface{}) {
	r.Status = s
	r.Message = m
	r.Data = d
}
