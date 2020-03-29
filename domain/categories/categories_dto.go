package categories

type Category struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"Description"`
	CreatedAt    string `json:"created_at"`
}

func (c Category) Validate() {

}
