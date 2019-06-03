const greet = require('./src/greet');

greet.emit('clap');
greet.emit('shout', 'HELLO WORLD');
greet.emit('call', 'Waldo', (name) => {
  console.log(`Hello ${name}, this is an event`);
});
