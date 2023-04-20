package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname  string `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" binding:"require,email" gorm:"unique " `
	Password  string `json:"password" form:"password" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" form:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" form:"address" gorm:"type:text"`
	Subscribe bool   `json:"subscribe" gorm:"type: varchar(255)"`
	Role      string `json:"role" gorm:"type: varchar(255)"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
