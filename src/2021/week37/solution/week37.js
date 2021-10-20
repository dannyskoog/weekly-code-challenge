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

    return Math.round(ticketPrice - ticketPrice * discount)
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
    },
]

const discount = calculateFamilyDiscount(family)
const ticketPrice = calculateFamilyTicketPrice(discount)
console.log(ticketPrice)
