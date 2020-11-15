import { Stan } from 'node-nats-streaming';
import { Events, Commands } from '.';

type Subjects = Events | Commands

interface Event {
    subject: Subjects;
    data: any;
}

export abstract class Publisher<T extends Event> {
    abstract subject: T['subject'];
    private client: Stan;

    constructor(client: Stan) {
        this.client = client
    }

    publish(data: T['data']) {
        return new Promise((resolve,reject)=> {
            this.client.publish(this.subject, JSON.stringify(data), (err)=> {
               if (err) {
                   return reject(err)
               }
               resolve();
                console.log('Event Published', this.subject)
            })
        })
       
    }
}