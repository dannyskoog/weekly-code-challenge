/**
 * Code challenge week 26, 2021 (language: Javascript, playground: https://jsfiddle.net/ with “Console” available at the bottom right corner)
 * 
 * Description:
 *     The code below is supposed to keep track of and output the number of times a particular button has been clicked. But it seems as it’s
 *     not working properly!
 *
 * Questions:
 *     1. Why is it not working properly?
 *     2. How can it be fixed?
 */

function Button(name) {
    this.name = name;

    return this;
}

function WebPage(name) {
    this.name = name;
    this.clickCounter = 0;

    this.addButton = function(button) {
        button.onClick = function() {
            console.log(`Button has been clicked ${this.clickCounter || 0} times`)
        }
        this.button = button;
    }

    this.clickButton = function() {
        this.clickCounter++;
        this.button.onClick();
    }

    return this;
}

var page = new WebPage('Homepage')
var button = new Button('SubmitButton')

page.addButton(button);
page.clickButton();
page.clickButton();
page.clickButton();
