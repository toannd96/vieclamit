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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		feeds.TopCV(handle.Repo)
	}()

	wg.Wait()

	// handle.SearchJobByLocation("Hà nội")
	// handle.SearchJobBySkill("golang")
	// handle.SearchJobByCompany("fpt")

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
