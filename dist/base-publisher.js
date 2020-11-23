"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Publisher = void 0;
class Publisher {
    constructor(client) {
        this.client = client;
    }
    publish(message) {
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
exports.Publisher = Publisher;
