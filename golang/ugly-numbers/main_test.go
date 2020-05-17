package main

import (
	"testing"
)

func TestUgly(t *testing.T) {

	if !ugly(1) {
		t.Errorf("Ugly(1) returned false; wanted true")
	}

	if !ugly(6) {
		t.Errorf("Ugly(6) returned false; wanted true")
	}

	if ugly(7) {
		t.Errorf("Ugly(7) returned true; wanted false")
	}

	if !ugly(8) {
		t.Errorf("Ugly(8) returned false; wanted true")
	}

	if !ugly(10) {
		t.Errorf("Ugly(10) returned false; wanted true")
	}

	if ugly(14) {
		t.Errorf("Ugly(14) returned true; wanted false")
	}

	if ugly(-6) {
		t.Errorf("Ugly(-6) returned true; wanted false")
	}

	if ugly(33) {
		t.Errorf("Ugly(33) returned true; wanted false")
	}

	if ugly(42) {
		t.Errorf("Ugly(42) returned true; wanted false")
	}

}
