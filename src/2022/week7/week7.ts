/**
* Code challenge week 7, 2022 (language/tool: Typescript, playground: https://www.typescriptlang.org/play/)
*
* Description:
*     The code below is supposed to calculate and log the ticket price for a whole year. However, when running the code it yields an error.
*     And after a quick investigation we find out that we have a typo in our 'seasons' array (i.e. we have 'falll' instead of 'fall'). So we
*     correct it and run the code again. And now the ticket price is properly calculated and logged. But we're still anxious about it
*     happening again in the future.
*
* Questions:
*     1. How could the code be adjusted so that the 'season' parameter is ensured to be a valid value at compile-time?
*/

const getTicketPriceForWholeSeason = (season: string): number => {
    switch (season) {
        case 'winter':
            return 100;
        case 'spring':
            return 200;
        case 'summer':
            return 300;
        case 'fall':
            return 200;
    }

    throw new Error('Invalid season');
}

const getTicketPriceForWholeYear = (): number => {
    const seasons = ['winter', 'spring', 'summer', 'falll'];
    const ticketPrice = seasons.reduce<number>((prev, curr) => {
        const seasonTicketPrice = getTicketPriceForWholeSeason(curr);

        return prev + seasonTicketPrice;
    }, 0);
    const discount = 200;

    return ticketPrice - discount;
}

const yearTicketPrice = getTicketPriceForWholeYear();
console.log(yearTicketPrice);
