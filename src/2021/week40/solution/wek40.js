/**************************************************************************************************************************************************
Expression:                             Statement:                                                                                      #:  Result:
**************************************************************************************************************************************************/

console.log([] == ![]);                 // An empty array EQUALS the negation of an empty array                                         1   true

console.log(true == []);                // The boolean value 'true' EQUALS an empty array                                               2   false
console.log(true == ![]);               // The boolean value 'true' EQUALS the negation of an empty array                               3   false

console.log(!!'false' == !!'true');     // A double negated string value of 'false' EQUALS a double negated string value of 'true'      4   true

console.log(null == 0);                 // null EQUALS number 0                                                                         5   false
console.log(null > 0);                  // null IS BIGGER THAN number 0                                                                 6   false
console.log(null >= 0);                 // null IS BIGGER THAN OR EQUALS number 0                                                       7   true

console.log(1 < 2 < 3);                 // The result of if number 1 is LESS THAN number 2 is LESS THAN number 3                        8   true
console.log(3 > 2 > 1);                 // The result of if number 3 is LESS THAN number 2 is LESS THAN number 1                        9   false

console.log(NaN === NaN);               // NaN (Not a Number) EQUALS NaN (Not a Number)                                                 10  false

console.log(Math.min() < Math.max());   // 'Math.min()' IS LESS THAN 'Math.max()'                                                       11  false

// Reference: https://github.com/denysdovhan/wtfjs
