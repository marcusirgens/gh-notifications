package main

import (
	"fmt"
	"github.com/google/go-github/v32/github"
	"github.com/logrusorgru/aurora/v3"
)

const timeFormat = "Mon 15:04"

func printNotifications(ns []*github.Notification) {
	if len(ns) == 0 {
		// Do absolutely nothing in this case. No point in spamming the user.
		return
	}

	fmt.Println("GitHub notifications:")

	for _, n := range ns {
		printNotification(n)
	}
}

func printNotification(n *github.Notification) {

	updateTime := fmt.Sprintf("%s", aurora.Cyan(n.UpdatedAt.Format(timeFormat)))
	repo := aurora.Blue(*n.Repository.FullName)
	subject := *n.Subject.Title

	message := fmt.Sprintf("%s in %s: %s", updateTime, repo, subject)

	fmt.Printf("\t%s\n", message)
}
