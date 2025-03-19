package tparse_test

import (
	"lonely/tparse"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var torrentFiles = map[string]string{
	"f1": "test_torrent_files/f1.torrent",
	"f2": "test_torrent_files/f2.torrent",
	"f3": "test_torrent_files/f3.torrent",
	"f4": "test_torrent_files/f4.torrent",
	"f5": "test_torrent_files/f5.torrent",
}

func TestReadTorrentFile(t *testing.T) {
	type testCase struct {
		// Input Parms
		filePath string

		//Expected Values
		expected []byte
	}

	tests := []testCase{
		{filePath: torrentFiles["f1"]},
		{filePath: torrentFiles["f2"]},
		{filePath: torrentFiles["f3"]},
		{filePath: torrentFiles["f4"]},
		{filePath: torrentFiles["f5"]},
	}

	for i, test := range tests {
		data, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Fatalf("Error reading expected test file (%s): %v", test.filePath, err)
		}
		tests[i].expected = data
	}

	for _, test := range tests {
		t.Run("testing the open function", func(t *testing.T) {
			actual, err := tparse.ReadTorrentFile(test.filePath)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}

}

func TestOpen(t *testing.T) {
	type testCase struct {
		// Input Parms
		torrentData []byte

		//Expected Values
		expected tparse.BencodeTorrent
	}

	f1data, err := tparse.ReadTorrentFile(torrentFiles["f1"])

	if err != nil {
		t.Fatalf("error reading torrent file %v", err)
	}
	f2data, err := tparse.ReadTorrentFile(torrentFiles["f2"])

	if err != nil {
		t.Fatalf("error reading torrent file %v", err)
	}

	f3data, err := tparse.ReadTorrentFile(torrentFiles["f3"])

	if err != nil {
		t.Fatalf("error reading torrent file %v", err)
	}

	f4data, err := tparse.ReadTorrentFile(torrentFiles["f4"])

	if err != nil {
		t.Fatalf("error reading torrent file %v", err)
	}

	f5data, err := tparse.ReadTorrentFile(torrentFiles["f5"])

	if err != nil {
		t.Fatalf("error reading torrent file %v", err)
	}

	fs1 := tparse.BencodeTorrent{
		"http://tracker.example.com/announce",
		tparse.BencodeInfo{
			"abcdef1234567890",
			262144,
			10485760,
			"example_file_1.txt",
		},
	}

	fs2 := tparse.BencodeTorrent{
		"http://tracker.test.com/announce",
		tparse.BencodeInfo{
			"123456abcdef7890",
			524288,
			20971520,
			"example_file_2.bin",
		},
	}

	fs3 := tparse.BencodeTorrent{
		"http://tracker.torrents.com/announce",
		tparse.BencodeInfo{
			"9a8b7c6d5e4f3g2h",
			131072,
			7340032,
			"example_file_3.iso",
		},
	}

	fs4 := tparse.BencodeTorrent{
		"http://tracker.speed.net/announce",
		tparse.BencodeInfo{
			"0a1b2c3d4e5f6g7h",
			1048576,
			41943040,
			"example_file_4.mp4",
		},
	}

	fs5 := tparse.BencodeTorrent{
		"http://tracker.fasttrack.org/announce",
		tparse.BencodeInfo{
			"deadbeefcafebabe",
			65536,
			5242880,
			"example_file_5.pdf",
		},
	}

	tests := []testCase{
		{torrentData: f1data, expected: fs1},
		{torrentData: f2data, expected: fs2},
		{torrentData: f3data, expected: fs3},
		{torrentData: f4data, expected: fs4},
		{torrentData: f5data, expected: fs5},
	}

	for _, test := range tests {
		t.Run("testing the open function", func(t *testing.T) {
			actual, err := tparse.Open(test.torrentData)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, *actual)
		})
	}
}

// func TestParseMagnetQuery(t *testing.T) {
// 	//
// }
