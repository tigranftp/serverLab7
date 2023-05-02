package main

import (
	"db_lab7/API"
	"fmt"
)

func main() {

	API, err := API.NewAPI()
	if err != nil {
		fmt.Println(err)
	}

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//go func() {
	//	select {
	//	case <-c:
	//		API.Stop()
	//		return
	//	}
	//}()

	err = API.Start()
	if err != nil {
		fmt.Println(err)
	}
}
