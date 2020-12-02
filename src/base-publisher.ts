import { Stan } from "node-nats-streaming";
import { Message } from ".";

export abstract class Publisher<T extends Message> {
  abstract subject: T["subject"];
  public streamName: string;
  private client: Stan;

  constructor(client: Stan, context: string) {
    this.client = client;
    this.streamName = this.format(context);
  }

  format(context: string) {
    return this.subject + context;
  }

  publish(message: T) {
    return new Promise((resolve, reject) => {
      this.client.publish(this.streamName, JSON.stringify(message), (err) => {
        if (err) {
          return reject(err);
        }
        resolve();
        console.log("Message Published", this.subject);
      });
    });
  }
}
