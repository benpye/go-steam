package steam

import "github.com/benpye/go-steam/protocol"

// This event is emitted for every CMsgClientUserNotifications message and likewise only used for
// trade offers. Unlike the the above it is also emitted when the count of a type that was tracked
// before by this Notifications instance reaches zero.
type NotificationEvent struct {
	Type  NotificationType
	Count uint
}

func (e *NotificationEvent) GetJobID() protocol.JobID {
	return 0
}
