package api

type emailStruct struct {
	ID        int          `json:"id"`
	Created   string       `json:"created"`
	Subject   string       `json:"subject"`
	MessageID string       `json:"message_id"`
	ParentID  *int         `json:"parent_id"`
	ThreadID  *int         `json:"thread_it"`
	List      listsStruct  `json:"list"`
	Sender    *userStruct  `json:"sender"`
	Patchset  *patchStruct `json:"patchset"`
}

type emailPagerStruct struct {
	Next           *int           `json:"next,string"`
	Results        []*emailStruct `json:"results"`
	Total          int            `json:"total"`
	ResultsPerPage int            `json:"results_per_page"`
}

type patchStruct struct {
	ID        int     `json:"id"`
	Created   string  `json:"created"`
	Updated   string  `json:"updated"`
	Subject   string  `json:"subject"`
	Prefix    string  `json:"prefix"`
	Version   int     `json:"version"`
	Status    string  `json:"status"`
	Submitter string  `json:"submitter"`
	ReplyTo   *string `json:"reply_to"`
	MessageID string  `json:"message_id"`
}

type listPermissions struct {
	NonSubscriber []string `json:"nonsubscriber"`
	Subscriber    []string `json:"subscriber"`
	Account       []string `json:"account"`
}

type listsStruct struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Owner       userStruct      `json:"owner"`
	Created     string          `json:"created"`
	Updated     string          `json:"updated"`
	Description string          `json:"description"`
	Permissions listPermissions `json:"permissions"`
}

type listsPagerStruct struct {
	Next           *int           `json:"next,string"`
	Results        []*listsStruct `json:"results"`
	Total          int            `json:"total"`
	ResultsPerPage int            `json:"results_per_page"`
}
