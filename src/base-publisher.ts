import { Stan } from "node-nats-streaming";
import { Message } from ".";

export abstract class Publisher<T extends Message> {
  abstract subject: T["subject"];
  private client: Stan;

  constructor(client: Stan) {
    this.client = client;
  }

  publish(message: T) {
    return new Promise((resolve, reject) => {
      this.client.publish(this.subject, JSON.stringify(message), (err) => {
        if (err) {
          return reject(err);
        }
        resolve();
        console.log("Message Published", this.subject);
      });
    });
  }
}
