const WebSocket = require('ws');
 
const ws = new WebSocket('ws://127.0.0.1:8080/ws');

function heartbeat() {
    clearTimeout(this.pingTimeout);
   
    // Use `WebSocket#terminate()`, which immediately destroys the connection,
    // instead of `WebSocket#close()`, which waits for the close timer.
    // Delay should be equal to the interval at which your server
    // sends out pings plus a conservative assumption of the latency.
    this.pingTimeout = setTimeout(() => {
      this.terminate();
    }, 30000 + 1000);
}

// ws.on('open', heartbeat);
ws.on('ping', () => {
    ws.pong('pong')
});
ws.on('close', function clear() {
//   clearTimeout(this.pingTimeout);
    console.log('close')
    
});
 
ws.on('open', function open() {
  ws.send('something');
});
 
ws.on('message', function incoming(data) {
  console.log({data});
});