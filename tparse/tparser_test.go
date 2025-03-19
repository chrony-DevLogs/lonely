package tparse_test

import (
	"lonely/tparse"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestOpen(t *testing.T) {
// 	type testCase struct {
// 		// Input Parms
// 		torrentData []byte

// 		//Expected Values
// 		expected map[string]any
// 		err      error
// 	}

//		t.Run("testing the open function", func(t *testing.T) {
//			//
//		})
//	}
func TestReadTorrentFile(t *testing.T) {
	type testCase struct {
		// Input Parms
		filePath string

		//Expected Values
		expected []byte
	}

	tests := []testCase{{filePath: "../torrentFiles/debian-12.10.0-amd64-netinst.iso.torrent"}, {filePath: "../torrentFiles/GTAVCNE_EN_RU.torrent"}, {filePath: "../torrentFiles/sleepingDogs.torrent"}}

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

// func TestParseMagnetQuery(t *testing.T) {
// 	//
// }
