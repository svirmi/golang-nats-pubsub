import { connect, StringCodec } from "nats";

import express from 'express';
import { createServer } from 'http';
import { WebSocketServer } from 'ws';

  // make ws app
const app = express();
const server = createServer(app);

const wss = new WebSocketServer({ server });
  
server.listen(process.env.PORT || 8999, () => {
  console.log(`Server started on port ${server.address().port}`);
});

const main = async() => {

// create a codec
const sc = StringCodec();

// to create a connection to a nats-server:
const nc = await connect({ servers: "nats://nats:4222", debug: false });

// create a simple subscriber and iterate over messages
// matching the subscription
const sub = nc.subscribe("sensorData");

for await (const m of sub) {

  wss.clients.forEach(function each(client) {
    client.send(sc.decode(m.data));
  });

  // console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
}

console.log("subscription closed");

await nc.drain();

console.log("connection closed");
}

main();

