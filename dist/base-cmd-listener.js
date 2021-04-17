"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommandListener = void 0;
class CommandListener {
    constructor(client, db) {
        this.ackWait = 5 * 1000; // 5000 ms
        this.client = client;
        this.db = db;
    }
    get cmdStream() {
        return `${this.category}.command`;
    }
    subscriptionOptions() {
        return this.client
            .subscriptionOptions()
            .setDeliverAllAvailable()
            .setManualAckMode(true)
            .setAckWait(this.ackWait)
            .setDurableName(this.queueGroupName);
    }
    listen() {
        const subscription = this.client.subscribe(this.cmdStream, this.queueGroupName, this.subscriptionOptions());
        subscription.on("message", (msg) => {
            console.log(`Message received: ${this.cmdStream} / ${this.queueGroupName}`);
            const parsedData = this.parseMessage(msg);
            this.onMessage(parsedData, msg);
        });
    }
    parseMessage(msg) {
        const data = msg.getData();
        return typeof data === "string"
            ? JSON.parse(data)
            : JSON.parse(data.toString("utf-8")); // parse buffer
    }
}
exports.CommandListener = CommandListener;
