import { readFileSync } from 'fs'
import { resolve } from 'path'
import Tree from './Tree'

const tree = Tree.readFromBuffer(
    readFileSync(resolve(__dirname, 'tree.txt'))
)
tree.print()