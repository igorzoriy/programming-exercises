//1-1" => 1
//"1-2,1-2" => 1
//"3-2,2-1,1-4,4-4,5-4,4-2,2-1" => 4
//"5-5,5-5,4-4,5-5,5-5,5-5,5-5,5-5,5-5,5-5" => 7


console.log(domino("1-1"))
console.log(domino("1-2,1-2"))
console.log(domino("3-2,2-1,1-4,4-4,5-4,4-2,2-1"))
console.log(domino("5-5,5-5,4-4,5-5,5-5,5-5,5-5,5-5,5-5,5-5"))

function domino(str) {
    var dominos = str.split(',').map((item) => item.split('-'))

    var hits = 0
    var hitsArray = []

    for (var i = 0; i < dominos.length - 1; i++) {
        var left = dominos[i]
        var right = dominos[i + 1]

        if (left[1] === right[0]) {
            hits++
        } else {
            if (hits > 0) {
                hitsArray.push(hits + 1)
                hits = 0
            }
        }
    }
    hitsArray.push(hits + 1)

    return Math.max.apply(null, hitsArray)
}
