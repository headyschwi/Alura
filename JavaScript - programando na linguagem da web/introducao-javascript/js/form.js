const botaoAdicionar = document.querySelector("#adicionar-paciente")

botaoAdicionar.addEventListener("click", (e)=>{
    e.preventDefault()

    const form = document.querySelector("#form-adiciona")
    const paciente = obtemPacienteDoFormulario(form)
    
    const erros = validaPaciente(paciente)

    if(erros.length > 0){
        exibeErros(erros)
        return
    }

    adicionaPaciente(paciente)

    const ul = document.querySelector("#mensagens-erros")
    ul.innerHTML = ""

    form.reset()
})

function obtemPacienteDoFormulario(form){
    const paciente = {
        peso: form.peso.value,
        nome: form.nome.value,
        altura: form.altura.value,
        gordura: form.gordura.value,
        imc: calculaImc(form.peso.value, form.altura.value)
    }

    return paciente
}

function validaPaciente(paciente){
    var erros = []

    if(paciente.nome.length == 0) erros.push("O nome não pode ser em branco")
    if(paciente.peso.length == 0) erros.push("O peso não pode ser em branco")
    if(paciente.altura.length == 0) erros.push("A altura não pode ser em branco")
    if(paciente.gordura.length == 0) erros.push("A gordura não pode ser em branco")
    if(!validaPeso(paciente.peso)) erros.push("Peso inválido")
    if(!validaAltura(paciente.altura)) erros.push("Altura inválida")

    return erros

}

function exibeErros(erros){
    const ul = document.querySelector("#mensagens-erros")
    ul.innerHTML = ""

    erros.forEach(erro => {
        const li = document.createElement("li")
        li.textContent = erro
        ul.appendChild(li)
    });
}

function adicionaPaciente(paciente){
    const pacienteTr = montaTr(paciente)
    const tabela = document.querySelector("#tabela-pacientes")
    tabela.appendChild(pacienteTr)
}

function montaTr(paciente){
    const pacienteTr = document.createElement("tr")

    const nomeTd = montaTd(paciente.nome, "info-nome")
    const pesoTd = montaTd(paciente.peso, "info-peso")
    const alturaTd = montaTd(paciente.altura, "info-altura")
    const gorduraTd = montaTd(paciente.gordura, "info-gordura")
    const imcTd = montaTd(paciente.imc, "info-imc")

    pacienteTr.appendChild(nomeTd)
    pacienteTr.appendChild(pesoTd)
    pacienteTr.appendChild(alturaTd)
    pacienteTr.appendChild(gorduraTd)
    pacienteTr.appendChild(imcTd)

    return pacienteTr
}

function montaTd(dado, classe){
    const td = document.createElement("td")
    td.classList.add(classe)
    td.textContent = dado
    return td
}