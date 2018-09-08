export default class Vertex {
    private value: number
    private children: Vertex[] = []

    constructor(value: number) {
        this.value = value
    }

    public getValue() {
        return this.value
    }

    public addChild(child: Vertex) {
        this.children.push(child)
    }

    public getChildren(): Vertex[] {
        return this.children
    }
}