import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	used_letters := ""

	biggest_substring_length := 0

	for _, char := range s {

		char_string := string(char)

		if strings.Contains(used_letters, char_string) {

			if len(used_letters) > biggest_substring_length {
				biggest_substring_length = len(used_letters)
			}

			used_letters = removeCharFromS(used_letters, char_string)
		}

		used_letters += char_string

	}

	if len(used_letters) > biggest_substring_length {
		biggest_substring_length = len(used_letters)
	}

	return biggest_substring_length
}

func removeCharFromS(used_letters string, char_to_remove string) string {
	for char_pos, char := range used_letters {
		if string(char) == char_to_remove {
			return used_letters[char_pos+1:]
		}
	}
	return used_letters
}
