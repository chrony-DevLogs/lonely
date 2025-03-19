package tparse

import (
	"fmt"
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

func Open(torrentData []byte) (*map[string]any, error) {
	var bto map[string]any

	err := bencode.DecodeBytes(torrentData, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

func ReadTorrentFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("reading file err: %w", err)
	}

	return data, nil
}
