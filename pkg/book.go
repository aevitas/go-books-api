package pkg

type Book struct {
	Id     uint64  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
