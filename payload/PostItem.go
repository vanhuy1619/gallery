package payload

type PostItem struct {
	Id        int            `json:"id"`
	Username  string         `json:"username"`
	Content   string         `json:"content"`
	Image     map[string]any `json:"image"`
	DateUp    string         `json:"date_up"`
	UpdateVer map[string]any `json:"update_ver"`
}
