package gohook

import (
	"time"
	"json"
)

// SenderType represents the structure of the sender field in push/ping events.
type SenderType struct {
	Login             json.Number `json:"login"`
	ID                int         `json:"id"`
	AvatarURL         json.Number `json:"avatar_url"`
	GravatarID        json.Number `json:"gravatar_id"`
	URL               json.Number `json:"url"`
	HTMLURL           json.Number `json:"html_url"`
	FollowersURL      json.Number `json:"followers_url"`
	FollowingURL      json.Number `json:"following_url"`
	GistsURL          json.Number `json:"gists_url"`
	StarredURL        json.Number `json:"starred_url"`
	SubscriptionsURL  json.Number `json:"subscriptions_url"`
	OrganizationsURL  json.Number `json:"organizations_url"`
	ReposURL          json.Number `json:"repos_url"`
	EventsURL         json.Number `json:"events_url"`
	ReceivedEventsURL json.Number `json:"received_events_url"`
	Type              json.Number `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
}

// UserType represents the structure of a user as used in commits.
type UserType struct {
	Name     json.Number `json:"name"`
	Email    json.Number `json:"email"`
	Username json.Number `json:"username"`
}

// PusherType represents the structure of a user as used in PushEvents.
type PusherType struct {
	Name  json.Number `json:"name"`
	Email json.Number `json:"email"`
}

// OrgType represents the structure of an organization as used in ping/push events.
type OrgType struct {
	Login            json.Number `json:"login"`
	ID               int         `json:"id"`
	URL              json.Number `json:"url"`
	ReposURL         json.Number `json:"repos_url"`
	EventsURL        json.Number `json:"events_url"`
	MembersURL       json.Number `json:"members_url"`
	PublicMembersURL json.Number `json:"public_members_url"`
	AvatarURL        json.Number `json:"avatar_url"`
	Description      json.Number `json:"description"`
}

// CommitType represents the structure of a commit as used in push events.
type CommitType struct {
	ID        json.Number   `json:"id"`
	Distinct  bool          `json:"distinct"`
	Message   json.Number   `json:"message"`
	Timestamp json.Number   `json:"timestamp"`
	URL       json.Number   `json:"url"`
	Author    UserType      `json:"author"`
	Committer UserType      `json:"committer"`
	Added     []string      `json:"added"`
	Removed   []string      `json:"removed"`
	Modified  []string      `json:"modified"`
}

// RepoType represents the structure of a repository as used in push events.
type RepoType struct {
	ID               int                  `json:"id"`
	Name             json.Number          `json:"name"`
	FullName         json.Number          `json:"full_name"`
	Owner            SenderType           `json:"owner"`
	Private          bool                 `json:"private"`
	HTMLURL          json.Number          `json:"html_url"`
	Description      json.Number          `json:"description"`
	Fork             bool                 `json:"fork"`
	URL              json.Number          `json:"url"`
	ForksURL         json.Number          `json:"forks_url"`
	KeysURL          json.Number          `json:"keys_url"`
	CollaboratorsURL json.Number          `json:"collaborators_url"`
	TeamsURL         json.Number          `json:"teams_url"`
	HooksURL         json.Number          `json:"hooks_url"`
	IssueEventsURL   json.Number          `json:"issue_events_url"`
	EventsURL        json.Number          `json:"events_url"`
	AssigneesURL     json.Number          `json:"assignees_url"`
	BranchesURL      json.Number          `json:"branches_url"`
	TagsURL          json.Number          `json:"tags_url"`
	BlobsURL         json.Number          `json:"blobs_url"`
	GitTagsURL       json.Number          `json:"git_tags_url"`
	GitRefsURL       json.Number          `json:"git_refs_url"`
	TreesURL         json.Number          `json:"trees_url"`
	StatusesURL      json.Number          `json:"statuses_url"`
	LanguagesURL     json.Number          `json:"languages_url"`
	StargazersURL    json.Number          `json:"stargazers_url"`
	ContributorsURL  json.Number          `json:"contributors_url"`
	SubscribersURL   json.Number          `json:"subscribers_url"`
	SubscriptionURL  json.Number          `json:"subscription_url"`
	CommitsURL       json.Number          `json:"commits_url"`
	GitCommitsURL    json.Number          `json:"git_commits_url"`
	CommentsURL      json.Number          `json:"comments_url"`
	IssueCommentURL  json.Number          `json:"issue_comment_url"`
	ContentsURL      json.Number          `json:"contents_url"`
	CompareURL       json.Number          `json:"compare_url"`
	MergesURL        json.Number          `json:"mergers_url"`
	ArchiveURL       json.Number          `json:"archive_url"`
	DownloadsURL     json.Number          `json:"downloads_url"`
	IssuesURL        json.Number          `json:"issues_url"`
	PullsURL         json.Number          `json:"pulls_url"`
	MilestonesURL    json.Number          `json:"milestones_url"`
	NotificationsURL json.Number          `json:"notifications_url"`
	LabelsURL        json.Number          `json:"labels_url"`
	ReleasesURL      json.Number          `json:"releases_url"`
	CreatedAt        time.Time 		      `json:"created_at"`
	UpdatedAt        time.Time 		      `json:"updated_at"`
	PushedAt         time.Time            `json:"pushed_at"`
	GitURL           json.Number          `json:"git_url"`
	SSHURL           json.Number          `json:"ssh_url"`
	CloneURL         json.Number          `json:"clone_url"`
	SvnURL           json.Number          `json:"svn_url"`
	Homepage         json.Number          `json:"homepage"`
	Size             int                  `json:"size"`
	StargazersCount  int                  `json:"stargazers_count"`
	WatchersCount    int                  `json:"watchers_count"`
	Language         json.Number          `json:"language"`
	HasIssues        bool                 `json:"has_issues"`
	HasDownloads     bool                 `json:"has_downloads"`
	HasWiki          bool                 `json:"has_wiki"`
	HasPages         bool                 `json:"has_pages"`
	ForksCount       int                  `json:"forks_count"`
	MirrorURL        json.Number          `json:"mirror_url"`
	OpenIssuesCount  int                  `json:"open_issues_count"`
	Forks            int                  `json:"forks"`
	OpenIssues       int                  `json:"open_issues"`
	Watchers         int                  `json:"watchers"`
	DefaultBranch    json.Number          `json:"default_branch"`
	Stargazers       int                  `json:"stargazers"`
	MasterBranch     json.Number          `json:"master_branch"`
	Organization     json.Number          `json:"organization"`
}

// RepoType represents the structure of a repository as used in push events.
type RepoPushType struct {
	ID               int                  `json:"id"`
	Name             json.Number          `json:"name"`
	FullName         json.Number          `json:"full_name"`
	Owner            SenderType           `json:"owner"`
	Private          bool                 `json:"private"`
	HTMLURL          json.Number          `json:"html_url"`
	Description      json.Number          `json:"description"`
	Fork             bool                 `json:"fork"`
	URL              json.Number          `json:"url"`
	ForksURL         json.Number          `json:"forks_url"`
	KeysURL          json.Number          `json:"keys_url"`
	CollaboratorsURL json.Number          `json:"collaborators_url"`
	TeamsURL         json.Number          `json:"teams_url"`
	HooksURL         json.Number          `json:"hooks_url"`
	IssueEventsURL   json.Number          `json:"issue_events_url"`
	EventsURL        json.Number          `json:"events_url"`
	AssigneesURL     json.Number          `json:"assignees_url"`
	BranchesURL      json.Number          `json:"branches_url"`
	TagsURL          json.Number          `json:"tags_url"`
	BlobsURL         json.Number          `json:"blobs_url"`
	GitTagsURL       json.Number          `json:"git_tags_url"`
	GitRefsURL       json.Number          `json:"git_refs_url"`
	TreesURL         json.Number          `json:"trees_url"`
	StatusesURL      json.Number          `json:"statuses_url"`
	LanguagesURL     json.Number          `json:"languages_url"`
	StargazersURL    json.Number          `json:"stargazers_url"`
	ContributorsURL  json.Number          `json:"contributors_url"`
	SubscribersURL   json.Number          `json:"subscribers_url"`
	SubscriptionURL  json.Number          `json:"subscription_url"`
	CommitsURL       json.Number          `json:"commits_url"`
	GitCommitsURL    json.Number          `json:"git_commits_url"`
	CommentsURL      json.Number          `json:"comments_url"`
	IssueCommentURL  json.Number          `json:"issue_comment_url"`
	ContentsURL      json.Number          `json:"contents_url"`
	CompareURL       json.Number          `json:"compare_url"`
	MergesURL        json.Number          `json:"mergers_url"`
	ArchiveURL       json.Number          `json:"archive_url"`
	DownloadsURL     json.Number          `json:"downloads_url"`
	IssuesURL        json.Number          `json:"issues_url"`
	PullsURL         json.Number          `json:"pulls_url"`
	MilestonesURL    json.Number          `json:"milestones_url"`
	NotificationsURL json.Number          `json:"notifications_url"`
	LabelsURL        json.Number          `json:"labels_url"`
	ReleasesURL      json.Number          `json:"releases_url"`
	CreatedAt        int64 		          `json:"created_at"`
	UpdatedAt        time.Time 		      `json:"updated_at"`
	PushedAt         int64                `json:"pushed_at"`
	GitURL           json.Number          `json:"git_url"`
	SSHURL           json.Number          `json:"ssh_url"`
	CloneURL         json.Number          `json:"clone_url"`
	SvnURL           json.Number          `json:"svn_url"`
	Homepage         json.Number          `json:"homepage"`
	Size             int                  `json:"size"`
	StargazersCount  int                  `json:"stargazers_count"`
	WatchersCount    int                  `json:"watchers_count"`
	Language         json.Number          `json:"language"`
	HasIssues        bool                 `json:"has_issues"`
	HasDownloads     bool                 `json:"has_downloads"`
	HasWiki          bool                 `json:"has_wiki"`
	HasPages         bool                 `json:"has_pages"`
	ForksCount       int                  `json:"forks_count"`
	MirrorURL        json.Number          `json:"mirror_url"`
	OpenIssuesCount  int                  `json:"open_issues_count"`
	Forks            int                  `json:"forks"`
	OpenIssues       int                  `json:"open_issues"`
	Watchers         int                  `json:"watchers"`
	DefaultBranch    json.Number          `json:"default_branch"`
	Stargazers       int                  `json:"stargazers"`
	MasterBranch     json.Number          `json:"master_branch"`
	Organization     json.Number          `json:"organization"`
}

// PushEvent represents the basic, top-level structure of a push event.
type PushEvent struct {
	Ref          json.Number       `json:"ref"`
	Before       json.Number       `json:"before"`
	After        json.Number       `json:"after"`
	Created      bool              `json:"created"`
	Deleted      bool              `json:"deleted"`
	Forced       bool              `json:"forced"`
	BaseRef      json.Number       `json:"base_ref"`
	Compare      json.Number       `json:"compare"`
	Commits      []CommitType `json:"commits"`
	HeadCommit   CommitType   `json:"head_commit"`
	Repository   RepoPushType `json:"repository"`
	Pusher       PusherType   `json:"pusher"`
	Organization OrgType      `json:"organization"`
	Sender       SenderType   `json:"sender"`
}

type ConfigType struct {
	URL         json.Number `json:"url"`
	ContentType json.Number `json:"content_type"`
}

type HookType struct {
	ID        int             `json:"id"`
	URL       json.Number     `json:"url"`
	TestURL   json.Number     `json:"test_url"`
	PingURL   json.Number     `json:"ping_url"`
	Name      json.Number     `json:"name"`
	Events    []string        `json:"events"`
	Active    bool            `json:"active"`
	Config    ConfigType      `json:"config"`
	UpdatedAt time.Time       `json:"updated_at"`
	CreatedAt time.Time       `json:"created_at"`
}

// PingEvent represents the basic, top-level structure of a ping event.
type PingEvent struct {
	Zen          json.Number     `json:"zen"`
	HookID       int             `json:"hook_id"`
	Hook         HookType        `json:"hook"`
	Organization OrgType         `json:"organization"`
	Sender       SenderType      `json:"sender"`
}

type CommentType struct {
	URL       json.Number     `json:"url"`
	HTMLURL   json.Number     `json:"html_url"`
	ID        int             `json:"id"`
	User      SenderType      `json:"user"`
	Position  json.Number     `json:"position"`
	Line      json.Number     `json:"line"`
	Path      json.Number     `json:"path"`
	CommitID  json.Number     `json:"commit_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Body      json.Number     `json:"body"`
}

