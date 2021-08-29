package users

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var users = []User {
	User {
		ID: uuid.New().String(),
		Name: "Fajar Ardiyanto",
		Username: "fajar",
		Email: "fajar@localhost.com",
		Password: "ayambawang",
	},
	User {
		ID: uuid.New().String(),
		Name: "ADMIN",
		Username: "admin",
		Email: "admin@localhost.com",
		Password: "ayambawang",
	},
}

func LoadSeed(db *gorm.DB) {
	var logger log.Logger

	if ok := db.HasTable(&User{}); !ok {
		if err := db.Debug().CreateTable(&User{}).Error; err != nil {
			level.Error(logger).Log("cannot create table: ", err.Error())
		}

		if err := db.Debug().AutoMigrate(&User{}).Error; err != nil {
			level.Error(logger).Log("cannot migrate table: ", err.Error())
		}

		for i, _ := range users {
			if err := db.Debug().Model(&User{}).Create(&users[i]).Error; err != nil {
				level.Error(logger).Log("cannot seed users table: ", err.Error())
			}
		}
	}
}
