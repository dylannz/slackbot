package handlers

import "../api/git"

//HandleMessage is the handler for a Real Time Slack message Struct
func HandleMessage(s string) string {
	return git.GetPullRequestData()

	// return "Hey"
}
