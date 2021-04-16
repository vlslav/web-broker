package storage

type (

	Storage struct {
		repo RepoStorage
	}

	RepoStorage interface{
		New()
	}
	FileRepo Storage
)

// func New() RepoStorage {
// 	var 
// 	return &Storage{}
// }