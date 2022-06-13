package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gojp/kana"
	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	text := string(b)
	t := tokenizer.New()
	tokens := t.Tokenize(text)

	words := []string{}
	for _, token := range tokens {
		if token.Surface == "BOS" || token.Surface == "EOS" {
			continue
		}
		words = append(words, token.Surface)
	}

	for _, word := range words {
		tmp := word
		if kana.IsKana(tmp) || kana.IsHiragana(tmp) {
			tmp = kana.KanaToRomaji(tmp)
		}
		m := map[rune]int{
			'a': 0,
			'i': 0,
			'u': 0,
			'e': 0,
			'o': 0,
		}
		for _, r := range []rune(tmp) {
			if _, ok := m[r]; ok {
				m[r] += 1
			}
		}
		c := 0
		for _, v := range m {
			if v == 1 {
				c += v
			}
		}
		if c == 5 {
			fmt.Println(word)
		}
	}
}
