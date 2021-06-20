package api

import (
	"fmt"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
)

// TicketStatus is the status of a ticket
var TicketStatus string

// TrackerName is the name of a tracker
var TrackerName string

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

type commentStruct struct {
	ID        int          `json:"id"`
	Created   string       `json:"created"`
	Submitter userStruct   `json:"submitter"`
	Text      string       `json:"text"`
	Ticket    ticketStruct `json:"ticket"`
}

type eventStruct struct {
	ID            int            `json:"id"`
	Created       string         `json:"created"`
	EventType     []string       `json:"event_type"`
	OldStatus     *string        `json:"old_status"`
	OldResoltion  *string        `json:"old_resolution"`
	NewStatus     *string        `json:"new_status"`
	NewResolution *string        `json:"new_resolution"`
	User          *userStruct    `json:"user"`
	Ticket        *ticketStruct  `json:"ticket"`
	Comment       *commentStruct `json:"comment"`
	Label         *[]string      `json:"label"`
	ByUser        *userStruct    `json:"by_user"`
	FromTicket    *ticketStruct  `json:"from_ticket"`
}

// PrintTickets prints out tickets of a user
func PrintTickets(args []string) error {
	var tickets ticketPagerStruct

	if TrackerName != "" {
		err := getTickets(&tickets, TrackerName)
		if err != nil {
			return err
		}

		for _, ticket := range tickets.Results {
			fmt.Print(ticket.filterByStatus())
		}
		return nil
	}

	var trackers trackerPagerStruct
	err := getTrackers(&trackers)
	if err != nil {
		return err
	}

	for _, tracker := range trackers.Results {
		err = getTickets(&tickets, tracker.Name)
		if err != nil {
			return err
		}

		for i, ticket := range tickets.Results {
			if i == 0 {
				fmt.Printf("Tracker %s:\n\n", tracker.Name)
			}
			fmt.Print(ticket.filterByStatus())
		}
	}

	return nil
}

func getTickets(response *ticketPagerStruct, trackerName string) error {
	var url string
	if UserName != "" {
		url = fmt.Sprintf("%s/api/user/%s/trackers/%s/tickets",
			config.GetURL("todo"), helpers.TransformCanonical(UserName), trackerName)
	} else {
		url = fmt.Sprintf("%s/api/trackers/%s/tickets", config.GetURL("todo"), trackerName)
	}

	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	return nil
}

func (ticket ticketStruct) filterByStatus() string {
	if TicketStatus == "all" || TicketStatus == ticket.Status {
		return fmt.Sprintln(ticket)
	}
	return ""
}

func getTrackers(response *trackerPagerStruct) error {
	var url string
	if UserName != "" {
		url = fmt.Sprintf("%s/api/user/%s/trackers", config.GetURL("todo"),
			helpers.TransformCanonical(UserName))
	} else {
		url = fmt.Sprintf("%s/api/trackers", config.GetURL("todo"))
	}

	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	return nil
}

func (ticket ticketStruct) String() string {
	str := fmt.Sprintf("Ticket %d: %s\n", ticket.ID, ticket.Title)
	str += fmt.Sprintf("Submitter: %s\n", ticket.Submitter.Name)
	str += fmt.Sprintf("Status: %s with %s\n", ticket.Status, ticket.Resolution)
	// Trim newline in tickets submitted by email
	str += fmt.Sprintln(strings.TrimSpace(ticket.Description))

	return str
}
