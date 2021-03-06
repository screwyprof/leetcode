// Check If N and Its Double Exist
// See https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
//
// Given an array arr of integers, check if there exists two integers N and M
// such that N is the double of M ( i.e. N = 2 * M).
// More formally check if there exists two indices i and j such that:
//   1. i != j
//   2. 0 <= i, j < arr.length
//   3. arr[i] == 2 * arr[j]
//
// Example 1:
//   Input: arr = [10,2,5,3]
//   Output: true
//   Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
//
// Example 2:
//   Input: arr = [7,1,14,11]
//   Output: true
//   Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
//
// Example 3:
//   Input: arr = [3,1,7,11]
//   Output: false
//   Explanation: In this case does not exist N and M, such that N = 2 * M.
//
// Hint 1:
//   Loop from i = 0 to arr.length, maintaining in a hashTable the array elements from [0, i - 1].
//
// Hint 2:
//   On each step of the loop check if we have seen the element 2 * arr[i] or arr[i] / 2 was seen if arr[i] % 2 == 0.
package chkdup
