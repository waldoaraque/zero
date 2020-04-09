//try/catch => administrar errores

try {
    throw new Error('Algo salió mal')
} catch { //en ES10 no es necesario pasar el error
    console.log('mi propio error')
}