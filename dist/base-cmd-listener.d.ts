import { Pool } from "pg";
import { Stan, Message as MessageUtils } from "node-nats-streaming";
import { Message, Categories } from "./events";
export declare abstract class CommandListener {
    abstract category: Categories;
    abstract queueGroupName: string;
    abstract onMessage(message: Message, utils: MessageUtils): void;
    protected db: Pool;
    protected client: Stan;
    protected ackWait: number;
    constructor(client: Stan, db: Pool);
    get cmdStream(): string;
    subscriptionOptions(): import("node-nats-streaming").SubscriptionOptions;
    listen(): void;
    parseMessage(msg: MessageUtils): any;
}
