import Vertex from './Vertex'

function fillPrintMatrix(vertex: Vertex, matrix: string[][], depth: number) {
    const line: string[] = []
    for (let i = 0; i <= depth; i++) {
        line[i] = i < depth ? '-' : vertex.getValue().toString()
    }
    matrix.push(line)

    const nextDepth = ++depth
    vertex.getChildren().forEach(item => {
        fillPrintMatrix(item, matrix, nextDepth)
    })
}

export default class Tree {
    static readFromBuffer(buffer: Buffer): Tree {
        const items = buffer.toString().split('\n')

        let root: Vertex
        const vertexList: Vertex[] = []

        items.forEach(item => {
            if (item.length === 0) {
                return
            }

            const data = item.split(' ')
            const index = parseInt(data[0], 10)
            const parentIndex = parseInt(data[1], 10)
            const value = parseInt(data[2], 10)

            const vertex = new Vertex(value)
            vertexList[index] = vertex

            if (parentIndex < 0) {
                if (!root) {
                    root = vertex
                } else {
                    throw new Error(`Invalid buffer format. Found two or more roots.`)
                }
            } else if (!vertexList[parentIndex]) {
                throw new Error(`Invalid parent index ${parentIndex}`)
            } else {
                vertexList[parentIndex].addChild(vertex)
            }
        })

        if (!root) {
            throw new Error(`Invalid buffer format. Root not found.`)
        }
        return  new Tree(root)
    }

    private root: Vertex

    constructor(root: Vertex) {
        this.root = root
    }

    public print() {
        const matrix: string[][] = []
        fillPrintMatrix(this.root, matrix, 0)
        console.log(matrix)
    }
}