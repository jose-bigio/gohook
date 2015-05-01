package gohook

import (
	"encoding/json"
)

// SenderType represents the structure of the sender field in push/ping events.
type SenderType struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// UserType represents the structure of a user as used in commits.
type UserType struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// PusherType represents the structure of a user as used in PushEvents.
type PusherType struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// OrgType represents the structure of an organization as used in ping/push events.
type OrgType struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

// CommitType represents the structure of a commit as used in push events.
type CommitType struct {
	ID        string   `json:"id"`
	Distinct  bool     `json:"distinct"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
	Author    UserType `json:"author"`
	Committer UserType `json:"committer"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

// RepoType represents the structure of a repository as used in push events.
type RepoType struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	FullName         string     `json:"full_name"`
	Owner            PusherType `json:"owner"`
	Private          bool       `json:"private"`
	HTMLURL          string     `json:"html_url"`
	Description      string     `json:"description"`
	Fork             bool       `json:"fork"`
	URL              string     `json:"url"`
	ForksURL         string     `json:"forks_url"`
	KeysURL          string     `json:"keys_url"`
	CollaboratorsURL string     `json:"collaborators_url"`
	TeamsURL         string     `json:"teams_url"`
	HooksURL         string     `json:"hooks_url"`
	IssueEventsURL   string     `json:"issue_events_url"`
	EventsURL        string     `json:"events_url"`
	AssigneesURL     string     `json:"assignees_url"`
	BranchesURL      string     `json:"branches_url"`
	TagsURL          string     `json:"tags_url"`
	BlobsURL         string     `json:"blobs_url"`
	GitTagsURL       string     `json:"git_tags_url"`
	GitRefsURL       string     `json:"git_refs_url"`
	TreesURL         string     `json:"trees_url"`
	StatusesURL      string     `json:"statuses_url"`
	LanguagesURL     string     `json:"languages_url"`
	StargazersURL    string     `json:"stargazers_url"`
	ContributorsURL  string     `json:"contributors_url"`
	SubscribersURL   string     `json:"subscribers_url"`
	SubscriptionURL  string     `json:"subscription_url"`
	CommitsURL       string     `json:"commits_url"`
	GitCommitsURL    string     `json:"git_commits_url"`
	CommentsURL      string     `json:"comments_url"`
	IssueCommentURL  string     `json:"issue_comment_url"`
	ContentsURL      string     `json:"contents_url"`
	CompareURL       string     `json:"compare_url"`
	MergesURL        string     `json:"mergers_url"`
	ArchiveURL       string     `json:"archive_url"`
	DownloadsURL     string     `json:"downloads_url"`
	IssuesURL        string     `json:"issues_url"`
	PullsURL         string     `json:"pulls_url"`
	MilestonesURL    string     `json:"milestones_url"`
	NotificationsURL string     `json:"notifications_url"`
	LabelsURL        string     `json:"labels_url"`
	ReleasesURL      string     `json:"releases_url"`
	CreatedAt        int        `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
	PushedAt         int        `json:"pushed_at"`
	GitURL           string     `json:"git_url"`
	SSHURL           string     `json:"ssh_url"`
	CloneURL         string     `json:"clone_url"`
	SvnURL           string     `json:"svn_url"`
	Homepage         string     `json:"homepage"`
	Size             int        `json:"size"`
	StargazersCount  int        `json:"stargazers_count"`
	WatchersCount    int        `json:"watchers_count"`
	Language         string     `json:"language"`
	HasIssues        bool       `json:"has_issues"`
	HasDownloads     bool       `json:"has_downloads"`
	HasWiki          bool       `json:"has_wiki"`
	HasPages         bool       `json:"has_pages"`
	ForksCount       int        `json:"forks_count"`
	MirrorURL        string     `json:"mirror_url"`
	OpenIssuesCount  int        `json:"open_issues_count"`
	Forks            int        `json:"forks"`
	OpenIssues       int        `json:"open_issues"`
	Watchers         int        `json:"watchers"`
	DefaultBranch    string     `json:"default_branch"`
	Stargazers       int        `json:"stargazers"`
	MasterBranch     string     `json:"master_branch"`
	Organization     string     `json:"organization"`
}

// PushEvent represents the basic, top-level structure of a push event.
type PushEvent struct {
	Ref          string       `json:"ref"`
	Before       string       `json:"before"`
	After        string       `json:"after"`
	Created      bool         `json:"created"`
	Deleted      bool         `json:"deleted"`
	Forced       bool         `json:"forced"`
	BaseRef      string       `json:"base_ref"`
	Compare      string       `json:"compare"`
	Commits      []CommitType `json:"commits"`
	HeadCommit   CommitType   `json:"head_commit"`
	Repository   RepoType     `json:"repository"`
	Pusher       PusherType   `json:"pusher"`
	Organization OrgType      `json:"organization"`
	Sender       SenderType   `json:"sender"`
}

// PingEvent represents the basic, top-level structure of a ping event.
type PingEvent struct {
	Zen          string          `json:"zen"`
	HookID       int             `json:"hook_id"`
	Hook         json.RawMessage `json:"hook"`
	Organization OrgType         `json:"organization"`
	Sender       SenderType      `json:"sender"`
}

// EventType is an alias for string that provides type safety for the event types.
type EventType string

// These constants define types for the implemented types of packets.
const (
	PingEventType = EventType("ping")
	PushEventType = EventType("push")
)

// EventAndType holds an event and its type, to be used later in a type assertion on the event.
type EventAndType struct {
	Event interface{}
	Type  EventType
}
