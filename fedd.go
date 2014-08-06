package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	userdata, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	savedir := userdata.HomeDir
	scandir := filepath.VolumeName(wd)
	if filepath.VolumeName(savedir) != scandir {
		savedir = scandir
	}
	savedir += string(filepath.Separator)
	scandir += string(filepath.Separator)

	log.Println("Scanning Volume:", scandir)
	log.Println("Saving Index to:", savedir)

	data := index(scandir)

	log.Println("Encoding Index")
	save(savedir+"index.fed", data)
	log.Println("Finished")
}
