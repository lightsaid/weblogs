package data

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int     `db:"id" json:"id"`
	Email    string  `db:"email" json:"email"`
	UserName string  `db:"username" json:"username"`
	Password string  `db:"password" json:"password"`
	Avatar   *string `db:"avatar" json:"avatar"`
	Role     int     `db:"if_admin" json:"role"`
	Active   int     `db:"active" json:"active"`
	// sqlte3 会将 datetime 转字符串
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type UserModel struct {
	DB *sqlx.DB
}

// Insert 插入一个用户
func (m *UserModel) Insert(email, username, password, avatar string) error {
	query := `insert into users(email, username, password, avatar) values($1, $2, $3, $4)`

	result, err := m.DB.Exec(query, email, username, password, avatar)
	if err != nil {
		return fmt.Errorf("insert user error: %w", err)
	}
	if id, err := result.LastInsertId(); err != nil || id <= 0 {
		if err != nil {
			return fmt.Errorf("insert user with lastInsertId error: %w", err)
		}
		return errors.New("insert user lastInsertId <= 0")
	}
	return nil
}

// GetById 根据ID获取一个未删除的用户
func (m *UserModel) GetById(id int) (*User, error) {
	query := `select id, email, password, username, avatar, role, active from users where id=$1 and active != -1 limit 1;`

	var user User
	err := m.DB.Get(&user, query, id)
	if err != nil {
		return nil, fmt.Errorf("GetById id:%d error: %w", id, err)
	}
	return &user, nil
}

// GetByEmail 根据Email获取一个未删除的用户
func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `select id, email, password, username, avatar, role, active from users where email=$1 and active != -1 limit 1;`

	var user User
	err := m.DB.Get(&user, query, email)
	if err != nil {
		return nil, fmt.Errorf("GetByEmail email:%v error: %w", email, err)
	}
	return &user, nil
}

// Update 更新一个未删除用户信息
func (m *UserModel) Update(id int, user *User) error {
	findUser, err := m.GetById(id)
	if err != nil {
		return err
	}
	username := user.UserName
	avatar := user.Avatar
	update := time.Now()
	if username == "" {
		username = findUser.UserName
	}
	if avatar == nil {
		avatar = findUser.Avatar
	}
	query := `
		update users 
		set username = $2
			avatar = $3,
			active = $4,
			updated_at = $5
		where id = $1`

	result, err := m.DB.Exec(query, id, username, avatar, user.Active, update)
	if err != nil {
		return fmt.Errorf("update user error: %w", err)
	}
	if n, err := result.RowsAffected(); err != nil || n <= 0 {
		if err != nil {
			return fmt.Errorf("update user which RowsAffected error: %w", err)
		}
		return errors.New("update user RowsAffected <= 0")
	}
	return nil
}

// Delete 删除一个用户
func (m *UserModel) Delete(id int) error {
	query := `update users set active = -1 where id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}

	if n, err := result.RowsAffected(); err != nil || n <= 0 {
		if err != nil {
			return fmt.Errorf("delete user which RowsAffected error: %w", err)
		}
		return errors.New("user not found")
	}
	return nil
}
