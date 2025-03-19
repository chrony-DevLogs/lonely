package main

import "lonely/tparse"

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

	// data, err := tparse.ReadTorrentFile("GTAVCNE_EN_RU.torrent")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// torrentFIle, err := tparse.Open(data)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// torretfile := *torrentFIle

	// fmt.Println(torretfile["announce-list"])

	tparse.ParseMagnetQuery("magnet:?xt=urn:btih:FE60B29767946ECDCD087E5AC9E66E480C4755D1&dn=Dead+Cells%3A+Medley+of+Pain+Bundle+%28v35+%2B+5+DLCs+%2B+12+Bonus+OSTs%2C+MULTi13%29+%5BFitGirl+Repack%2C+Selective+Download+-+from+1.4+GB%5D&tr=udp%3A%2F%2Fopentor.net%3A6969&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.gbitt.info%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.ccp.ovh%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.ccp.ovh%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.theoks.net%3A6969%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=http%3A%2F%2Fopen.acgtracker.com%3A1096%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce")

}
