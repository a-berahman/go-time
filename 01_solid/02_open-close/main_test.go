package main

import (
	"fmt"
	"testing"
)

func TestSordDamage(t *testing.T) {
	specs := []struct {
		sword interface {
			Damage() int
		}
		exp int
	}{
		{
			sword: Sword{name: "Silver"},
			exp:   2,
		}, {
			sword: MagicSword{Sword: Sword{name: "Dragon"}},
			exp:   42,
		},
	}

	for index, v := range specs {
		if got := v.sword.Damage(); got != v.exp {
			t.Errorf("%d expected to get damage %d, got %d", index, v.exp, got)
		}

	}
}

func TestSordDamageToString(t *testing.T) {
	specs := []struct {
		sword fmt.Stringer
		exp   string
	}{
		{
			sword: Sword{name: "Silver"},
			exp:   "Silver is a sword that can deal 2 points to oppnents",
		}, {
			sword: MagicSword{Sword: Sword{name: "Dragon"}},
			exp:   "Dragon is a sword that can deal 42 points to oppnents",
		},
	}

	for index, v := range specs {
		if got := v.sword.String(); got != v.exp {
			t.Errorf("%d expected to get\n %q,\n got\n %q", index, v.exp, got)
		}

	}
}
