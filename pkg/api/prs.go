package api

import (
	"log"

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
		Search struct {
			Nodes []struct {
				PullRequest PullRequest `graphql:"... on PullRequest"`
			}
		} `graphql:"search(type: ISSUE, first: $limit, query: $query)"`
	}
	vars := map[string]interface{}{
		"query": graphql.String("is:pr author:Cian911"),
		"limit": graphql.Int(10),
	}

	err := c.client.Query("SearchPullRequests", &queryResult, vars)
	if err != nil {
		log.Fatalf("Error fetching pull requests: %v", err)
		return nil
	}

	prs = make([]PullRequest, 0, len(queryResult.Search.Nodes))
	for _, node := range queryResult.Search.Nodes {
		prs = append(prs, node.PullRequest)
	}

	return prs
}

// Approve pull request for a given repository
func ApprovePullRequest() {}

// Merge pull request for a given repository
func MergePullRequest() {}
