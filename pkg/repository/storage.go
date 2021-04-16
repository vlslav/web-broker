package repository

type (

	Storage struct {
		repo RepoStorage
	}

	RepoStorage interface{
		New()
	}

)

// func New() RepoStorage {
// 	var 
// 	return &Storage{}
// }