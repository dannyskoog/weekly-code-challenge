class DomElementCache {
    constructor() {
        this.cache = {};
    }

    set(key, value) {
        this.cache[key] = new WeakRef(value);
    }

    get(key) {
        return this.cache[key]?.deref();
    }
}

// Simulate existing DOM element
let domElement = { content: 'Foo bar' };

const cache = new DomElementCache();
cache.set('my-element', domElement);

const cachedDomElement = cache.get('my-element');
console.log(cachedDomElement);

// Simulate DOM element removal
domElement = undefined;

setTimeout(() => {
    const cachedDomElement = cache.get('my-element');
    console.log(cachedDomElement);
}, 10000);

// Note: The above solution is only confirmed to work with Node.js (>= 14.15.4)
// due to certain garbage collection behaviour being of essence.
