package main

import (
	"flag"
	"fmt"
	"github.com/reco/pkg/containter"
	"github.com/reco/pkg/dto"
	"github.com/reco/pkg/services"
	"sync"
)

const workers = 3

func main() {
	limit := flag.Int("limit", 1, "maximum amount users per page to extract")
	flag.Parse()

	wg := &sync.WaitGroup{}

	c := containter.Get()
	list, nextOffset, err := c.DataProvider.GetUsersList("", *limit)
	if err != nil {
		//@todo: add logs
		fmt.Println("Failed to get users")
		return
	}

	if err = storeUsers(list); err != nil {
		//@todo: add logs
		return
	}

	var offsets = make(chan string, workers)
	offsets <- nextOffset

	processUsers(offsets, wg, *limit)

	wg.Wait()
	fmt.Println("Done")
}

func processUsers(offsets chan string, wg *sync.WaitGroup, limit int) {
	//we run several workers
	for num := range [workers][]int{} {

		wg.Add(1)

		//run the routine
		//	in each routine we trigger request to API and retrieve users
		//  whenever we pull the users, we receive next offset
		//	each offset we put back into the channel so other routines process the next page
		//	once we receive a null offset, we close channel and exit routines
		//	each data we receive, we put into the file(this also can be extracted to a separate routine

		go worker(wg, num, offsets, limit)
	}

	return
}

func worker(wg *sync.WaitGroup, num int, offsets chan string, limit int) {
	defer wg.Done()

	fmt.Println(fmt.Sprintf("Running worker: #%d", num))

	c := containter.Get()
	for offset := range offsets {
		list, nextOffset, err := c.DataProvider.GetUsersList(offset, limit)
		if err != nil {
			fmt.Println("Failed to load the data")
			//@todo: handle the issue
			continue
		}

		if nextOffset != "" {
			offsets <- nextOffset
		} else {
			close(offsets)
		}

		if err = storeUsers(list); err != nil {
			fmt.Println("failed to store users")
			continue
		}
	}

	fmt.Println(fmt.Sprintf("Finished worker: #%d", num))
}

func storeUsers(list []dto.UserDataItem) error {
	c := containter.Get()

	for _, user := range list {
		if err := c.UsersService.Store(services.User{
			ID:   user.Gid,
			Name: user.Name,
		}); err != nil {
			return err
		}
	}

	return nil
}
