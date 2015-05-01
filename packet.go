package gohook

import "encoding/json"

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
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	FullName         string          `json:"full_name"`
	Owner            SenderType      `json:"owner"`
	Private          bool            `json:"private"`
	HTMLURL          string          `json:"html_url"`
	Description      string          `json:"description"`
	Fork             bool            `json:"fork"`
	URL              string          `json:"url"`
	ForksURL         string          `json:"forks_url"`
	KeysURL          string          `json:"keys_url"`
	CollaboratorsURL string          `json:"collaborators_url"`
	TeamsURL         string          `json:"teams_url"`
	HooksURL         string          `json:"hooks_url"`
	IssueEventsURL   string          `json:"issue_events_url"`
	EventsURL        string          `json:"events_url"`
	AssigneesURL     string          `json:"assignees_url"`
	BranchesURL      string          `json:"branches_url"`
	TagsURL          string          `json:"tags_url"`
	BlobsURL         string          `json:"blobs_url"`
	GitTagsURL       string          `json:"git_tags_url"`
	GitRefsURL       string          `json:"git_refs_url"`
	TreesURL         string          `json:"trees_url"`
	StatusesURL      string          `json:"statuses_url"`
	LanguagesURL     string          `json:"languages_url"`
	StargazersURL    string          `json:"stargazers_url"`
	ContributorsURL  string          `json:"contributors_url"`
	SubscribersURL   string          `json:"subscribers_url"`
	SubscriptionURL  string          `json:"subscription_url"`
	CommitsURL       string          `json:"commits_url"`
	GitCommitsURL    string          `json:"git_commits_url"`
	CommentsURL      string          `json:"comments_url"`
	IssueCommentURL  string          `json:"issue_comment_url"`
	ContentsURL      string          `json:"contents_url"`
	CompareURL       string          `json:"compare_url"`
	MergesURL        string          `json:"mergers_url"`
	ArchiveURL       string          `json:"archive_url"`
	DownloadsURL     string          `json:"downloads_url"`
	IssuesURL        string          `json:"issues_url"`
	PullsURL         string          `json:"pulls_url"`
	MilestonesURL    string          `json:"milestones_url"`
	NotificationsURL string          `json:"notifications_url"`
	LabelsURL        string          `json:"labels_url"`
	ReleasesURL      string          `json:"releases_url"`
	CreatedAt        json.RawMessage `json:"created_at"`
	UpdatedAt        json.RawMessage `json:"updated_at"`
	PushedAt         json.RawMessage `json:"pushed_at"`
	GitURL           string          `json:"git_url"`
	SSHURL           string          `json:"ssh_url"`
	CloneURL         string          `json:"clone_url"`
	SvnURL           string          `json:"svn_url"`
	Homepage         string          `json:"homepage"`
	Size             int             `json:"size"`
	StargazersCount  int             `json:"stargazers_count"`
	WatchersCount    int             `json:"watchers_count"`
	Language         string          `json:"language"`
	HasIssues        bool            `json:"has_issues"`
	HasDownloads     bool            `json:"has_downloads"`
	HasWiki          bool            `json:"has_wiki"`
	HasPages         bool            `json:"has_pages"`
	ForksCount       int             `json:"forks_count"`
	MirrorURL        string          `json:"mirror_url"`
	OpenIssuesCount  int             `json:"open_issues_count"`
	Forks            int             `json:"forks"`
	OpenIssues       int             `json:"open_issues"`
	Watchers         int             `json:"watchers"`
	DefaultBranch    string          `json:"default_branch"`
	Stargazers       int             `json:"stargazers"`
	MasterBranch     string          `json:"master_branch"`
	Organization     string          `json:"organization"`
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

type ConfigType struct {
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
}

type HookType struct {
	ID        int        `json:"id"`
	URL       string     `json:"url"`
	TestURL   string     `json:"test_url"`
	PingURL   string     `json:"ping_url"`
	Name      string     `json:"name"`
	Events    []string   `json:"events"`
	Active    bool       `json:"active"`
	Config    ConfigType `json:"config"`
	UpdatedAt string     `json:"updated_at"`
	CreatedAt string     `json:"created_at"`
}

// PingEvent represents the basic, top-level structure of a ping event.
type PingEvent struct {
	Zen          string     `json:"zen"`
	HookID       int        `json:"hook_id"`
	Hook         HookType   `json:"hook"`
	Organization OrgType    `json:"organization"`
	Sender       SenderType `json:"sender"`
}

