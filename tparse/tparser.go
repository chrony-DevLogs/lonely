package tparse

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/zeebo/bencode"
)

type BencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type BencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     BencodeInfo `bencode:"info"`
}

func Open(torrentData []byte) (*BencodeTorrent, error) { // parse a torrent file into a map
	bto := BencodeTorrent{}

	err := bencode.DecodeBytes(torrentData, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

func ReadTorrentFile(filePath string) ([]byte, error) { // open a torrent file using the os FS
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("reading file err: %w", err)
	}

	return data, nil
}

func ParseMagnetQuery(magnet string) (map[string]string, error) { // coverts a magnet link to a map of data

	prefix := "magnet:?"

	if !strings.Contains(magnet, prefix) {

		return nil, errors.New("error not a magnet link")
	}
	magnet = strings.Replace(magnet, prefix, "", 1)

	parsedLink, err := url.ParseQuery(magnet)

	if err != nil {
		return nil, err
	}

	magnetData := make(map[string]string)

	for k, v := range parsedLink {
		magnetData[k] = v[0]
	}

	return magnetData, nil

}
