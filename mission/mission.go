package mission

// Mission .
type Mission struct {
	Name          string   `json:"mission_name"`
	ID            string   `json:"mission_id"`
	Manufacturers []string `json:"manufacturers"`
	PayloadIds    []string `json:"payload_ids"`
	WikipediaLink string   `json:"wikipedia"`
	WebsiteLink   string   `json:"website"`
	TwitterLink   string   `json:"twitter"`
	Description   string   `json:"description"`
}
