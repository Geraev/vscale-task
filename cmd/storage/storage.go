package storage

const (
	StatusAccepted = "accepted"
	StatusFailed   = "failed"
	StatusComplete = "complete"
	StatusDeleted  = "deleted"
)

type Storage interface {
	AddServer(groupID, ctid int64) bool
	RemoveServer(groupID, ctid int64) bool
	GetServerList(groupID int64) ([]int64, bool)
	NextGroupID() int64
	RemoveGroupID(groupID int64)
	GetGroupStatus(groupID int64) (string, bool)
	SetGroupStatus(groupID int64, status string) bool
}
