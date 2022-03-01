package main

import (
	"os"
	"time"

	"vieclamit/broadcasts"
	"vieclamit/database"
	"vieclamit/feeds"
	"vieclamit/handle"
	repoimpl "vieclamit/repository/repo_impl"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		return
	}
}

func main() {

	// conn db
	mg := &database.Mongo{}
	mg.CreateConn()

	handle := handle.Handle{
		Repo: repoimpl.NewRepo(mg),
	}

	// conn telegram
	telegramConfig := &broadcasts.Telegram{
		Token: os.Getenv("TELEGRAM_TOKEN"),
		Repo:  repoimpl.NewRepo(mg),
	}
	go telegramConfig.NewTelegram()

	// run crawl
	go feeds.TopCV(handle.Repo)

	// schedule crawl
	go schedule(6*time.Hour, handle, 1)
	schedule(10*time.Hour, handle, 2)
}

func schedule(timeSchedule time.Duration, handle handle.Handle, index int) {
	ticker := time.NewTicker(timeSchedule)
	func() {
		for {
			switch index {
			case 1:
				<-ticker.C
				feeds.TopCV(handle.Repo)
			case 2:
				<-ticker.C
				handle.CheckJobDeadlineExpired()
			}
		}
	}()
}
