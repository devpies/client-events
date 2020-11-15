import { Stan } from 'node-nats-streaming';
import { Events, Commands } from '.';
declare type Subjects = Events | Commands;
interface Event {
    subject: Subjects;
    data: any;
}
export declare abstract class Publisher<T extends Event> {
    abstract subject: T['subject'];
    private client;
    constructor(client: Stan);
    publish(data: T['data']): Promise<unknown>;
}
export {};
