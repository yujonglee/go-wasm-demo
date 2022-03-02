// Arrow function now works
function add(a, b) {
    return a + b;
}
function hello() {
    return "Hello WASM!";
}

// let and const is not globally scoped.
var env = "ENV";
globalThis.env2 = "ENV2";
