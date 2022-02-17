enum Season {
    Winter = 'winter',
    Spring = 'spring',
    Summer = 'summer',
    Fall = 'fall',
}

const getTicketPriceForWholeSeason = (season: Season): number => {
    switch (season) {
        case Season.Winter:
            return 100;
        case Season.Spring:
            return 200;
        case Season.Summer:
            return 300;
        case Season.Fall:
            return 200;
    }

    throw new Error('Invalid season')
}

const getTicketPriceForWholeYear = (): number => {
    const seasons = Object.values(Season)
    const ticketPrice = seasons.reduce<number>((prev, curr) => {
        const seasonTicketPrice = getTicketPriceForWholeSeason(curr)

        return prev + seasonTicketPrice;
    }, 0)
    const discount = 200

    return ticketPrice - discount;
}

const yearTicketPrice = getTicketPriceForWholeYear()
console.log(yearTicketPrice)
