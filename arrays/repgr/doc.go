// Replace Elements with Greatest Element on Right Side
// See https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/
//
// Given an array arr, replace every element in that array with the greatest element among the elements to its right,
// and replace the last element with -1.
//
// Example 1:
//   Input: arr = [17,18,5,4,6,1]
//   Output: [18,6,6,6,1,-1]
//
// Constraints:
//   1 <= arr.length <= 10^4
//   1 <= arr[i] <= 10^5
//
// Hint 1:
//   Loop through the array starting from the end.
// Hint 2:
//   Keep the maximum value seen so far.
package repgr
