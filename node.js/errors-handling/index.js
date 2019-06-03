const errors = require('./src/errors');
const handle = require('./src/handling');

//handle.errorFirstCallbackWrong();


  try {
    errors.assertErr();
  } catch (e) {
    console.log(`Error: ${e}`);
  } finally {
    console.log(`Finally!`);
  }
