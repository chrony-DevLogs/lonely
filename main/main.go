package main

import (
	"fmt"
	"lonely/tparse"
)

func main() {

	//os file read write worm up

	// get user input!
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("write anything: ")
	// userInput, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Errorf("writing file err", err)
	// }
	// // store to file!
	// os.WriteFile("newfile.txt", []byte(userInput), os.FileMode.Perm(0644))

	// read from eleps.torrent

	data, err := tparse.ReadTorrentFile("torrentFiles/sleepingDogs.torrent")

	if err != nil {
		fmt.Println(err)
	}

	torrentFile, err := tparse.Open(data)

	if err != nil {
		fmt.Println(err)
	}

	torretfile := torrentFile

	fmt.Println(torretfile)

}
