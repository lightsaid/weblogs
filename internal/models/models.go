package models

// User 用户
type User struct {
	ID       int     `db:"id" json:"id"`
	Email    string  `db:"email" json:"email"`
	Username string  `db:"username" json:"username"`
	Password string  `db:"password" json:"password"`
	Avatar   *string `db:"avatar" json:"avatar"`
	IfAdmin  int     `db:"if_admin" json:"if_admin"`
	Active   int     `db:"active" json:"active"`
	// sqlte3 会将 datetime 转字符串
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

// Post 文章/帖子
type Post struct {
	ID        int     `db:"id" json:"id"`
	UserId    int     `db:"user_id" json:"user_id"`
	Author    string  `db:"author" json:"author"`
	Title     string  `db:"title" json:"title"`
	Content   string  `db:"content" json:"content"`
	Thumb     *string `db:"thumb" json:"thumb"`
	Readings  int     `db:"readings" json:"readings"`
	Comments  int     `db:"comments" json:"comments"`
	Likes     int     `db:"likes" json:"likes"`
	Active    int     `db:"active" json:"active"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}

// Category 分类
type Category struct {
	ID       int     `db:"id" json:"id"`
	UserID   int     `db:"user_id" json:"user_id"`
	ParentID *int    `db:"parent_id" json:"parent_id"`
	IfParent int     `db:"if_parent" json:"if_parent"`
	Name     string  `db:"name" json:"name"`
	Thumb    *string `db:"thumb" json:"thumb"`
}

// Attribute 属性,包含(tag|mark)
type Attribute struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Kind   string `db:"kind" json:"kind"`
	Name   string `db:"name" json:"name"`
}

// Comment 评论
type Comment struct {
	ID        int    `db:"id" json:"id"`
	PostID    int    `db:"post_id" json:"post_id"`
	Email     string `db:"email" json:"email"`
	Content   string `db:"content" json:"content"`
	ParentID  *int   `db:"parent_id" json:"parent_id"`
	Likes     int    `db:"likes" json:"likes"`
	Replynum  int    `db:"replynum" json:"replynum"`
	Active    int    `db:"active" json:"active"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

// Like 点赞文章/评论
type Like struct {
	ID        int    `db:"id" json:"id"`
	IPAddress string `db:"ip_address" json:"ip_address"`
	Kind      int    `db:"kind" json:"kind"`
	UserID    int    `db:"user_id" json:"user_id"`
	PCID      int    `db:"pc_id" json:"pc_id"` // 文章ID或者评论ID
}
