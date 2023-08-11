package model

type User struct {
	ID       int     `validate:"required,min=1" json:"id,omitempty"`
	Username string  `validate:"required" json:"username,omitempty"`
	Age      int     `validate:"min=1,max=120" json:"age,omitempty"`
	Classes  []Class `json:"classes,omitempty"`
}

type Class struct {
	ID      int    `validate:"required,min=1"`
	Name    string `validate:"required"`
	Teacher string `validate:"required"`
}
