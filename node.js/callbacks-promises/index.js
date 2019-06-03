const call = require('./src/call');

console.log(call.sync('Sync Function'));

call.withCallback('CallBack Function', call.sync);

call.withPromise('Promise Function')
  .then( name => console.log(name) );
