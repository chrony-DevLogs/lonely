package main

import (
	"fmt"
	"log"
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

}
