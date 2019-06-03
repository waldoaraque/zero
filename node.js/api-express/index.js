const express = require("express");
const app = express();
const { PORT } = require("./config");

require("./routes/api")(app);
require("./routes/views")(app);

function init(){
    console.log("Stargint express instance...");
    app.listen(PORT, ()=>{
        console.log("Express server is running.");
    });
}

init();
