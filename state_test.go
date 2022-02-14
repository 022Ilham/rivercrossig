package state

import (
	"fmt"
	"testing"
)

func TestViewState(t *testing.T) {
	wanted := "[Kylling rev korn hs ---\\ \\__/ _________________/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q,ønsket %q.", state, wanted)
	}
}

func TestPut(t *testing.T) {
	wanted := "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
	got := Put("korn")
	if got != wanted {
		t.Errorf("Feil,fikk %q, ønsket %q.", got, wanted)
	}
}
func TestCrossRiverTo(t *testing.T) {
	wanted := "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
	Put("korn")
	CrossRiverTo("east")
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q,ønsket %q.", state, wanted)
	}
}
func TestCharString(t *testing.T) {
	t.Run("hs", func(t *testing.T) {
		c := HS
		got := fmt.Sprintf("%s", c)
		expected := "hs"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("kylling", func(t *testing.T) {
		c := kylling
		got := fmt.Sprintf("%s", c)
		expected := "kylling"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("korn", func(t *testing.T) {
		c := korn
		got := fmt.Sprintf("%s", c)
		expected := "korn"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("rev", func(t *testing.T) {
		c := rev
		got := fmt.Sprintf("%s", c)
		expected := "rev"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}

func TestCharEats(t *testing.T) {
	t.Run("rev spiser kylling når hs er borte", func(t *testing.T) {
		s := state{west, east, east, east}
		got := rev.eats(kylling, s)
		expected := true
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("rev spiser ikke kylling når hs er tilstedet", func(t *testing.T) {
		s := state{east, east, east, east}
		got := rev.eats(kylling, s)
		expected := false
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("kylling spiser korn når hs er borte", func(t *testing.T) {
		s := state{east, west, west, west}
		got := kylling.eats(korn, s)
		expected := true
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("Kylling spiser ikke korn når hs er tilstedet", func(t *testing.T) {
		s := state{east, west, west, west}
		got := kylling.eats(korn, s)
		expected := true
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})
}
