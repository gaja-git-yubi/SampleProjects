package main
 
import "testing"

func TestAdd(t *testing.T) {
    got := add(2, 3)
    want := 5 

    if got != want {
        t.Errorf("add(2, 3) = %d; want %d", got, want)
    }
}


func TestSubtract(t *testing.T) {
    got := subtract(5, 3)
    want := 2

    if got != want {
        t.Errorf("subtract(5, 3) = %d; want %d", got, want)
    }
}

func TestMultiply(t *testing.T) {
    got := multiply(2, 3)
    want := 6

    if got != want {
        t.Errorf("multiply(2, 3) = %d; want %d", got, want)
    }
}

func TestDivide(t *testing.T) {
    got := divide(6, 3)
    want := 2

    if got != want {
        t.Errorf("divide(6, 3) = %d; want %d", got, want)
    }
}

func TestModulo(t *testing.T) {
    got := modulus(6, 4)
    want := 2

    if got != want {
        t.Errorf("modulus(6, 3) = %d; want %d", got, want)
    }
}