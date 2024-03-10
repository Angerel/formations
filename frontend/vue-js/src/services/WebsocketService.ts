import {environment} from "../../environment/environment.dev";
import { v4 as uuid } from 'uuid';

type Observer = {
    next: (res:any)=>void,
    error?:(err:any)=>void
}

export class WebsocketService {
    static instance: WebsocketService;

    private ws!: WebSocket;
    private observers: {[key:string]:Observer} = {};
    private isReady = false;

    private constructor() {
        this.ws = new WebSocket(environment.wsURL);

        this.ws.onmessage = (msg) => {
            const toReturn = JSON.parse(msg.data)
            Object.keys(this.observers).forEach((key)=>{
                this.observers[key].next(toReturn)
            })
        }

        this.ws.onopen = () => {
            this.isReady = true;
        }
    }

    static async getInstance() : Promise<WebsocketService> {
        if(!this.instance){
            this.instance = new WebsocketService();
        }
        return new Promise<WebsocketService>((resolve) => {
            const check = () => {
                if(this.instance.isReady) {
                    resolve(this.instance)
                }else{
                    setTimeout(check, 250);
                }
            }
            check()
        })
    }

    subscribe(observer: Observer) {
        const key = uuid()
        this.observers[key] = observer;
        return () => {
            delete this.observers[key]
        }
    }

    async sendAndReceive(content: string) {
        if(this.ws){
            this.ws.send(content)
            let subscriptionCleanup = () => {};
            const toReturn = await new Promise<any>((resolve, reject)=>{
                subscriptionCleanup = this.subscribe({
                    next: (res:string) => {
                        resolve(res)
                    },
                    error: (err:string) => {
                        reject(err);
                    }
                })
            })
            subscriptionCleanup();
            return toReturn
        }
        return Promise.resolve("")
    }

    authenticate(name: string){
        return this.sendAndReceive(JSON.stringify({
            action: "authenticate",
            options: JSON.stringify({name: name})
        }))
    }

    getHistory(){
        return this.sendAndReceive(JSON.stringify({
            action:"getHistory",
            options: JSON.stringify({})
        }))
    }

    postMessage(message: string){
        return this.sendAndReceive(JSON.stringify({
            action:"postMessage",
            options: JSON.stringify({message: message})
        }))
    }
}