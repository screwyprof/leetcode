// Valid Mountain Array
// See https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/
//
// Given an array A of integers, return true if and only if it is a valid mountain array.
// Recall that A is a mountain array if and only if:
// 1 . A.length >= 3
// 2. There exists some i with 0 < i < A.length - 1 such that:
//   * A[0] < A[1] < ... A[i-1] < A[i]
//   * A[i] > A[i+1] > ... > A[A.length - 1]
//
// Example 1:
//   Input: [2,1]
//   Output: false
//
// Example 2:
//   Input: [3,5,5]
//   Output: false
//
// Example 3:
//   Input: [0,3,2,1]
//   Output: true
//
// Hint:
//   It's very easy to keep track of a monotonically increasing or decreasing ordering of elements.
//   You just need to be able to determine the start of the valley in the mountain and from that point onwards,
//   it should be a valley i.e. no mini-hills after that. Use this information in regards to the values in the array
//   and you will be able to come up with a straightforward solution.
package mntarr
