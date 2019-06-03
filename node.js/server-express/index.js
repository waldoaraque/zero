const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.send(`I'm in home page.`);
});

app.get('/app', (req, res) => {
  res.send(`I'm in /app page.`);
});

app.get('*', (req, res) => {
  res.send(`I don't know where is here.`);
})

app.listen(port, () => {
  console.log(`Server working at port: ${port}`);
});
