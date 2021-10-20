/**
 * Code challenge week 40, 2021 (language: Javascript, playground: https://jsfiddle.net/ with “Console” available at the bottom right corner)
 * 
 * Description:
 *     That code below has a set of expressions and statements. And it’s your job to guess (without running the code) if each expression
 *     (and effectively the statement as well) is true or false. Tricky huh?!
 */

/*****************************************************************************************************************************************
Expression:                             Statement:                                                                                      #:
*****************************************************************************************************************************************/

console.log([] == ![]);                 // An empty array EQUALS the negation of an empty array                                         1

console.log(true == []);                // The boolean value 'true' EQUALS an empty array                                               2
console.log(true == ![]);               // The boolean value 'true' EQUALS the negation of an empty array                               3

console.log(!!'false' == !!'true');     // A double negated string value of 'false' EQUALS a double negated string value of 'true'      4

console.log(null == 0);                 // null EQUALS number 0                                                                         5
console.log(null > 0);                  // null IS BIGGER THAN number 0                                                                 6
console.log(null >= 0);                 // null IS BIGGER THAN OR EQUALS number 0                                                       7

console.log(1 < 2 < 3);                 // The result of if number 1 is LESS THAN number 2 is LESS THAN number 3                        8
console.log(3 > 2 > 1);                 // The result of if number 3 is LESS THAN number 2 is LESS THAN number 1                        9

console.log(NaN === NaN);               // NaN (Not a Number) EQUALS NaN (Not a Number)                                                 10

console.log(Math.min() < Math.max());   // 'Math.min()' IS LESS THAN 'Math.max()'                                                       11
