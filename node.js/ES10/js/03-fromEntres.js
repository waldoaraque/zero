const map = new Map([ 
    ['nombre', 'Producto 1'], 
    ['precio', 20] 
])

//fromEntries() => permite crear un objeto de una serie de llaves y valores

const objeto = Object.fromEntries(map)
console.log(objeto)

//se pueden utilizar metodos la entidad Object
delete objeto.precio

console.log(objeto)