type CommentType struct {
	URL       string     `json:"url"`
	HTMLURL   string     `json:"html_url"`
	ID        int        `json:"id"`
	User      SenderType `json:"user"`
	Position  string     `json:"position"`
	Line      string     `json:"line"`
	Path      string     `json:"path"`
	CommitID  string     `json:"commit_id"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	Body      string     `json:"body"`
}

type CommitCommentEvent struct {
	Comment    CommentType `json:"comment"`
	Repository RepoType    `json:"repository"`
	Sender     SenderType  `json:"sender"`
}

type IssueCommentType struct {
	ID        int        `json:"id"`
	URL       string     `json:"url"`
	HTMLURL   string     `json:"html_url"`
	Body      string     `json:"body"`
	User      SenderType `json:"user"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"created_at"`
}

type LabelType struct {
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type MilestoneType struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int        `json:"id"`
	Number       int        `json:"number"`
	State        string     `json:"state"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      SenderType `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	ClosedAt     string     `json:"closed_at"`
	DueOn        string     `json:"due_on"`
}

type ShortPullRequestType struct {
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}

type IssueType struct {
	ID          int                  `json:"id"`
	URL         string               `json:"url"`
	HTMLURL     string               `json:"html_url"`
	Number      int                  `json:"number"`
	State       string               `json:"state"`
	Title       string               `json:"title"`
	Body        string               `json:"body"`
	User        SenderType           `json:"user"`
	Labels      []LabelType          `json:"labels"`
	Assignee    SenderType           `json:"assignee"`
	Milestone   MilestoneType        `json:"milestone"`
	Comments    int                  `json:"comments"`
	PullRequest ShortPullRequestType `json:"pull_request"`
	ClosedAt    string               `json:"closed_at"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at"`
}

type IssueCommentEvent struct {
	Action     string           `json:"action"`
	Issue      IssueType        `json:"issue"`
	Comment    IssueCommentType `json:"comment"`
	Repository RepoType         `json:"repository"`
	Sender     SenderType       `json:"sender"`
}

