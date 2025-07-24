package main

import (
	"math/rand"
	"testing"

	"github.com/shopspring/decimal"
)

func TestToyExample(t *testing.T) {
	alpha := decimal.NewFromInt(4)
	beta := decimal.NewFromInt(9)
	p := decimal.NewFromInt(23)
	q := decimal.NewFromInt(11)

	x := decimal.NewFromInt(6)
	k := decimal.NewFromInt(7)

	c := decimal.NewFromInt(4)

	y1 := alpha.Pow(x).Mod(p)
	y2 := beta.Pow(x).Mod(p)

	if y1.Cmp(decimal.NewFromInt(2)) != 0 {
		t.Errorf("Expected y1 to be 2, got %s", y1.String())
	}
	if y2.Cmp(decimal.NewFromInt(3)) != 0 {
		t.Errorf("Expected y2 to be 3, got %s", y2.String())
	}

	r1 := alpha.Pow(k).Mod(p)
	r2 := beta.Pow(k).Mod(p)

	if r1.Cmp(decimal.NewFromInt(8)) != 0 {
		t.Errorf("Expected r1 to be 8, got %s", r1.String())
	}
	if r2.Cmp(decimal.NewFromInt(4)) != 0 {
		t.Errorf("Expected r2 to be 4, got %s", r2.String())
	}

	s := Solve(k, c, x, q)
	expectedS := decimal.NewFromInt(5)
	if s.Cmp(expectedS) != 0 {
		t.Errorf("Expected s to be %s, got %s", expectedS.String(), s.String())
	}

	result := Verify(r1, r2, y1, y2, alpha, beta, s, c, p)
	if !result {
		t.Error("Verification failed, expected true but got false")
	}

	// Fake secret
	xFake := decimal.NewFromInt(7)
	sFake := Solve(k, c, xFake, q)

	result = Verify(r1, r2, y1, y2, alpha, beta, sFake, c, p)
	if result {
		t.Error("Verification should have failed with fake secret, expected false but got true")
	}

}

func TestToyExampleWithRNG(t *testing.T) {
	alpha := decimal.NewFromInt(4)
	beta := decimal.NewFromInt(9)
	p := decimal.NewFromInt(23)
	q := decimal.NewFromInt(11)

	x := decimal.NewFromInt(6)
	k := decimal.NewFromInt(int64(rand.Intn(int(q.IntPart()))))

	c := decimal.NewFromInt(int64(rand.Intn(int(q.IntPart()))))

	y1 := alpha.Pow(x).Mod(p)
	y2 := beta.Pow(x).Mod(p)

	if y1.Cmp(decimal.NewFromInt(2)) != 0 {
		t.Errorf("Expected y1 to be 2, got %s", y1.String())
	}
	if y2.Cmp(decimal.NewFromInt(3)) != 0 {
		t.Errorf("Expected y2 to be 3, got %s", y2.String())
	}

	r1 := alpha.Pow(k).Mod(p)
	r2 := beta.Pow(k).Mod(p)

	s := Solve(k, c, x, q)

	result := Verify(r1, r2, y1, y2, alpha, beta, s, c, p)
	if !result {
		t.Error("Verification failed, expected true but got false")
	}

}
