import { Stan, Message as MessageUtils } from "node-nats-streaming";
import { Message } from ".";
export declare abstract class Listener<T extends Message> {
    abstract type: T["type"];
    abstract queueGroupName: string;
    abstract onMessage(message: T, msg: MessageUtils): void;
    protected client: Stan;
    protected ackWait: number;
    constructor(client: Stan);
    subscriptionOptions(): import("node-nats-streaming").SubscriptionOptions;
    listen(): void;
    parseMessage(msg: MessageUtils): any;
}
