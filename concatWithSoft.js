function concatWithSort (first, second) {
    let [fi, si] = [0, 0]
    let result = []

    while (true) {
        if (fi >= first.length) {
            return result.concat(second.slice(si))
        }
        if (si >= second.length) {
            return result.concat(first.slice(fi))
        }

        if (first[fi] < second[si]) {
            result.push(first[fi])
            fi++
        } else if (first[fi] > second[si]) {
            result.push(second[si])
            si++
        } else {
            result.push(first[fi], second[si])
            fi++
            si++
        }
    }
}


console.log(concatWithSort([], []))
console.log(concatWithSort([1, 2, 3], []))
console.log(concatWithSort([], [1, 2, 3]))
console.log(concatWithSort([1, 2, 3], [1, 2, 3]))
console.log(concatWithSort([1, 2, 5], [3, 4]))
console.log(concatWithSort([1, 2, 5], [3, 7, 8, 9]))
console.log(concatWithSort([10, 11, 15, 22], [3, 4, 20]))
