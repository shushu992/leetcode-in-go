package pow

import (
    "reflect"
    "testing"
)

func TestMyPow1(t *testing.T) {
    result := myPow(0, 10)
    want := 0.0

    assert(t, result, want)
}

func TestMyPow2(t *testing.T) {
    result := myPow(-1, 10)
    want := 1.0

    assert(t, result, want)
}

func TestMyPow3(t *testing.T) {
    result := myPow(2.00000, 10)
    want := 1024.00000

    assert(t, result, want)
}

func TestMyPow4(t *testing.T) {
    result := myPow(2.10000, 3)
    want := 9.26100

    assert(t, result, want)
}

func TestMyPow5(t *testing.T) {
    result := myPow(2.00000, -2)
    want := 0.25000

    assert(t, result, want)
}

func assert(t *testing.T, result float64, want float64) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("result: %f; want: %f", result, want)
    }
}
