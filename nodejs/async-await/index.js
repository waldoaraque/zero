const call = require('./src/call');
/*
call.greet('This is not asynchronism');

call.withPromise('This is', 'asynchronism')
  .then( name => console.log(name));
*/

async function callWithPromise() {
  const name = await call.withPromise('This is', 'asynchronism');
  console.log(name);
}

callWithPromise();
