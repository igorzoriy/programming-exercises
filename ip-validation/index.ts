const isValid = (ip: string): boolean => {
    const sections = ip.split(".")

    return sections.length === 4
        && sections.filter(item => item[0] === "0").length === 0
        && sections.map(item => parseInt(item, 10)).filter(item => item > 0 && item < 255).length === 4
}

console.log("1.2.3", isValid("1.2.3"))
console.log("1.2.3.4.5", isValid("1.2.3.4.5"))
console.log("01.02.03.04", isValid("01.02.03.04"))
console.log("123.045.067.089", isValid("123.045.067.089"))
console.log("1..3.4", isValid("1..3.4"))
console.log("123.456.78.90", isValid("123.456.78.90"))
console.log("1.2.3.a", isValid("1.2.3.a"))

console.log("1.2.3.4", isValid("1.2.3.4"))
console.log("123.45.67.89", isValid("123.45.67.89"))





