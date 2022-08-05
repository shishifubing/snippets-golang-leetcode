/* https://leetcode.com/problems/guess-number-higher-or-lower/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Guess Number Higher or Lower.
Memory Usage: 1.9 MB, less than 80.74% of Go online submissions for Guess Number Higher or Lower.

We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

    -1: Your guess is higher than the number I picked (i.e. num > pick).
    1: Your guess is lower than the number I picked (i.e. num < pick).
    0: your guess is equal to the number I picked (i.e. num == pick).

Return the number that I picked.



Example 1:

Input: n = 10, pick = 6
Output: 6

Example 2:

Input: n = 1, pick = 1
Output: 1

Example 3:

Input: n = 2, pick = 1
Output: 1



Constraints:

    1 <= n <= 231 - 1
    1 <= pick <= n

*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
package main

func guess(input int) int {
	return 0
}

func guessNumber(n int) int {
	// checking edge cases
	if guess(1) == 0 {
		return 1
	} else if guess(n) == 0 {
		return n
	}
	left, right := 1, n
	for right >= left {
		// overflow protection
		number := left + (right-left)/2
		switch guess(number) {
		case 0:
			// found the target
			return number
		case -1:
			// the number is bigger -> the target is to the left -> discard right
			right = number - 1
		case 1:
			// the number is smaller -> the target is to the right -> discard left
			left = number + 1
		}
	}
	// ide shows an error, this return is unreachable in this issue
	// since we are using continuous numbers
	return 0
}
