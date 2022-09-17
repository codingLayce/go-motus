package game

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"wordle/utils/mock"
)

func TestGameLoadWords(t *testing.T) {
	expectedWords := map[int]int{
		4: 1997,
		5: 5070,
		6: 9938,
		7: 15178,
		8: 19076,
		9: 19555,
	}

	for wordSize, expectedCount := range expectedWords {
		t.Run(fmt.Sprintf("Load words of size %d", wordSize), func(t *testing.T) {
			game := Game{}

			err := game.loadWords(wordSize)
			require.NoError(t, err)

			require.Len(t, game.words, expectedCount)
			for _, word := range game.words {
				require.Len(t, word, wordSize)
			}
		})
	}
}

func TestGameLoadWords_Error(t *testing.T) {
	for name, tc := range map[string]struct {
		openFile      func(name string) (*os.File, error)
		expectedError string
	}{
		"Error while opening file": {
			openFile: func(name string) (*os.File, error) {
				return nil, errors.New("openFile")
			},
			expectedError: "openFile",
		},
	} {
		t.Run(name, func(t *testing.T) {
			mock.Do(t, &openFile, tc.openFile)

			game := Game{}
			err := game.loadWords(4)

			require.EqualError(t, err, tc.expectedError)
		})
	}
}

func TestPickRandomWord(t *testing.T) {
	words := []string{"petit", "grand", "moyen", "personne", "animal"}

	game := Game{}
	word := game.PickRandomWord(words)

	require.Contains(t, words, word)
}

func TestTryWord(t *testing.T) {
	for name, tc := range map[string]struct {
		mysteryWord     string
		attemptedWord   string
		expectedFound   bool
		expectedDetails []LetterDetail
		expectError     bool
	}{
		"Right word": {
			mysteryWord:   "petit",
			attemptedWord: "petit",
			expectedFound: true,
		},
		"Wrong word - No matching letter": {
			mysteryWord:     "petit",
			attemptedWord:   "grand",
			expectedFound:   false,
			expectedDetails: []LetterDetail{NotPresent, NotPresent, NotPresent, NotPresent, NotPresent},
		},
		"Wrong word - With placed letter": {
			mysteryWord:     "petit",
			attemptedWord:   "pales",
			expectedFound:   false,
			expectedDetails: []LetterDetail{Placed, NotPresent, NotPresent, Present, NotPresent},
		},
		"Wrong word - Not same length": {
			mysteryWord:     "petit",
			attemptedWord:   "length doesn't match the expected",
			expectedFound:   false,
			expectedDetails: []LetterDetail{Placed, NotPresent, NotPresent, Present, NotPresent},
			expectError:     true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			game := Game{}
			game.mystery = tc.mysteryWord

			found, details, err := game.AttemptWord(tc.attemptedWord)
			if tc.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			require.Equal(t, tc.expectedFound, found)

			require.Len(t, details, len(tc.expectedDetails))
			for i := range tc.expectedDetails {
				require.Equal(t, tc.expectedDetails[i], details[i])
			}
		})
	}
}
