import { Stan, Message as MessageUtils } from "node-nats-streaming";
import { Message } from ".";
export declare abstract class Listener<T extends Message> {
    abstract subject: T["subject"];
    abstract queueGroupName: string;
    abstract onMessage(message: T, msg: MessageUtils): void;
    private client;
    protected ackWait: number;
    constructor(client: Stan);
    subscriptionOptions(): import("node-nats-streaming").SubscriptionOptions;
    listen(): void;
    parseMessage(msg: MessageUtils): any;
}
