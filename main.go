package main

import (
	"sync"
	"time"

	"vieclamit/database"
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
	wg.Add(2)

	go func() {
		defer wg.Done()
		// feeds.TopCV(handle.Repo)
	}()

	go func() {
		defer wg.Done()
		// feeds.CareerBuilder(handle.Repo)
	}()

	wg.Wait()

	// Schedule crawl
	go schedule(24*time.Hour, handle, 1)
	schedule(24*time.Hour, handle, 2)
}

func schedule(timeSchedule time.Duration, handle handle, inndex int) {
	ticker := time.NewTicker(timeSchedule)
	func() {
		for {
			switch inndex {
			case 1:
				<-ticker.C

			case 2:
				<-ticker.C
			}
		}
	}()
}
