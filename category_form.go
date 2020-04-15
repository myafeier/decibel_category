package category

type PostForm struct {
	Name      string `json:"name"`
	Icon      string `json:"icon,omitempty"`
	ListOrder int    `json:"list_order"`
	Pid       int64  `json:"pid"`
}

type DeleteForm struct {
	Id int64 `json:"id"`
}
