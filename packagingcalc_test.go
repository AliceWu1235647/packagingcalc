package packagingcalc

import "testing"

func TestPPWRAndCBM(t *testing.T) {
	if cbm := CartonCBM(500, 400, 300, 10); cbm != 0.6 {
		t.Fatalf("expected 0.6 CBM, got %v", cbm)
	}
	pct, ok := PPWREmptySpaceRatio(1000, 600)
	if !ok || pct != 40 {
		t.Fatalf("expected 40%% compliant, got %v %v", pct, ok)
	}
	if g := PPWRRecyclabilityGrade(96.2); g != "A" {
		t.Fatalf("expected Grade A, got %s", g)
	}
}
