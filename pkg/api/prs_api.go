package api

import (
	"log"
	"time"

	"github.com/cli/go-gh"
	gClient "github.com/cli/go-gh/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type GhClient struct {
	client gClient.GQLClient
}

type PullRequest struct {
	Number int
	Title  string
	Body   string
	Author struct {
		Login string
	}
	Mergeable      string
	State          string
	IsDraft        bool
	ReviewDecision string
	Additions      int
	Deletions      int
	HeadRepository struct {
		Name string
	}
	UpdatedAt time.Time
}

func New() *GhClient {
	var err error
	newClient, err := gh.GQLClient(nil)
	if err != nil {
		log.Fatalf("Error initalizing client: %v", err)
	}

	c := GhClient{
		client: newClient,
	}

	return &c
}

// Get a list of pull requests for a repository
func (c *GhClient) ListPullRequests() (prs []PullRequest) {
	var queryResult struct {
		Repository struct {
			PullRequests struct {
				Nodes []struct {
					PullRequest
				}
			} `graphql:"pullRequests(first:100, states:OPEN)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	vars := map[string]interface{}{
		"owner": graphql.String("Cian911"),
		"name":  graphql.String("switchboard"),
	}

	err := c.client.Query("SearchPullRequests", &queryResult, vars)
	if err != nil {
		log.Fatalf("Error fetching pull requests: %v", err)
		return nil
	}

	prs = make([]PullRequest, 0, len(queryResult.Repository.PullRequests.Nodes))
	for _, node := range queryResult.Repository.PullRequests.Nodes {
		prs = append(prs, node.PullRequest)
	}

	return prs
}

// Approve pull request for a given repository
func ApprovePullRequest() {}

// Merge pull request for a given repository
func MergePullRequest() {}
