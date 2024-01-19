package reverse_words_in_a_string

/**
 * https://leetcode.com/problems/reverse-words-in-a-string/
 *
 * 1 <= s.length <= 10^4
 * s contains English letters (upper-case and lower-case), digits, and spaces ' '.
 * There is at least one word in s.
 */
func reverseWords(s string) string {
    input := []byte(s)
    size1 := len(input)

    boundary := make([]int, size1+1)
    size2 := 0

    word := false

    for i := 0; i < size1; i++ {
        if input[i] == 32 { // space
            if word {
                word = false

                boundary[size2] = i
                size2++
            }
        } else {
            if !word {
                word = true

                boundary[size2] = i
                size2++
            }
        }
    }

    if word {
        boundary[size2] = size1
        size2++
    }

    output := make([]byte, size1)
    size3 := 0

    for i := size2 - 1; ; {
        start := boundary[i-1]
        end := boundary[i]

        for {
            if start == end {
                break
            }

            output[size3] = input[start]
            size3++
            start++
        }

        i -= 2

        if i < 0 {
            break
        }

        output[size3] = 32
        size3++
    }

    return string(output[:size3])
}
