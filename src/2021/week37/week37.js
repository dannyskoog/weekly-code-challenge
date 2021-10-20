/**
 * Code challenge week 37, 2021 (language: Javascript, playground: https://jsfiddle.net/ with “Console” available at the bottom right corner)
 * 
 * Description:
 *     The code below is supposed to calculate the ticket price for a family. And depending on the family members’ ages they get an
 *     accumulated discount. But this family is confused! When they walked up to the cashier to pay they were expecting to hear a ticket
 *     price of 40. But instead they got to hear something else. Hum… Maybe he should’ve taken that extra math class anyway, the father
 *     mumbles silently.
 * 
 * Questions:
 *     1. Why didn’t the family get to hear the ticket price they expected?
 *     2. How can the code be adjusted to result in the expected ticket price of 40 for the family?
 */

 function calculateFamilyDiscount(family) {
    const discountByAgeMap = {
        '10': 0.3,
        '38': 0.2,
        '41': 0.1
    }

    return family.reduce((prev, curr) => {
        return prev + (discountByAgeMap[curr.age] || 0)
    }, 0)
}

function calculateFamilyTicketPrice(discount) {
    const ticketPrice = 100

    return ticketPrice - ticketPrice * discount
}

const family = [
    {
        name: 'Adam',
        age: 41,
    },
    {
        name: 'Eva',
        age: 38,
    },
    {
        name: 'Lilian',
        age: 10,
    }
]

const discount = calculateFamilyDiscount(family)
const ticketPrice = calculateFamilyTicketPrice(discount)
console.log(ticketPrice)
