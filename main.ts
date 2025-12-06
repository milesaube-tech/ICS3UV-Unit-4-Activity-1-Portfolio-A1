/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-05
 * @fileoverview This program to count down from 100 in increments of 5
 */

// set variables
let Counter: number = 0
let Output: string = ""

// for loop
for (
  let Counter = 100;
  Counter > 0;
  Counter = Counter - 5
) {
  Output = Output + Counter + ","
}
  Output = Output +  "0"
console.log(Output)