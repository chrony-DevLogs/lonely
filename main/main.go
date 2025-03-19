package main

import (
	"fmt"
	"log"
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

	// fileNumber := "5"

	// data, err := tparse.ReadTorrentFile("tparse/test_torrent_files/f" + fileNumber + ".torrent")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// torretfile, err := tparse.Open(data)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// file, err := os.Create("tparse/test_torrent_files/" + fileNumber + "Expected")
	// if err != nil {
	// 	log.Fatalf("Error creating file: %v", err)
	// }
	// defer file.Close()

	// encoder := gob.NewEncoder(file)
	// err = encoder.Encode(torretfile)
	// if err != nil {
	// 	log.Fatalf("Error encoding data: %v", err)
	// }
	// fmt.Println("Data written to file successfully in binary format!")

	// Step 1: Open the binary file for reading

	d, err := tparse.ParseMagnetQuery("magnet:?xt=urn:btih:c3d4e5f67890abcdef1234567890abcdef123456&dn=example_file_3.iso&tr=http://tracker.torrents.com/announce")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(d)
}
