package initial

import(
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .evc file!")
	}
}