package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"strings"
	"time"
)

type User struct {
	ID                    uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username              string    `gorm:"size:255;not null;unique" json:"username"`
	Email                 string    `gorm:"size:100;not null;unique" json:"email"`
	Password              string    `gorm:"size:100;not null;" json:"password"`
	Department            string    `gorm:"size:100;not null;" json:"department"`
	CountOfTasks          uint32    `gorm:"default:0" json:"countOfTasks"`
	CountOfCompletedTasks uint32    `gorm:"default:0" json:"countOfCompletedTasks"`
	IsActive              bool      `gorm:"default:true" json:"isActive"`
	IsDeleted             bool      `gorm:"default:false" json:"isDeleted"`
	CreatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Department = html.EscapeString(strings.TrimSpace(u.Department))
	u.CountOfTasks = 0
	u.CountOfCompletedTasks = 0
	u.IsActive = true
	u.IsDeleted = false
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Required Username")
		}

		if u.Department == "" {
			return errors.New("Required Department")
		}

		if u.CountOfCompletedTasks < 0 {
			return errors.New("CountOfCompletedTasks must be greater than 0")
		}

		if u.CountOfTasks < 0 {
			return errors.New("CountOfTasks must be greater than 0")
		}

		if u.Password == "" {
			return errors.New("Required Password")
		}

		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil

	default:
		if u.Username == "" {
			return errors.New("Required Username")
		}

		if u.Password == "" {
			return errors.New("Required Password")
		}

		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB, pageNumber uint32) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(pageNumber * 100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":              u.Password,
			"nickname":              u.Username,
			"email":                 u.Email,
			"department":            u.Department,
			"is_active":             u.IsActive,
			"is_deleted":            u.IsDeleted,
			"countOfTasks":          u.CountOfTasks,
			"countOfCompletedTasks": u.CountOfCompletedTasks,
			"update_at":             time.Now(),
		})
	if db.Error != nil {
		return &User{}, db.Error
	}
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
