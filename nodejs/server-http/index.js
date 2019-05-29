const http = require('http');
const fs = require('fs');

const server = http
  .createServer( (req, res) => {
  	fs.readFile('./resources/index.html', (err, data) => {
  		if (err) {
        console.error(err);
        return;
      }
  		res.end(data)
  	});
  });

server.listen(3000);
