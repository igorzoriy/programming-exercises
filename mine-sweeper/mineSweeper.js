var fs = require('fs')
var path = require('path')

function calcMines (file) {
    var board = fs.readFileSync(file).toString().split('\n')
    if (!board[board.length - 1].length) {
        board.pop()
    }
    board = board.map((item) => item.split(' '))

    for (var x = 0; x < board.length; x++) {
        for (var y = 0; y < board[x].length; y++) {
            if (board[x][y] !== 'X') {
                continue
            }

            var moore = [
                [x - 1, y - 1],
                [x - 1, y],
                [x - 1, y + 1],
                [x,     y - 1],
                [x,     y + 1],
                [x + 1, y - 1],
                [x + 1, y],
                [x + 1, y + 1],
            ]

            for (var k = 0; k < moore.length; k++) {
                var kx = moore[k][0]
                var ky = moore[k][1]
                if (!board[kx] || !board[kx][ky]) {
                    continue
                }
                if (board[kx][ky] === "X") {
                    continue
                }

                if (board[kx][ky] === "O") {
                    board[kx][ky] = 0
                }
                board[kx][ky]++
            }
        }
    }

    return board.map(item => item.join(' ')).join('\n')
}

console.log(calcMines(path.join(__dirname, './mine-sweeper.txt')))
