package reverse_bits

import "testing"

func TestReverseBits1(t *testing.T) {
    var input uint32 = 0

    result := reverseBits(input)
    var want uint32 = 0

    assert(t, result, want)
}

func TestReverseBits2(t *testing.T) {
    var input uint32 = 0b11111111111111111111111111111111

    result := reverseBits(input)
    var want uint32 = 0b11111111111111111111111111111111

    assert(t, result, want)
}

func TestReverseBits3(t *testing.T) {
    var input uint32 = 0b00000000000000000000000000000100

    result := reverseBits(input)
    var want uint32 = 0b00100000000000000000000000000000

    assert(t, result, want)
}

func TestReverseBits4(t *testing.T) {
    var input uint32 = 0b00000010100101000001111010011100

    result := reverseBits(input)
    var want uint32 = 0b00111001011110000010100101000000

    assert(t, result, want)
}

func TestReverseBits5(t *testing.T) {
    var input uint32 = 0b11111111111111111111111111111101

    result := reverseBits(input)
    var want uint32 = 0b10111111111111111111111111111111

    assert(t, result, want)
}

func assert(t *testing.T, result uint32, want uint32) {
    t.Helper()

    if result != want {
        t.Fatalf("\nwant: %b\ngot:  %b", want, result)
    }
}
