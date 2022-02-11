const personObj = {
    fullName: 'John Doe',
    country: 'USA',
    city: 'New York',
    street: 'Main st',
    postalCode: 12345,
};

let addressObj = {};

({ fullName, ...addressObj } = personObj);

console.log(addressObj);