type IssuesEvent struct {
	Action     string     `json:"action"`
	Issue      IssueType  `json:"issue"`
	Assignee   SenderType `json:"assignee"`
	Label      LabelType  `json:"label"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type CreateEvent struct {
	RefType      string     `json:"ref_type"`
	Ref          string     `json:"ref"`
	MasterBranch string     `json:"master_branch"`
	Description  string     `json:"description"`
	PusherType   string     `json:"pusher_type"`
	Repository   RepoType   `json:"repository"`
	Sender       SenderType `json:"sender"`
}

type DeleteEvent struct {
	RefType    string     `json:"ref_type"`
	Ref        string     `json:"ref"`
	PusherType string     `json:"pusher_type"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type RepositoryEvent struct {
	Action       string     `json:"action"`
	Repository   RepoType   `json:"repository"`
	Organization OrgType    `json:"organization"`
	Sender       SenderType `json:"sender"`
}

type DeploymentType struct {
	URL           string            `json:"url"`
	ID            int               `json:"id"`
	SHA           string            `json:"sha"`
	Ref           string            `json:"ref"`
	Task          string            `json:"task"`
	Payload       map[string]string `json:"payload"`
	Environment   string            `json:"environment"`
	Description   string            `json:"description"`
	Creator       SenderType        `json:"creator"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
	StatusesURL   string            `json:"statuses_url"`
	RepositoryURL string            `json:"repository_url"`
}

type DeploymentEvent struct {
	Deployment  DeploymentType    `json:"deployment"`
	ID          int               `json:"id"`
	SHA         string            `json:"sha"`
	Ref         string            `json:"ref"`
	Task        string            `json:"task"`
	Name        string            `json:"name"`
	Environment string            `json:"environment"`
	Payload     map[string]string `json:"payload"`
	Description string            `json:"description"`
	Repository  RepoType          `json:"repository"`
	Sender      SenderType        `json:"sender"`
}

type DeploymentStatusType struct {
	URL           string     `json:"url"`
	ID            int        `json:"id"`
	State         string     `json:"state"`
	Creator       SenderType `json:"creator"`
	Description   string     `json:"description"`
	TargetURL     string     `json:"target_url"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	DeploymentURL string     `json:"deployment_url"`
	RepositoryURL string     `json:"repository_url"`
}

type DeploymentStatusEvent struct {
	Deployment       DeploymentType       `json:"deployment"`
	DeploymentStatus DeploymentStatusType `json:"deployment_status"`
	ID               int                  `json:"id"`
	State            string               `json:"state"`
	TargetURL        string               `json:"target_url"`
	Description      string               `json:"description"`
	Repository       RepoType             `json:"repository"`
	Sender           SenderType           `json:"sender"`
}

type ForkEvent struct {
	Forkee     RepoType   `json:"forkee"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type PageType struct {
	PageName string `json:"page_name"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Action   string `json:"action"`
	SHA      string `json:"sha"`
	HTMLURL  string `json:"html_url"`
}

type GollumEvent struct {
	Pages      []PageType `json:"pages"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type MemberEvent struct {
	Action     string     `json:"action"`
	Member     SenderType `json:"member"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type TeamType struct {
	Name            string `json:"name"`
	ID              int    `json:"id"`
	Slug            string `json:"slug"`
	Permission      string `json:"permission"`
	URL             string `json:"url"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
}

type MembershipEvent struct {
	Action       string     `json:"action"`
	Scope        string     `json:"scope"`
	Member       SenderType `json:"member"`
	Team         TeamType   `json:"team"`
	Sender       SenderType `json:"sender"`
	Organization OrgType    `json:"organization"`
}

type ErrorType struct {
	Message string `json:"message"`
}

type BuildType struct {
	URL       string     `json:"url"`
	Status    string     `json:"status"`
	Error     ErrorType  `json:"error"`
	Pusher    SenderType `json:"pusher"`
	Commit    string     `json:"commit"`
	Duration  int        `json:"duration"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type PageBuildEvent struct {
	ID         int        `json:"id"`
	Build      BuildType  `json:"build"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type PublicEvent struct {
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type BaseHeadType struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	User  SenderType `json:"user"`
	Repo  RepoType   `json:"repo"`
}

type Link struct {
	HREF string `json:"href"`
}

type PullRequestType struct {
	ID                int             `json:"id"`
	URL               string          `json:"url"`
	HTMLURL           string          `json:"html_url"`
	DiffURL           string          `json:"diff_url"`
	PatchURL          string          `json:"patch_url"`
	IssueURL          string          `json:"issue_url"`
	CommitsURL        string          `json:"commits_url"`
	ReviewCommentsURL string          `json:"review_comments_url"`
	ReviewCommentURL  string          `json:"review_comment_url"`
	CommentsURL       string          `json:"comments_url"`
	StatusesURL       string          `json:"statuses_url"`
	Number            int             `json:"number"`
	State             string          `json:"state"`
	Title             string          `json:"title"`
	Body              string          `json:"body"`
	CreatedAt         string          `json:"created_at"`
	UpdatedAt         string          `json:"updated_at"`
	ClosedAt          string          `json:"closed_at"`
	MergedAt          string          `json:"merged_at"`
	Head              BaseHeadType    `json:"head"`
	Repo              RepoType        `json:"repo"`
	Base              BaseHeadType    `json:"Base"`
	Links             map[string]Link `json:"_links"`
	User              SenderType      `json:"User"`
	MergeCommitSHA    string          `json:"merge_commit_sha"`
	Merged            bool            `json:"merged"`
	Mergeable         bool            `json:"mergeable"`
	MergedBy          SenderType      `json:"merged_by"`
	Comments          int             `json:"comments"`
	Commits           int             `json:"commits"`
	Additions         int             `json:"additions"`
	Deletions         int             `json:"deletions"`
	ChangedFiles      int             `json:"changed_files"`
}

type PullRequestEvent struct {
	Action      string          `json:"action"`
	Number      int             `json:"number"`
	PullRequest PullRequestType `json:"pull_request"`
	Repository  RepoType        `json:"repository"`
	Sender      SenderType      `json:"sender"`
}

type PullRequestReviewCommentType struct {
	URL              string          `json:"url"`
	ID               int             `json:"id"`
	DiffHunk         string          `json:"diff_hunk"`
	Path             string          `json:"path"`
	Position         int             `json:"position"`
	OriginalPosition int             `json:"original_position"`
	CommitID         string          `json:"commit_id"`
	OriginalCommitID string          `json:"original_commit_id"`
	User             SenderType      `json:"user"`
	Body             string          `json:"body"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
	HTMLURL          string          `json:"html_url"`
	PullRequestURL   string          `json:"pull_request_url"`
	Links            map[string]Link `json:"_links"`
}

type PullRequestReviewCommentEvent struct {
	Action      string                       `json:"action"`
	Comment     PullRequestReviewCommentType `json:"comment"`
	PullRequest PullRequestType              `json:"pull_request"`
	Repository  RepoType                     `json:"repository"`
	Sender      SenderType                   `json:"sender"`
}

type AssetType struct {
	URL                string     `json:"url"`
	BrowserDownloadURL string     `json:"browser_download_url"`
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	Label              string     `json:"label"`
	State              string     `json:"state"`
	ContentType        string     `json:"content_type"`
	Size               int        `json:"size"`
	DownloadCount      int        `json:"download_count"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	Uploader           SenderType `json:"uploader"`
}

type ReleaseType struct {
	URL             string      `json:"url"`
	HTMLURL         string      `json:"html_url"`
	AssetsURL       string      `json:"assets_url"`
	UploadURL       string      `json:"upload_url"`
	TarballURL      string      `json:"tarball_url"`
	ZipballURL      string      `json:"zipball_url"`
	ID              int         `json:"id"`
	TagName         string      `json:"tag_name"`
	TargetCommitish string      `json:"target_commitish"`
	Name            string      `json:"name"`
	Body            string      `json:"body"`
	Draft           bool        `json:"draft"`
	Prerelease      bool        `json:"prerelease"`
	CreatedAt       string      `json:"created_at"`
	PublishedAt     string      `json:"published_at"`
	Author          SenderType  `json:"author"`
	Assets          []AssetType `json:"assets"`
}

type ReleaseEvent struct {
	Action     string      `json:"action"`
	Release    ReleaseType `json:"release"`
	Repository RepoType    `json:"repository"`
	Sender     SenderType  `json:"sender"`
}

type ShortCommitType struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type BranchType struct {
	Name   string          `json:"master"`
	Commit ShortCommitType `json:"commit"`
}

type StatusEvent struct {
	ID          int          `json:"id"`
	SHA         string       `json:"sha"`
	Name        string       `json:"name"`
	TargetURL   string       `json:"target_url"`
	Context     string       `json:"context"`
	Description string       `json:"description"`
	State       string       `json:"state"`
	Commit      CommitType   `json:"commit"`
	Branches    []BranchType `json:"branches"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Repository  RepoType     `json:"repository"`
	Sender      SenderType   `json:"sender"`
}

type TeamAddEvent struct {
	Team         TeamType   `json:"team"`
	Repository   RepoType   `json:"repository"`
	Organization OrgType    `json:"organization"`
	Sender       SenderType `json:"sender"`
}

type WatchEvent struct {
	Action       string     `json:"action"`
	Repository   RepoType   `json:"repository"`
	Sender       SenderType `json:"sender"`
	Organization OrgType    `json:"organization"`
}

// EventType is an alias for string that provides type safety for the event types.
type EventType string

// These constants define types for the implemented types of packets.
const (
	PingEventType                     = EventType("ping")
	PushEventType                     = EventType("push")
	CommitCommentEventType            = EventType("commit_comment")
	IssueCommentEventType             = EventType("issue_comment")
	IssuesEventType                   = EventType("issues")
	CreateEventType                   = EventType("create")
	DeleteEventType                   = EventType("delete")
	RepositoryEventType               = EventType("repository")
	DeploymentEventType               = EventType("deployment")
	DeploymentStatusEventType         = EventType("deployment_status")
	ForkEventType                     = EventType("fork")
	GollumEventType                   = EventType("gollum")
	MemberEventType                   = EventType("member")
	MembershipEventType               = EventType("membership")
	PageBuildEventType                = EventType("page_build")
	PublicEventType                   = EventType("public")
	PullRequestEventType              = EventType("pull_request")
	PullRequestReviewCommentEventType = EventType("pull_request_review_comment")
	ReleaseEventType                  = EventType("release")
	StatusEventType                   = EventType("status")
	TeamAddEventType                  = EventType("team_add")
	WatchEventType                    = EventType("watch")
)

// EventAndType holds an event and its type, to be used later in a type assertion on the event.
type EventAndType struct {
	Event interface{}
	Type  EventType
}
