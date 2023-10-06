package model

type UserModel struct {
	Id       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Email    string `json:"email" gorm:"column:email" binding:"required"`
	Password string `json:"password" gorm:"column:password" binding:"required"`
	Name     string `json:"name" gorm:"column:name" binding:"required"`
	Avatar   string `json:"avatar" gorm:"column:avatar" binding:"required"`
	Role     uint8  `json:"role" gorm:"column:role" binding:"required"`
	Re       uint8  `json:"re" gorm:"column:re" binding:"required"`
}

// TableName ... 物理表名
func (u *UserModel) TableName() string {
	return "users"
}

// Create doc
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetUserByEmailAndPassword(email string, password string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Table("users").Where("email = ? AND password = ?", email, password).First(u)
	return u, d.Error
}
