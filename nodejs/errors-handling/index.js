const errors = require('./src/errors');

try {
  errors.standardErr.range();
} catch (e) {
  console.log(`Error: ${e}`);
} finally {
  console.log(`Finally!`);
}
