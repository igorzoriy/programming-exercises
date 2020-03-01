function chain (...classes) {
    let proto = {};
    for (let cls of classes) {
        for (let method of Object.getOwnPropertyNames(cls.prototype)) {
            if (method === "constructor") {
                continue
            }
            proto[method] = cls.prototype[method]
        }
    }
    function newClass() {}
    newClass.prototype = proto
    return newClass
}

function Shape5() {}
Shape5.prototype.shape = function () {console.log("shape5") }

function Polygon5() {}
Polygon5.prototype.polygon = function () { console.log("polygon5") }

function Square5() {}
Square5.prototype.square = function () { console.log("square5") }

class Shape6 {
    shape() { console.log("shape6") }
}

class Polygon6 {
    polygon() { console.log("polygon6") }
}

class Square6 {
    square() { console.log("square6") }
}

var ExtendedSquare5 = chain(Square5, Polygon5, Shape5);
var ExtendedSquare6 = chain(Square6, Polygon6, Shape6);

var es5 = new ExtendedSquare5();
es5.shape();
es5.polygon();
es5.square();

var es6 = new ExtendedSquare6();
es6.shape();
es6.polygon();
es6.square();