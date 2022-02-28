package main

import (
	"os"
	"sync"
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

	// run crawl
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		feeds.TopCV(handle.Repo)
	}()
	wg.Wait()

	// conn telegram
	telegramConfig := &broadcasts.Telegram{
		Token: os.Getenv("TELEGRAM_TOKEN"),
		Repo:  repoimpl.NewRepo(mg),
	}
	telegramConfig.NewTelegram()

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
