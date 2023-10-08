export namespace main {
	
	export class todo {
	    id: number;
	    description: string;
	    due: string;
	    status: string;
	    created: string;
	
	    static createFrom(source: any = {}) {
	        return new todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.due = source["due"];
	        this.status = source["status"];
	        this.created = source["created"];
	    }
	}

}

