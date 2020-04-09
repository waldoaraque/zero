const sumar = (a, b) => a + b

console.log(sumar.toString())


function Automovil(modelo, color) {
    this.modelo = modelo
    this.color = color
}

Automovil.prototype.toString = function autoString() {
    const datos = `Auto: ${this.modelo}- Color: ${this.color}`
    return datos
}

const auto = new Automovil('Camaro', 'Amarillo')

console.log(auto.toString())
