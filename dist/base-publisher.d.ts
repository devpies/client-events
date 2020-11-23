import { Stan } from "node-nats-streaming";
import { Message } from ".";
export declare abstract class Publisher<T extends Message> {
    abstract subject: T["subject"];
    private client;
    constructor(client: Stan);
    publish(message: T): Promise<unknown>;
}