type CommitCommentEvent struct {
	Comment    CommentType `json:"comment"`
	Repository RepoType    `json:"repository"`
	Sender     SenderType  `json:"sender"`
}

type IssueCommentType struct {
	ID        int             `json:"id"`
	URL       json.Number     `json:"url"`
	HTMLURL   json.Number     `json:"html_url"`
	Body      json.Number     `json:"body"`
	User      SenderType      `json:"user"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"created_at"`
}

type LabelType struct {
	URL   json.Number `json:"url"`
	Name  json.Number `json:"name"`
	Color json.Number `json:"color"`
}

type MilestoneType struct {
	URL          json.Number     `json:"url"`
	HTMLURL      json.Number     `json:"html_url"`
	LabelsURL    json.Number     `json:"labels_url"`
	ID           int             `json:"id"`
	Number       int             `json:"number"`
	State        json.Number     `json:"state"`
	Title        json.Number     `json:"title"`
	Description  json.Number     `json:"description"`
	Creator      SenderType      `json:"creator"`
	OpenIssues   int             `json:"open_issues"`
	ClosedIssues int             `json:"closed_issues"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	ClosedAt     *time.Time      `json:"closed_at"`
	DueOn        json.Number     `json:"due_on"`
}

type ShortPullRequestType struct {
	URL      json.Number `json:"url"`
	HTMLURL  json.Number `json:"html_url"`
	DiffURL  json.Number `json:"diff_url"`
	PatchURL json.Number `json:"patch_url"`
}

type IssueType struct {
	ID          int                       `json:"id"`
	URL         json.Number               `json:"url"`
	HTMLURL     json.Number               `json:"html_url"`
	Number      int                       `json:"number"`
	State       json.Number               `json:"state"`
	Title       json.Number               `json:"title"`
	Body        json.Number               `json:"body"`
	User        SenderType                `json:"user"`
	Labels      []LabelType               `json:"labels"`
	Assignee    SenderType                `json:"assignee"`
	Milestone   MilestoneType             `json:"milestone"`
	Comments    int                       `json:"comments"`
	PullRequest ShortPullRequestType      `json:"pull_request"`
	ClosedAt    *time.Time                `json:"closed_at"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}

