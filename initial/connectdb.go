package initial

import(
	"gorm.io/gorm"

	"os"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB(){
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URI")), &gorm.Config{TranslateError: true})

	if err != nil {
		return
	}

	DB = db
}