package game

import (
	"fmt"
	"testing"
)

func TestEnchantedSword_Damage(t *testing.T) {
	type fields struct {
		Sword Sword
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "sljfa",
			fields: fields{
				Sword: Sword{},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := EnchantedSword{
				Sword: tt.fields.Sword,
			}
			if got := receiver.Damage(); got != tt.want {
				t.Errorf("Damage() = %v, want %v", got, tt.want)
			}

		})
	}
}

type tstInt interface {
	Damage() int
}

func passInt(dae tstInt) {
	fmt.Println(dae)
}

func TestSwordDamage(t *testing.T) {

	e := EnchantedSword{}
	fmt.Println(e.Damage())
	passInt(e)

	specs := []struct {
		sword tstInt
		exp   int
	}{
		{
			sword: Sword{name: "Silver Saber"},
			exp:   2,
		},
		{
			sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
			exp:   42,
		},
	}
	for specIndex, spec := range specs {
		if got := spec.sword.Damage(); got != spec.exp {
			t.Errorf("[spec %d] expected to get damage %d; got %d",
				specIndex, spec.exp, got)
		}
	}
}

func TestNoShadowed(t *testing.T) {
	specs := []struct {
		sword fmt.Stringer
		exp   string
	}{
		{
			sword: Sword{name: "Silver Saber"},
			exp: "Silver Saber is a sword that can deal 2 points of damage to opponents",
		},
		{
			//no shadowing
			sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
			exp: "Dragon's Greatsword is a sword that can deal 2 points of damage to opponents",
		},
	}
	for specIndex, spec := range specs {
		if got := spec.sword.String(); got != spec.exp {
			t.Errorf("[spec %d] expected to get\n%q\ngot:\n%q", specIndex, spec.exp, got)
		}
	}
}
