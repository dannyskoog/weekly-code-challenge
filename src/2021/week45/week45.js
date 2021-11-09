/**
 * Code challenge week 45, 2021 (language: Javascript, playground: https://jsfiddle.net/ with “Console” available at the bottom right corner)
 *
 * Description:
 *     The code below is supposed to cache a DOM element, retrieve it from the cache and then print it. Furthermore the DOM element
 *     is then deleted whereupon it's attempted to be retrieved from the cache again. And the idea is that, this time around, it shouldn't
 *     be in the cache. But once running the code it's obvious that we're lacking something.
 *
 * Questions:
 *     1. How can the code be adjusted so that the cache returns 'undefined' after the DOM element has been removed?
 */

class DomElementCache {
    constructor() {
        this.cache = {};
    }

    set(key, value) {
        this.cache[key] = value;
    }

    get(key) {
        return this.cache[key];
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