type IssueCommentEvent struct {
	Action     json.Number           `json:"action"`
	Issue      IssueType             `json:"issue"`
	Comment    IssueCommentType      `json:"comment"`
	Repository RepoType              `json:"repository"`
	Sender     SenderType            `json:"sender"`
}

type IssuesEvent struct {
	Action     json.Number     `json:"action"`
	Issue      IssueType       `json:"issue"`
	Assignee   SenderType      `json:"assignee"`
	Label      LabelType       `json:"label"`
	Repository RepoType        `json:"repository"`
	Sender     SenderType      `json:"sender"`
}

type CreateEvent struct {
	RefType      json.Number     `json:"ref_type"`
	Ref          json.Number     `json:"ref"`
	MasterBranch json.Number     `json:"master_branch"`
	Description  json.Number     `json:"description"`
	PusherType   json.Number     `json:"pusher_type"`
	Repository   RepoType        `json:"repository"`
	Sender       SenderType      `json:"sender"`
}

type DeleteEvent struct {
	RefType    json.Number     `json:"ref_type"`
	Ref        json.Number     `json:"ref"`
	PusherType json.Number     `json:"pusher_type"`
	Repository RepoType        `json:"repository"`
	Sender     SenderType      `json:"sender"`
}

type RepositoryEvent struct {
	Action       json.Number     `json:"action"`
	Repository   RepoType        `json:"repository"`
	Organization OrgType         `json:"organization"`
	Sender       SenderType      `json:"sender"`
}

