import { Stan, Message as MessageUtils } from "node-nats-streaming";
import { Message } from ".";

export abstract class Listener<T extends Message> {
  abstract subject: T["subject"];
  abstract queueGroupName: string;
  abstract onMessage(message: T, msg: MessageUtils): void;

  private client: Stan;
  protected ackWait = 5 * 1000; // 5000 milliseconds

  constructor(client: Stan) {
    this.client = client;
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
    const subscription = this.client.subscribe(
      this.subject,
      this.queueGroupName,
      this.subscriptionOptions()
    );

    subscription.on("message", (msg: MessageUtils) => {
      console.log(`Message received: ${this.subject} / ${this.queueGroupName}`);
      const parsedData = this.parseMessage(msg);
      this.onMessage(parsedData, msg);
    });
  }

  parseMessage(msg: MessageUtils) {
    const data = msg.getData();
    return typeof data === "string"
      ? JSON.parse(data)
      : JSON.parse(data.toString("utf-8")); // parse buffer
  }
}
