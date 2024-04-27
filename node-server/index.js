import { connect, StringCodec } from "nats";

const main = async() => {
// create a codec
const sc = StringCodec();

// to create a connection to a nats-server:
const nc = await connect({ servers: "nats://nats:4222", debug: false });

// create a simple subscriber and iterate over messages
// matching the subscription
const sub = nc.subscribe("sensorData");

for await (const m of sub) {
  console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
}

console.log("subscription closed");

await nc.drain();
console.log("connection closed");
}

main();