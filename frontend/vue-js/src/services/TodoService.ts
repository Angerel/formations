import {ListItem} from "@/models/ListItem";
import {FetchService} from "@/services/FetchService";

export class TodoService {
    getTodos(): Promise<ListItem[]> {
        return FetchService.fetch<ListItem[]>("http://localhost:8080/todos");
    }
}