type DeploymentType struct {
	URL           json.Number            `json:"url"`
	ID            int                    `json:"id"`
	SHA           json.Number            `json:"sha"`
	Ref           json.Number            `json:"ref"`
	Task          json.Number            `json:"task"`
	Payload       map[string]string      `json:"payload"`
	Environment   json.Number            `json:"environment"`
	Description   json.Number            `json:"description"`
	Creator       SenderType             `json:"creator"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	StatusesURL   json.Number            `json:"statuses_url"`
	RepositoryURL json.Number            `json:"repository_url"`
}

type DeploymentEvent struct {
	Deployment  DeploymentType         `json:"deployment"`
	ID          int                    `json:"id"`
	SHA         json.Number            `json:"sha"`
	Ref         json.Number            `json:"ref"`
	Task        json.Number            `json:"task"`
	Name        json.Number            `json:"name"`
	Environment json.Number            `json:"environment"`
	Payload     map[string]string      `json:"payload"`
	Description json.Number            `json:"description"`
	Repository  RepoType               `json:"repository"`
	Sender      SenderType             `json:"sender"`
}

type DeploymentStatusType struct {
	URL           json.Number     `json:"url"`
	ID            int             `json:"id"`
	State         json.Number     `json:"state"`
	Creator       SenderType      `json:"creator"`
	Description   json.Number     `json:"description"`
	TargetURL     json.Number     `json:"target_url"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeploymentURL json.Number     `json:"deployment_url"`
	RepositoryURL json.Number     `json:"repository_url"`
}

