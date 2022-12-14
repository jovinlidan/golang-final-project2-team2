package user_domain

import (
	"time"
)

type User struct {
	Id        int64      `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Age       int8       `json:"age"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//func (u *User) GetUserIdParam(c *gin.Context) (int64, error_utils.MessageErr) {
//	idParam := c.Param("userId")
//	userId, err := strconv.Atoi(idParam)
//	if err != nil {
//		return int64(0), error_utils.NewBadRequest("invalid user id params")
//	}
//
//	return int64(userId), nil
//}
