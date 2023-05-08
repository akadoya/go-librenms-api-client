package librenms

// AlertRule -
type AlertRule struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Rule        string       `json:"rule"`
	Severity    string       `json:"severity"`
	Extra       string       `json:"extra"`
	Disabled    int          `json:"disabled"`
	Query       string       `json:"query"`
	Builder     string       `json:"builder"`
	Proc        string       `json:"proc"`
	InvertMap   int          `json:"invert_map"`
	Devices     []int        `json:"device_ids"`
//	Devices     []Device     `json:"devices"`
}

/* Device -
type Device struct {
	ID       int    `json:"rule_id"`
	Name     string `json:"name"`
	...
}
*/