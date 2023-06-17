package notifications

import (
	"context"

	"github.com/owncast/owncast/db"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// AddNotification saves a new user notification destination.
func AddNotification(channel, destination string) error {
	return data.GetDatastore().GetQueries().AddNotification(context.Background(), db.AddNotificationParams{
		Channel:     channel,
		Destination: destination,
	})
}

// RemoveNotificationForChannel removes a notification destination.
func RemoveNotificationForChannel(channel, destination string) error {
	log.Debugln("Removing notification for channel", channel)
	return data.GetDatastore().GetQueries().RemoveNotificationDestinationForChannel(context.Background(), db.RemoveNotificationDestinationForChannelParams{
		Channel:     channel,
		Destination: destination,
	})
}

// GetNotificationDestinationsForChannel will return a collection of
// destinations to notify for a given channel.
func GetNotificationDestinationsForChannel(channel string) ([]string, error) {
	result, err := data.GetDatastore().GetQueries().GetNotificationDestinationsForChannel(context.Background(), channel)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query notification destinations for channel "+channel)
	}

	return result, nil
}
