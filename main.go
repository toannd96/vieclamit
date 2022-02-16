package main

import (
	"sync"
	"time"

	"vieclamit/database"
	"vieclamit/feeds"
	"vieclamit/handle"
	repoimpl "vieclamit/repository/repo_impl"
)

func main() {
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

	// search job by keyword
	// handle.SearchJobByLocation("Hà nội")
	// handle.SearchJobBySkill("golang")
	// handle.SearchJobByCompany("fpt")

	// url screenshot JD TopCV
	// feeds.ScreenshotJDTopCV("https://www.topcv.vn/viec-lam/blockchain-developers-luong-1-000-4-000-hcm/590697.html")
	// feeds.ScreenshotJDTopCV("https://www.topcv.vn/brand/smartosc/tuyen-dung/it-comtor-j592057.html")

	// schedule crawl
	go schedule(6*time.Hour, handle, 1)
	schedule(24*time.Hour, handle, 2)
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
