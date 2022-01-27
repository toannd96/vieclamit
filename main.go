package main

import (
	"sync"
	"time"

	"vieclamit/database"
	"vieclamit/feeds"
	"vieclamit/repository"
	repoimpl "vieclamit/repository/repo_impl"
)

type handle struct {
	Repo repository.Repository
}

func main() {
	mg := &database.Mongo{}
	mg.CreateConn()

	handle := handle{
		Repo: repoimpl.NewRepo(mg),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		feeds.GetDataOnePage(handle.Repo)
	}()

	wg.Wait()

	// Schedule crawl
	schedule(6*time.Hour, handle, 1)
}

func schedule(timeSchedule time.Duration, handle handle, inndex int) {
	ticker := time.NewTicker(timeSchedule)
	func() {
		for {
			switch inndex {
			case 1:
				<-ticker.C
				feeds.GetDataOnePage(handle.Repo)
			}
		}
	}()
}
