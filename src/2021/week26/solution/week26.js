function Button(name) {
    this.name = name;

    return this;
}

function WebPage(name) {
    this.name = name;
    this.clickCounter = 0;

    this.addButton = function(button) {
        button.onClick = () => {
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
