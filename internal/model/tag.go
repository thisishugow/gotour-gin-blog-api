package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint32 `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}
