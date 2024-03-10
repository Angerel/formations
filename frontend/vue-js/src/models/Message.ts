import type {User} from "@/models/User";

export type Message = {
    id : string,
    message: string,
    timestamp: string,
    user: User
}