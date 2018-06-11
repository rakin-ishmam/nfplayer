package model_test

import (
	"testing"

	"github.com/rakin-ishmam/nfplayer/model"
)

func TestFullName(t *testing.T) {
	p := model.Player{
		FName: "first",
		LName: "last",
	}

	exp := "first last"
	res := p.FullName()
	if exp != res {
		t.Errorf("expected full name is %v but got %v", exp, res)
	}
}
