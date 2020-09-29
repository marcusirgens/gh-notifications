package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"time"
)

// GetRecentNotifications gets GitHub notifications for the current user using the provided
// GitHub client. The request age is rounded to the nearest 10 minutes. This helps us
// cache the requests (as they will be pretty similar every time).
func GetRecentNotifications(ghc *github.Client) ([]*github.Notification, error) {
	ctx := context.Background()

	maxAge := time.Hour * 24 * 7 // a week seems reasonable
	// Round the age to nearest 10 minutes. This helps us cache the request :)
	roundedAge := roundTime(maxAge, time.Minute * 10)

	opts := &github.NotificationListOptions{
		All:           false, // don't get read notifications
		Participating: true, // only get stuff where you're "directly participating or mentioned".
		Since:         roundedAge,
	}

	ns := make([]*github.Notification, 0)

	for  {
		pgn, resp, err := ghc.Activity.ListNotifications(ctx, opts);

		if err != nil {
			return ns, fmt.Errorf("error when fetching notifications: %w", err)
		}

		ns = append(ns, pgn...)

		if resp.NextPage == 0 {
			// no more pages
			return ns, nil
		}

		opts.Page = resp.NextPage
	}
}

func roundTime(since time.Duration, round time.Duration) time.Time {
	base := time.Now().Add(-since)

	u := base.Unix()
	s := int64(round / time.Second)
	r := u % s
	nu := u - r

	return time.Unix(nu, 0)
}
