<script setup lang="ts">
import {ref} from "vue";
import {WebsocketService} from "@/services/WebsocketService";
import type {Message} from "@/models/Message"

let conversation = ref<Message[]>([])
let inputValue:string = "";

const websocketService = await WebsocketService.getInstance();

const uri = window.location.search.substring(1);
const params = new URLSearchParams(uri);
const userName = params.get("sender");
if(userName){
    websocketService.authenticate(userName)
}

websocketService.getHistory()
    .then((history:string) => {
        conversation.value = JSON.parse(history)
    })
websocketService.subscribe({
    next: (msg:any)=>{
        console.log(msg)
        conversation.value.push(msg)
        const convDisplay = document.getElementById("conv-display");
        if(convDisplay) {
            setTimeout(()=>{
                convDisplay.scrollTo({
                    top:convDisplay.scrollHeight,
                    behavior:"smooth"
                })
            }, 150);
        }
    }
})
function sendMsg() {
    void websocketService.postMessage(inputValue)
    inputValue = "";
}
function resize() {
    const element = document.getElementById("textValue")
    if(element) {
        element.style.height = "5px";
        element.style.height = (element.scrollHeight) + "px";
    }
}
</script>

<template>
    <div class="conv-wrapper">
        <div id="conv-display">
            <span v-for="(convo, index) in conversation" :key="index" :class="{mine : convo.user.name===userName}">
                <span class="sender">{{convo.user.name===userName?"Me":convo.user.name}}</span>
                <span class="message">{{convo.message}}</span>
            </span>
        </div>
        <form action="javascript:void(0);">
            <textarea id="textValue" v-model="inputValue" @input="resize()" />
            <input type="submit" id="submitBtn" @click="sendMsg()" />
            <span id="send-icon" @click="sendMsg()"></span>
        </form>
    </div>
</template>

<style scoped>
    .conv-wrapper {
        display: flex;
        flex-direction: column;
        width: 90vw;
        height: 70vh;
        margin-top: 10vh;
        background-color: #414141;
    }
    #conv-display {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        padding: 10px 25px;
        overflow-y: scroll;
        scrollbar-width: thin;
        scrollbar-color: var(--color-background) transparent;
    }

    #conv-display::-webkit-scrollbar {
        width: 9px;
    }

    #conv-display::-webkit-scrollbar-track {
        background: transparent;
    }

    #conv-display::-webkit-scrollbar-thumb {
        background-color: var(--color-background);
    }

    .message {
        display: flex;
        align-items: flex-start;
        flex-direction: column;
        width: max-content;
        max-width: 55vw;
        height: auto;
        border-radius: 10px;
        background-color: #b6b6b6;
        padding: 10px 25px;
        margin: 2px 0;
    }
    .sender {
        font-size: small;
        font-style: italic;
        height: 15px;
        color: var(--color-text);
    }
    .mine {
        display: flex;
        flex-direction: column;
        align-self: flex-end;
    }
    .mine .message {
        background-color: #efab0a;
    }
    .mine .sender {
        display: none;
    }
    form {
        display: flex;
        align-self: center;
        width: calc(100% - 50px);
        height: fit-content;
        margin: 25px;
    }

    #textValue {
        width: calc(100% - 50px);
        min-height: 35px;
        max-height: 100px;
        resize: none;
        border: 0;
        border-radius: 5px;
        overflow-y: scroll;
        scrollbar-width: none;
    }

    #textValue::-webkit-scrollbar {
        width: 0;
    }
    #textValue::-webkit-scrollbar-track {
        background: transparent;
    }
    #textValue::-webkit-scrollbar-thumb {
        background-color: transparent;
    }

    #submitBtn {
        display: none;
    }
    #send-icon {
        display: block;
        margin-left: 15px;
        width: 35px;
        height: 35px;
        -webkit-mask-size: cover;
        -webkit-mask-image: url(src/assets/paper-plane.svg);
        mask-size: cover;
        mask-image: url(src/assets/paper-plane.svg);
        background-color: #efab0a;
        cursor: pointer;
    }

</style>
