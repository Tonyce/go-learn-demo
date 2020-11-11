const http = require('http');

http.createServer((req, res) => {
    res.end("hello Nodejs")
}).listen(8989, () => {
    console.log('test-server start at 8989')
})