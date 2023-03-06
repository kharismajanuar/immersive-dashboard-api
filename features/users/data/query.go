package data

import (
	// "errors"
	"errors"
	"immersive-dashboard/features/users"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) RegisterData(newUser users.Core) error {
	data := CoreToUser(newUser)
	tx := uq.db.Create(&data)
	if tx.Error != nil {
		log.Println("error query", tx.Error)
		return tx.Error
	}
	return nil
}

func (uq *userQuery) LoginData(email string) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("email = ?", email).First(&tmp)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}

func (uq *userQuery) GetUser(pageNum int, keyword string) ([]users.Core, error) {
	tmp := []User{}
	pageSize := 2
	log.Println("key:", keyword)
	offset := (pageNum - 1) * pageSize
	// tx := uq.db.Raw("SELECT * FROM users LIMIT ?, OFFSET ? WHERE name = ?", pageSize, offset, keyword).Find(&tmp)
	tx := uq.db.Where("name LIKE ?", "%"+keyword+"%").Offset(offset).Limit(pageSize).Find(&tmp)
	if tx.RowsAffected < 1 {
		return nil, errors.New("users not found, no data displayed")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	listUser := ListUserToCore(tmp)
	return listUser, nil
}
