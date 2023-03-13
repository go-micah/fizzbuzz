package fizzbuzz

import "testing"

func TestCalculateFizz(t *testing.T) {

	want := "Fizz"
	got := Calculate(9)

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestCalculateBuzz(t *testing.T) {

	want := "Buzz"
	got := Calculate(5)

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestCalculateFizzBuzz(t *testing.T) {

	want := "FizzBuzz"
	got := Calculate(15)

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestCalculateOther(t *testing.T) {

	want := "7"
	got := Calculate(7)

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}
