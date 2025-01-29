# Longest Substring Without Repeating Characters

This program calculates the length of the longest substring without repeating characters in a given string.

## Functions

- **hasDuplicatesInString(s string, sub string) bool**: Checks if a substring `sub` contains any duplicate characters within the string `s`.
- **Max(a int, b int) int**: Returns the maximum of two integers `a` and `b`.
- **lengthOfLongestSubstring(s string) int**: Computes the length of the longest substring without repeating characters in the string `s`.

## How It Works

1. **Initial Checks**: If the string length is 0 or 1, it returns the length directly. For strings of length 2, it checks if both characters are the same and returns 1 or 2 accordingly.
2. **Main Logic**: Iterates through the string, building substrings and checking for duplicates using `hasDuplicatesInString`. It keeps track of the maximum length found.
3. **Output**: Returns the length of the longest substring without repeating characters.

## Example

For the input `"abcabcbb"`, the program will return `3`, as the longest substring without repeating characters is `"abc"`.
