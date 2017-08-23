package datastore

type Release struct {
	Collection
}

func NewReleaseDatastore() Release {
	return Release{Collection{}}
}