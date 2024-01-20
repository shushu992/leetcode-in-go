package number_of_1_bits

import "testing"

func TestHammingWeight1(t *testing.T) {
    var input uint32 = 0b0

    result := hammingWeight(input)
    want := 0

    assert(t, result, want)
}

func TestHammingWeight2(t *testing.T) {
    var input uint32 = 0b11111111111111111111111111111111

    result := hammingWeight(input)
    want := 32

    assert(t, result, want)
}

func TestHammingWeight3(t *testing.T) {
    var input uint32 = 0b00000000000000000000000000001011

    result := hammingWeight(input)
    want := 3

    assert(t, result, want)
}

func TestHammingWeight4(t *testing.T) {
    var input uint32 = 0b00000000000000000000000010000000

    result := hammingWeight(input)
    want := 1

    assert(t, result, want)
}

func TestHammingWeight5(t *testing.T) {
    var input uint32 = 0b11111111111111111111111111111101

    result := hammingWeight(input)
    want := 31

    assert(t, result, want)
}

func assert(t *testing.T, result int, want int) {
    t.Helper()

    if result != want {
        t.Fatalf("want: %d; result: %d", want, result)
    }
}
