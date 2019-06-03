module.exports = {
    sync: function(nombre) {
        return `Hello ${nombre}! with Node.js`;
    },
    withCallback: function(nombre, cb) {
        setTimeout( () => {
            console.log(cb(nombre));
        }, 5000);
    },
    withPromise: function(nombre) {
        return new Promise((resolve, reject) => {
            setTimeout( () => {
                resolve(nombre);
            }, 5000);
        });
    }
}
