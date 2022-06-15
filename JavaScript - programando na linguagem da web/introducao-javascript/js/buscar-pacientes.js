const botaoBuscar = document.querySelector("#buscar-pacientes")

botaoBuscar.addEventListener("click", function(){
    console.log("Buscando pacientes...")

    const xhr = new XMLHttpRequest()
    xhr.open("GET", "http://api-pacientes.herokuapp.com/pacientes")

    xhr.addEventListener("load", function(){
        
        const erro = document.querySelector("#erro-busca")
        if(xhr.status == 200){
            erro.classList.add("invisivel")
            const resposta = xhr.responseText
            const pacientes = JSON.parse(resposta)
    
            pacientes.forEach(paciente => {
                adicionaPaciente(paciente)
            });
        }
        else{
            erro.classList.remove("invisivel")
            console.log(xhr.status)
            console.log(xhr.responseText)
        }

    })

    xhr.send()
})