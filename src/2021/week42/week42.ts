/**
 * Code challenge week 42, 2021 (language: Typescript, playground: https://www.typescriptlang.org/play)
 * 
 * Description:
 *     The code below is supposed to handle vehicles. It prints information but also initiates the fuelling process. The code compiles
 *     perfectly but once ran it yields an error. It seems as we’re attempting to invoke a function (called fuel) that doesn’t exist on our
 *     bicycle object. Hmm… That doesn’t seem right?!
 * 
 * Questions:
 *     1. How can we adjust the code so that we no longer yield an error at runtime?
 *     2. How can we achieve #1 but also get compile time type checking of the fuel function?
 */

interface Car {
    brand: string;
    color: string;
    fuel: () => void;
}

interface Bicycle {
    brand: string;
    color: string;
}

const handleVehicles = (...vehicles: any[]) => {
    for (const vehicle of vehicles) {
        console.log(`Vehicle brand is: ${vehicle.brand}`);
        console.log(`Vehicle color is: ${vehicle.color}`);
    
        vehicle.fuel();
    }
};

const car: Car = { brand: 'Toyota', color: 'Red', fuel: () => console.log('Fueling...') };
const bicycle: Bicycle = { brand: 'Scott', color: 'Blue' };

handleVehicles(car, bicycle);
