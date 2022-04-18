package pr

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/cian911/gh-gomerge/pkg/api"
	"github.com/cian911/gh-gomerge/pkg/ui/components/table"
	"github.com/cian911/gh-gomerge/pkg/ui/constants"
)

type PullRequest struct {
	Data api.PullRequest
}

type sectionPullRequestsFetchedMsg struct {
	SectionId int
	Prs       []PullRequest
}

func (pr PullRequest) renderReviewStatus() string {
	reviewCellStyle := lipgloss.NewStyle()
	if pr.Data.ReviewDecision == "APPROVED" {
		return reviewCellStyle.Foreground(lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"}).Render("")
	}

	if pr.Data.ReviewDecision == "CHANGES_REQUESTED" {
		return reviewCellStyle.Foreground(lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"}).Render("")
	}

	return reviewCellStyle.Render(constants.WaitingGlyph)
}

func (pr PullRequest) renderState() string {
	mergeCellStyle := lipgloss.NewStyle()
	switch pr.Data.State {
	case "OPEN":
		return mergeCellStyle.Foreground(openPR).Render("")
	case "CLOSED":
		return mergeCellStyle.Foreground(closedPR).Render("")
	case "MERGED":
		return mergeCellStyle.Foreground(mergedPR).Render("")
	default:
		return mergeCellStyle.Foreground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#3E4057"}).Render("-")
	}
}

func (pr PullRequest) GetStatusChecksRollup() string {
	if pr.Data.Mergeable == "CONFLICTING" {
		return "FAILURE"
	}

	accStatus := "SUCCESS"
	/* mostRecentCommit := pr.Data.Commits.Nodes[0].Commit */
	/* for _, statusCheck := range mostRecentCommit.StatusCheckRollup.Contexts.Nodes { */
	/*   var conclusion string */
	/*   if statusCheck.Typename == "CheckRun" { */
	/*     conclusion = string(statusCheck.CheckRun.Conclusion) */
	/*     status := string(statusCheck.CheckRun.Status) */
	/*     if isStatusWaiting(status) { */
	/*       accStatus = "PENDING" */
	/*     } */
	/*   } else if statusCheck.Typename == "StatusContext" { */
	/*     conclusion = string(statusCheck.StatusContext.State) */
	/*   } */
	/*  */
	/*   if isConclusionAFailure(conclusion) { */
	/*     accStatus = "FAILURE" */
	/*     break */
	/*   } */
	/* } */

	return accStatus
}

func (pr PullRequest) renderCiStatus() string {

	accStatus := pr.GetStatusChecksRollup()
	ciCellStyle := lipgloss.NewStyle()
	if accStatus == "SUCCESS" {
		return ciCellStyle.Render(constants.SuccessGlyph)
	}

	if accStatus == "PENDING" {
		return ciCellStyle.Render(constants.WaitingGlyph)
	}

	return ciCellStyle.Render(constants.FailureGlyph)
}

/* func (pr PullRequest) renderLines() string { */
/* deletions := 0 */
/* if pr.Data.Deletions > 0 { */
/*   deletions = pr.Data.Deletions */
/* } */
/*  */
/* return lipgloss.NewStyle().Render( */
/*   fmt.Sprintf("%d / -%d", pr.Data.Additions, deletions), */
/* ) */
/* } */

func (pr PullRequest) renderTitle() string {
	return lipgloss.NewStyle().
		Inline(true).
		Foreground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#3E4057"}).
		Render(
			fmt.Sprintf("#%d %s",
				pr.Data.Number,
				titleText.Copy().Render(pr.Data.Title),
			),
		)
}

func (pr PullRequest) renderAuthor() string {
	return lipgloss.NewStyle().Render(pr.Data.Author.Login)
}

func (pr PullRequest) renderRepoName() string {
	// repoName := utils.TruncateString(pr.Data.HeadRepository.Name, 18)
	repoName := pr.Data.HeadRepository.Name
	return lipgloss.NewStyle().
		Render(repoName)
}

func (pr PullRequest) renderUpdateAt() string {
	return lipgloss.NewStyle().
		Render(pr.Data.UpdatedAt.String())
}

func (pr PullRequest) RenderState() string {
	switch pr.Data.State {
	case "OPEN":
		return " Open"
	case "CLOSED":
		return "﫧Closed"
	case "MERGED":
		return " Merged"
	default:
		return ""
	}
}

func (pr PullRequest) ToTableRow() table.Row {
	return table.Row{
		pr.renderUpdateAt(),
		pr.renderRepoName(),
		pr.renderTitle(),
		pr.renderAuthor(),
		pr.renderReviewStatus(),
		pr.renderState(),
		pr.renderCiStatus(),
		// pr.renderLines(),
	}
}

func isConclusionAFailure(conclusion string) bool {
	return conclusion == "FAILURE" || conclusion == "TIMED_OUT" || conclusion == "STARTUP_FAILURE"
}

func isStatusWaiting(status string) bool {
	return status == "PENDING" ||
		status == "QUEUED" ||
		status == "IN_PROGRESS" ||
		status == "WAITING"
}
