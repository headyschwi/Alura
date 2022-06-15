const titulo = document.querySelector(".titulo")
titulo.textContent = "Aparecida Nutricionista"


const clientes = document.querySelectorAll(".paciente")
for (let i = 0; i < clientes.length; i++) {
    const cliente = clientes[i]
    
    const pesoCliente = cliente.querySelector(".info-peso")
    const alturaCliente = cliente.querySelector(".info-altura")
    const imcCliente = cliente.querySelector(".info-imc")

    const altura = alturaCliente.textContent
    const peso = pesoCliente.textContent
    
    if(!validaAltura(altura)){
        alturaCliente.textContent = "Altura inválida"
        imcCliente.textContent = "Altura inválida"
        // cliente.style.backgroundColor = "lightcoral"
        cliente.classList.add("paciente-invalido")
    }
    
    if(!validaPeso(peso)){
        pesoCliente.textContent = "Peso inválido"
        imcCliente.textContent = "Peso inválido"
        // cliente.style.backgroundColor = "lightcoral"
        cliente.classList.add("paciente-invalido")
    }
    
    if(validaAltura(altura) && validaPeso(peso)){
        const imc = calculaImc(peso, altura)
        imcCliente.textContent = imc
    }
}

function validaPeso(peso){
    if(peso >= 1000 || peso <0){
        return false
    }
    return true
}

function validaAltura(altura){
    if(altura > 3 || altura <0){
        return false
    }

    return true
}

function calculaImc(peso, altura){
    const imc = peso / (altura*altura)
    return imc.toFixed(2)
}
