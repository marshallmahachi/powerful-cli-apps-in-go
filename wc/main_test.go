package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {

	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {

	b := bytes.NewBufferString("word1\nword2\nword3")

	exp := 3

	res := count(b, true)

	if exp != res {
		t.Errorf("Expected %d, got %d", exp, res)
	}

}
