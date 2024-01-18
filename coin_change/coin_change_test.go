package coin_change

import "testing"

func TestCoinChange1(t *testing.T) {
    coins := []int{1}
    amount := 0

    result := coinChange(coins, amount)
    want := 0

    assert(t, result, want)
}

func TestCoinChange2(t *testing.T) {
    coins := []int{2}
    amount := 3

    result := coinChange(coins, amount)
    want := -1

    assert(t, result, want)
}

func TestCoinChange3(t *testing.T) {
    coins := []int{1, 2, 5}
    amount := 11

    result := coinChange(coins, amount)
    want := 3

    assert(t, result, want)
}

func TestCoinChange4(t *testing.T) {
    coins := []int{2, 5, 10, 1}
    amount := 27

    result := coinChange(coins, amount)
    want := 4

    assert(t, result, want)
}

func assert(t *testing.T, result int, want int) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %d; want: %d", result, want)
    }
}
