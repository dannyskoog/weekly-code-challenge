/**
 * Code challenge week 25, 2021 (language: Javascript, playground: https://jsfiddle.net/ with “Console” available at the bottom right corner)
 *
 * Description:
 *     The code below is supposed to output the sum of numbers starting from 100000 and below (i.e. 100000 + 99999 + 99998 + etc.). The
 *     problem is that it’s broken!
 * 
 * Questions:
 *     1. Why is it broken?
 *     2. How can it be fixed?
 */

const sumBelow = (number, sum = 0) => (
    number === 0 
      ? sum
      : sumBelow(number - 1, sum + number)
  )

console.log(sumBelow(100000))
