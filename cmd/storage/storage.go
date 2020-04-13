package storage

type Storage interface {
	AddServer(groupID, ctid int64) (bool)
	GetServerList(groupID int64) ([]int64, bool)
	NextGroupID() (int64)
	RemoveGroupID(groupID int64)
}
