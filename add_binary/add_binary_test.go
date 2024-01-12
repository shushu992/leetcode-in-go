package add_binary

import (
    "testing"
)

func TestAddBinary1(t *testing.T) {
    result := addBinary("0", "0")
    want := "0"

    assert(t, result, want)
}

func TestAddBinary2(t *testing.T) {
    result := addBinary("0", "1")
    want := "1"

    assert(t, result, want)
}

func TestAddBinary3(t *testing.T) {
    result := addBinary("1", "1")
    want := "10"

    assert(t, result, want)
}

func TestAddBinary4(t *testing.T) {
    result := addBinary("11", "1")
    want := "100"

    assert(t, result, want)
}

func TestAddBinary5(t *testing.T) {
    result := addBinary("1010", "1011")
    want := "10101"

    assert(t, result, want)
}

func TestAddBinary6(t *testing.T) {
    result := addBinary("1111", "1111")
    want := "11110"

    assert(t, result, want)
}

func assert(t *testing.T, result string, want string) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %s; want: %s", result, want)
    }
}
