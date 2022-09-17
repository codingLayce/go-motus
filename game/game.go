package game

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"wordle/utils"
)

var AttemptLengthUnexpected = errors.New("attempt word length is unexpected")

const wordsPathPattern = "../data/words_%d.txt"

type LetterDetail int

const (
	NotPresent LetterDetail = iota
	Present
	Placed
)

type Game struct {
	words   []string
	mystery string
}

var openFile = func(path string) (*os.File, error) {
	return os.Open(path)
}

func (g *Game) loadWords(wordSize int) error {
	file, err := openFile(fmt.Sprintf(wordsPathPattern, wordSize))
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		g.words = append(g.words, scanner.Text())
	}

	return scanner.Err()
}

func (g *Game) PickRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

func (g *Game) AttemptWord(attempt string) (bool, []LetterDetail, error) {
	if attempt == g.mystery {
		return true, nil, nil
	}

	if len(attempt) != len(g.mystery) {
		return false, nil, AttemptLengthUnexpected
	}

	details := make([]LetterDetail, 0, len(g.mystery))
	cache := make(map[rune]int)
	mysteryChars := []rune(g.mystery)
	attemptChars := []rune(attempt)

	for i := 0; i < len(attemptChars); i++ {
		skip := 0
		skip, _ = cache[attemptChars[i]]

		idx := utils.IndexOf(mysteryChars, attemptChars[i], skip)
		if idx == -1 {
			details = append(details, NotPresent)
		} else if idx == i {
			details = append(details, Placed)
		} else {
			details = append(details, Present)
		}

		cache[attemptChars[i]] += 1
	}

	return false, details, nil
}
