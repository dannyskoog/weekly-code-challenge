/**
 * Code challenge week 35, 2021 (language: Javascript, playground: https://jsfiddle.net/ with “Console” available at the bottom right corner)
 * 
 * Description:
 *     The code below is supposed to queue up workers and eventually put them to work. The workers are supposed to be initialised in the
 *     order they were added to the queue and must finish before the next worker is initialised. But according to the the logging it seems
 *     as the order of when workers finish is not honoured at the moment. Hum… That’s strange!
 * 
 * Questions:
 *     1. Why isn’t the order of when workers finish honoured?
 *     2. How can it be adjusted to honour the order of when workers finish?
 */

 function addWorker(queue) {
    const id = queue.length + 1;
    const worker = () => {
        setTimeout(() => {
            console.log(`Worker ${id} is finished`);
        }, Math.floor(Math.random() * 10) + 1);
    };

    queue.push(worker);
}

let queue = [];

addWorker(queue);
addWorker(queue);
addWorker(queue);

for (const item of queue) {
    item();
}
