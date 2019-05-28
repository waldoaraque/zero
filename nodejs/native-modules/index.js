const fs = require('fs');
const fileops = require('./src/fileops');

let valuesNumbers;

fs.readFile('./resources/number.txt', 'utf8', (err, res) => {
  if (err) throw err; //stop the execution code..
  //console.log(res);
  const numbers = res.split('\n').map(n => Number(n));
  console.log(fileops.incrementValues(numbers));
  valuesNumbers = fileops.incrementValues(numbers);

  fs.writeFile('./resources/numberNew.txt', valuesNumbers.join('\n'), (err, res) => {
    if (err) throw err;
  });
});

let valuesNames

fs.readFile('./resources/name.txt', 'utf8', (err, res) => {
  if (err) throw err;
  const names = res.split('\n');
  console.log(fileops.callNames(names));
  valuesNames = fileops.callNames(names);

  fs.writeFile('./resources/nameNew.txt', valuesNames.join('\n'), (err, res) => {
    if (err) throw err;
  });

});
