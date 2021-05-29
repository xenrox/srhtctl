package api

type permissionStruct struct {
	Anonymous []string `json:"anonymous"`
	Submitter []string `json:"submitter"`
	User      []string `json:"user"`
}

type trackerStruct struct {
	ID                 int              `json:"id"`
	Owner              userStruct       `json:"owner"`
	Created            string           `json:"created"`
	Updated            string           `json:"updated"`
	Name               string           `json:"name"`
	Description        string           `json:"description"`
	DefaultPermissions permissionStruct `json:"default_permissions"`
}

type trackerPagerStruct struct {
	Next           *int             `json:"next,string"`
	Results        []*trackerStruct `json:"results"`
	Total          int              `json:"total"`
	ResultsPerPage int              `json:"results_per_page"`
}

type ticketStruct struct {
	ID          int              `json:"id"`
	Ref         string           `json:"ref"`
	Tracker     trackerStruct    `json:"tracker"`
	Title       string           `json:"title"`
	Created     string           `json:"created"`
	Updated     string           `json:"updated"`
	Submitter   userStruct       `json:"submitter"`
	Description string           `json:"description"`
	Status      string           `json:"status"`
	Resolution  string           `json:"resolution"`
	Permissions permissionStruct `json:"permissions"`
	Labels      []string         `json:"labels"`
	Assignees   []string         `json:"assignees"`
}

type ticketPagerStruct struct {
	Next           *int            `json:"next,string"`
	Results        []*ticketStruct `json:"results"`
	Total          int             `json:"total"`
	ResultsPerPage int             `json:"results_per_page"`
}
