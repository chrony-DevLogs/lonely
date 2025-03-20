package tparse

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
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

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func (t *TorrentFile) buildTrackerUrl(peerID [20]byte, port uint) (string, error) {
	base, err := url.Parse(t.Announce)

	if err != nil {
		return "", err
	}

	queries := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}

	base.RawQuery = queries.Encode()
	return base.String(), nil
}

func (i *BencodeInfo) hash() [20]byte {

	buf, err := bencode.EncodeBytes(i)

	if err != nil {
		log.Fatal(err)
	}
	h := sha1.Sum(buf)
	return h
}

func resolveSHA1Pieces(bto *BencodeTorrent) [][20]byte {
	const sah1Size = 20
	numberOfPieces := len(bto.Info.Pieces) / sah1Size

	pieces := make([][20]byte, 0, numberOfPieces)

	for i := 0; i < numberOfPieces; i++ {
		var sh1 [20]byte
		for j := i * 20; j < 20*(i+1); j++ {
			sh1[j%20] = bto.Info.Pieces[j]
		}
		pieces = append(pieces, sh1)
	}

	return pieces
}

func (bto *BencodeTorrent) ToTorrentFile() (TorrentFile, error) {

	torrentFile := TorrentFile{}

	torrentFile.Announce = bto.Announce
	torrentFile.InfoHash = bto.Info.hash()
	torrentFile.PieceHashes = resolveSHA1Pieces(bto)
	torrentFile.PieceLength = bto.Info.Length
	torrentFile.Length = bto.Info.Length
	torrentFile.Name = bto.Info.Name
	return torrentFile, nil
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
