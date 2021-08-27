export class Todo {
    constructor(public name: string, public done: boolean, public id?: string){
        this.id = id;
        this.name = name;
        this.done = done;
    };
}