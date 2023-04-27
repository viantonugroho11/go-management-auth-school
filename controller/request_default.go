package controller

// DefaultParameter ....
type DefaultParameter struct {
	ID      string    `param:"id"`
	Search  string `query:"search"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	Offset  int
	OrderBy string `query:"order_by"`
	Sort    string `query:"sort"`
}