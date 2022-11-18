package internal

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ProgressFile(done chan int64, destination string, total int64) {
	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:
			file, err := os.Open(destination)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100

			fmt.Println(int(percent))
		}
		if stop {
			fmt.Println(100)
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}

}