type DeploymentStatusEvent struct {
	Deployment       DeploymentType            `json:"deployment"`
	DeploymentStatus DeploymentStatusType      `json:"deployment_status"`
	ID               int                       `json:"id"`
	State            json.Number               `json:"state"`
	TargetURL        json.Number               `json:"target_url"`
	Description      json.Number               `json:"description"`
	Repository       RepoType                  `json:"repository"`
	Sender           SenderType                `json:"sender"`
}

type ForkEvent struct {
	Forkee     RepoType   `json:"forkee"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type PageType struct {
	PageName json.Number `json:"page_name"`
	Title    json.Number `json:"title"`
	Summary  json.Number `json:"summary"`
	Action   json.Number `json:"action"`
	SHA      json.Number `json:"sha"`
	HTMLURL  json.Number `json:"html_url"`
}

type GollumEvent struct {
	Pages      []PageType `json:"pages"`
	Repository RepoType   `json:"repository"`
	Sender     SenderType `json:"sender"`
}

type MemberEvent struct {
	Action     json.Number     `json:"action"`
	Member     SenderType      `json:"member"`
	Repository RepoType        `json:"repository"`
	Sender     SenderType      `json:"sender"`
}

type TeamType struct {
	Name            json.Number `json:"name"`
	ID              int         `json:"id"`
	Slug            json.Number `json:"slug"`
	Permission      json.Number `json:"permission"`
	URL             json.Number `json:"url"`
	MembersURL      json.Number `json:"members_url"`
	RepositoriesURL json.Number `json:"repositories_url"`
}

type MembershipEvent struct {
	Action       json.Number     `json:"action"`
	Scope        json.Number     `json:"scope"`
	Member       SenderType      `json:"member"`
	Team         TeamType        `json:"team"`
	Sender       SenderType      `json:"sender"`
	Organization OrgType         `json:"organization"`
}

type ErrorType struct {
	Message json.Number `json:"message"`
}

type BuildType struct {
	URL       json.Number     `json:"url"`
	Status    json.Number     `json:"status"`
	Error     ErrorType       `json:"error"`
	Pusher    SenderType      `json:"pusher"`
	Commit    json.Number     `json:"commit"`
	Duration  int             `json:"duration"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
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
	Label json.Number     `json:"label"`
	Ref   json.Number     `json:"ref"`
	SHA   json.Number     `json:"sha"`
	User  SenderType      `json:"user"`
	Repo  RepoType        `json:"repo"`
}

type Link struct {
	HREF json.Number `json:"href"`
}

