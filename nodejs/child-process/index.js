const childProcess = require('child_process');

function execCommand(command) {
  childProcess.exec(command, (err, stdout, stderr) => {
    if (err) {
      console.log(`Error ${err}`);
      return;
    }
    if (stdout) {
      console.log(`Standard out: \n ${stdout}`);
    }
    if (stderr) {
      console.log(`Standard out: \n ${stderr}`);
    }
  });
}

//execCommand('ls');
execCommand('node ./src/script.js --base=5');
