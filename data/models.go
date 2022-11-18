package data

import "github.com/jmoiron/sqlx"

type Models struct {
	Users UserModel
	Posts PostModel
	Tags  TagModel
}

func NewModels(db *sqlx.DB) Models {
	return Models{
		Users: UserModel{DB: db},
		Posts: PostModel{DB: db},
		Tags:  TagModel{DB: db},
	}
}
