package main

import (
	"bufio"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const (
	ENV_USERNAME = "username"
	ENV_TOKEN    = "token"

	TEST_MODE = false
)

var (
	username = ""
	token    = ""
	repos    []string
)

func main() {
	readRepos()
	readEnv()

	//读取所有的repos
	//request := gorequest.New()
	//_, body, errs := request.Get("https://api.github.com/user/repos").
	//	SetBasicAuth(username, token).
	//	End()
	//if errs != nil {
	//	log.Fatalf("request.Get errro:%v", errs)
	//}
	//
	//var resp = &Resp{}
	//if err := json.Unmarshal([]byte(body), resp); err != nil {
	//	log.Fatalf("json.Unmarshal error:%v", err)
	//}

	//readmes := make([]string, 0)
	//for _, v := range *resp {
	//	if needRepo(v.Name) {
	//		readmes = append(readmes, downloadReadme(v.Name))
	//
	//		if TEST_MODE {
	//			break
	//		}
	//	}
	//}

	readmes := make([]string, 0)
	for _, v := range repos {
		readmes = append(readmes, downloadReadme(v))
	}

	//生成all_readme
	f, err := os.Create("files/add_readme.md")
	if err != nil {
		log.Fatalf("os.Create error:%v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("# all_readme\n")
	for _, v := range readmes {
		w.WriteString("\n#" + v + "\n")
	}
	w.Flush()
}

func readEnv() {
	username = os.Getenv(ENV_USERNAME)
	if username == "" {
		log.Fatalf("env username error")
	}

	token = os.Getenv(ENV_TOKEN)
	if token == "" {
		log.Fatalf("env token error")
	}
}

func readRepos() {
	data, err := ioutil.ReadFile("data/repos.txt")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	ss := strings.Split(string(data), ",")
	repos = ss
}

func needRepo(name string) bool {
	for _, v := range repos {
		if v == name {
			return true
		}
	}

	return false
}

func downloadReadme(repo string) string {
	url := fmt.Sprintf("https://raw.githubusercontent.com/liguoqinjim/%s/master/README.md", repo)

	request := gorequest.New()

	_, body, errs := request.Get(url).
		SetBasicAuth(username, token).
		End()
	if errs != nil {
		log.Fatalf("request.Get erros:%v", errs)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("files/%s_README.md", repo), []byte(body), 0644); err != nil {
		log.Fatalf("ioutil.WriteFile error:%v", err)
	}

	return body
}

type Resp []struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
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
	} `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int         `json:"forks"`
	OpenIssues       int         `json:"open_issues"`
	Watchers         int         `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
}
