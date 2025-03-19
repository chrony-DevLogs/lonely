package tparse

import (
	"fmt"
	"net/url"
	"os"

	"github.com/zeebo/bencode"
)

// type bencodeInfo struct {
// 	Pieces      string `bencode:"pieces"`
// 	PieceLength int    `bencode:"piece length"`
// 	Length      int    `bencode:"length"`
// 	Name        string `bencode:"name"`
// }

// type bencodeTorrent struct {
// 	Announce string      `bencode:"announce"`
// 	Info     bencodeInfo `bencode:"info"`
// }

func Open(torrentData []byte) (*map[string]any, error) { // parse a torrent file into a map
	var bto map[string]any

	err := bencode.DecodeBytes(torrentData, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

func ReadTorrentFile(fileName string) ([]byte, error) { // open a torrent file using the os FS
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("reading file err: %w", err)
	}

	return data, nil
}

func ParseMagnetQuery(magnet string) (map[string][]string, error) { // coverts a magnet link to a map of data
	parsedLink, err := url.ParseQuery(magnet)

	if err != nil {
		return nil, err
	}

	magnetData := make(map[string][]string)

	for k, v := range parsedLink {
		magnetData[k] = v
	}

	return magnetData, nil

}
