package main

import (
	"fmt"
	"log"
	"lonely/project"
	"lonely/tparse"
)

func main() {
	btf, err := tparse.ReadTorrentFile("elligalTorrentFiles/sleepingDogs.torrent")
	if err != nil {
		log.Fatal(err)
	}
	bto, err := tparse.Open(btf)

	file, err := bto.ToTorrentFile()

	fmt.Println(file.InfoHash)

	p, err := project.GetProjectData()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

}
