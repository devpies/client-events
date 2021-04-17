import { Pool } from "pg";
import { Stan, Message as MessageUtils } from "node-nats-streaming";
import { Message, Categories } from "./events";

export abstract class CommandListener {
  abstract category: Categories;
  abstract queueGroupName: string;
  abstract onMessage(message: Message, utils: MessageUtils): void;

  protected db: Pool;
  protected client: Stan;
  protected ackWait = 5 * 1000; // 5000 ms

  constructor(client: Stan, db: Pool) {
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
    const subscription = this.client.subscribe(
      this.cmdStream,
      this.queueGroupName,
      this.subscriptionOptions()
    );

    subscription.on("message", (msg: MessageUtils) => {
      console.log(
        `Message received: ${this.cmdStream} / ${this.queueGroupName}`
      );
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
