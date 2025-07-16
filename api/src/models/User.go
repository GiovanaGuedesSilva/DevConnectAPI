package models

type User struct {
	ID        uint64 `json:"id,omitempty"` // omitempty means this field will be omitted when the value is zero
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
