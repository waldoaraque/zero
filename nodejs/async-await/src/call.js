module.exports = {
    greet: function(nombreCompleto) {
        console.log(`Hello ${nombreCompleto}!`);
    },
    withPromise: function(nombre, apellido) {
        console.log("Calling the function with promise...");
        return new Promise((resolve, reject)=>{
            console.log("Executing...");
            setTimeout(()=>{
                console.log("Solving!");
                resolve(`${nombre} ${apellido}`);
            }, 5000);
        })
    }
}
