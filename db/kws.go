package db

//对应dKws表
type Kws struct {
	Kw string `json:"kw"`
}

func NewKws(kw string) *Kws {
	return &Kws{
		Kw: kw,
	}
}
func (e *Kws) Dels(did int) bool {
	return true
}
