package power_of_two

import "testing"



func TestPowerOfTwo1(t *testing.T) {
    input := -2147483648 // -2^31

    result := isPowerOfTwo(input)
    assert(t, result, false)
}

func TestPowerOfTwo2(t *testing.T) {
    input := 2147483648 - 1 // 2^31-1

    result := isPowerOfTwo(input)
    assert(t, result, false)
}

func TestPowerOfTwo3(t *testing.T) {
    input := -2

    result := isPowerOfTwo(input)
    assert(t, result, false)
}

func TestPowerOfTwo4(t *testing.T) {
    input := 0

    result := isPowerOfTwo(input)
    assert(t, result, false)
}

func TestPowerOfTwo5(t *testing.T) {
    input := 1

    result := isPowerOfTwo(input)
    assert(t, result, true)
}

func TestPowerOfTwo6(t *testing.T) {
    input := 2

    result := isPowerOfTwo(input)
    assert(t, result, true)
}

func TestPowerOfTwo7(t *testing.T) {
    input := 3

    result := isPowerOfTwo(input)
    assert(t, result, false)
}

func TestPowerOfTwo8(t *testing.T) {
    input := 4

    result := isPowerOfTwo(input)
    assert(t, result, true)
}

func TestPowerOfTwo9(t *testing.T) {
    input := 16

    result := isPowerOfTwo(input)
    assert(t, result, true)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("want: %v; got: %v", want, result)
    }
}
