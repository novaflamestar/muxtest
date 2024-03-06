package modles

type Id struct {
	Id string `json:"id"`
}

type Ids struct {
	Count int  `json:"count"`
	Ids   []Id `json:"ids"`
}

func NewIds() Ids {
	return Ids{
		Count: 0,
		Ids:   make([]Id, 0),
	}
}
