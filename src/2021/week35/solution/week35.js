(async function() {
    function addWorker(queue) {
        const id = queue.length + 1;
        const worker = () => {
            return new Promise(resolve => {
                setTimeout(() => {
                    console.log(`Worker ${id} is finished`);
                    resolve();
                }, Math.floor(Math.random() * 10) + 1);
            })
        };

        queue.push(worker);
    }

    let queue = [];

    addWorker(queue);
    addWorker(queue);
    addWorker(queue);

    for (const item of queue) {
        await item();
    }
})()