type PullRequestType struct {
	ID                int                  `json:"id"`
	URL               json.Number          `json:"url"`
	HTMLURL           json.Number          `json:"html_url"`
	DiffURL           json.Number          `json:"diff_url"`
	PatchURL          json.Number          `json:"patch_url"`
	IssueURL          json.Number          `json:"issue_url"`
	CommitsURL        json.Number          `json:"commits_url"`
	ReviewCommentsURL json.Number          `json:"review_comments_url"`
	ReviewCommentURL  json.Number          `json:"review_comment_url"`
	CommentsURL       json.Number          `json:"comments_url"`
	StatusesURL       json.Number          `json:"statuses_url"`
	Number            int                  `json:"number"`
	State             json.Number          `json:"state"`
	Title             json.Number          `json:"title"`
	Body              json.Number          `json:"body"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
	ClosedAt          *time.Time           `json:"closed_at"`
	MergedAt          *time.Time           `json:"merged_at"`
	Head              BaseHeadType         `json:"head"`
	Repo              RepoType             `json:"repo"`
	Base              BaseHeadType         `json:"Base"`
	Links             map[string]Link      `json:"_links"`
	User              SenderType           `json:"User"`
	MergeCommitSHA    json.Number          `json:"merge_commit_sha"`
	Merged            bool                 `json:"merged"`
	Mergeable         bool                 `json:"mergeable"`
	MergedBy          SenderType           `json:"merged_by"`
	Comments          int                  `json:"comments"`
	Commits           int                  `json:"commits"`
	Additions         int                  `json:"additions"`
	Deletions         int                  `json:"deletions"`
	ChangedFiles      int                  `json:"changed_files"`
}

type PullRequestEvent struct {
	Action      json.Number          `json:"action"`
	Number      int                  `json:"number"`
	PullRequest PullRequestType      `json:"pull_request"`
	Repository  RepoType             `json:"repository"`
	Sender      SenderType           `json:"sender"`
}

type PullRequestReviewCommentType struct {
	URL              json.Number          `json:"url"`
	ID               int                  `json:"id"`
	DiffHunk         json.Number          `json:"diff_hunk"`
	Path             json.Number          `json:"path"`
	Position         int                  `json:"position"`
	OriginalPosition int                  `json:"original_position"`
	CommitID         json.Number          `json:"commit_id"`
	OriginalCommitID json.Number          `json:"original_commit_id"`
	User             SenderType           `json:"user"`
	Body             json.Number          `json:"body"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
	HTMLURL          json.Number          `json:"html_url"`
	PullRequestURL   json.Number          `json:"pull_request_url"`
	Links            map[string]Link      `json:"_links"`
}

type PullRequestReviewCommentEvent struct {
	Action      json.Number                  `json:"action"`
	Comment     PullRequestReviewCommentType `json:"comment"`
	PullRequest PullRequestType              `json:"pull_request"`
	Repository  RepoType                     `json:"repository"`
	Sender      SenderType                   `json:"sender"`
}

type AssetType struct {
	URL                json.Number     `json:"url"`
	BrowserDownloadURL json.Number     `json:"browser_download_url"`
	ID                 int             `json:"id"`
	Name               json.Number     `json:"name"`
	Label              json.Number     `json:"label"`
	State              json.Number     `json:"state"`
	ContentType        json.Number     `json:"content_type"`
	Size               int             `json:"size"`
	DownloadCount      int             `json:"download_count"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	Uploader           SenderType      `json:"uploader"`
}

type ReleaseType struct {
	URL             json.Number      `json:"url"`
	HTMLURL         json.Number      `json:"html_url"`
	AssetsURL       json.Number      `json:"assets_url"`
	UploadURL       json.Number      `json:"upload_url"`
	TarballURL      json.Number      `json:"tarball_url"`
	ZipballURL      json.Number      `json:"zipball_url"`
	ID              int              `json:"id"`
	TagName         json.Number      `json:"tag_name"`
	TargetCommitish json.Number      `json:"target_commitish"`
	Name            json.Number      `json:"name"`
	Body            json.Number      `json:"body"`
	Draft           bool             `json:"draft"`
	Prerelease      bool             `json:"prerelease"`
	CreatedAt       time.Time        `json:"created_at"`
	PublishedAt     time.Time        `json:"published_at"`
	Author          SenderType       `json:"author"`
	Assets          []AssetType      `json:"assets"`
}

type ReleaseEvent struct {
	Action     json.Number `json:"action"`
	Release    ReleaseType `json:"release"`
	Repository RepoType    `json:"repository"`
	Sender     SenderType  `json:"sender"`
}

type ShortCommitType struct {
	SHA json.Number `json:"sha"`
	URL json.Number `json:"url"`
}

type BranchType struct {
	Name   json.Number     `json:"master"`
	Commit ShortCommitType `json:"commit"`
}

type StatusEvent struct {
	ID          int          `json:"id"`
	SHA         json.Number  `json:"sha"`
	Name        json.Number  `json:"name"`
	TargetURL   json.Number  `json:"target_url"`
	Context     json.Number  `json:"context"`
	Description json.Number  `json:"description"`
	State       json.Number  `json:"state"`
	Commit      CommitType   `json:"commit"`
	Branches    []BranchType `json:"branches"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
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
	Action       json.Number     `json:"action"`
	Repository   RepoType        `json:"repository"`
	Sender       SenderType      `json:"sender"`
	Organization OrgType         `json:"organization"`
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
