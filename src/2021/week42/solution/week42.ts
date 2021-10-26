 interface Car {
    brand: string;
    color: string;
    fuel: () => void;
}

interface Bicycle {
    brand: string;
    color: string;
}

const isCar = (obj: object): obj is Car => {
    const car = obj as Car

    if (car == null) {
        return false
    }

    return typeof car.brand === 'string'
        && typeof car.color === 'string'
        && typeof car.fuel === 'function'
}

const handleVehicles = (...vehicles: (Car | Bicycle)[]) => {
    for (const vehicle of vehicles) {
        console.log(`Vehicle brand is: ${vehicle.brand}`);
        console.log(`Vehicle color is: ${vehicle.color}`);
    
        if (isCar(vehicle)) {
            vehicle.fuel();
        }
    }
};

const car: Car = { brand: 'Toyota', color: 'Red', fuel: () => console.log('Fueling...') };
const bicycle: Bicycle = { brand: 'Scott', color: 'Blue' };

handleVehicles(car, bicycle);
