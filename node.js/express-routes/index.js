const express = require("express");
const app = express();

/*
app.get("/", (req, res) => {
    res.send("Hola, estoy en la ruta '/'!");
});
*/

require('./routes/views')(app);
require('./routes/special')(app);

console.log("Iniciando Express.js");
app.listen(3000, () => {
    console.log("Express ha iniciado correctamente!");
});