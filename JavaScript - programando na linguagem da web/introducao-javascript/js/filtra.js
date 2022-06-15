const campoFiltro = document.querySelector("#filtrar-tabela")

campoFiltro.addEventListener("input", function(){
    const pacientes = document.querySelectorAll(".paciente")
    

    if(campoFiltro.value.length > 0){
        pacientes.forEach(paciente => {
            const nome = paciente.querySelector(".info-nome").textContent

            var expressao = new RegExp(this.value, "i")

            if(!expressao.test(nome)){
                paciente.classList.add("invisivel")
            }
            else{
                paciente.classList.remove("invisivel")
            }
        });
    }
    else{
        pacientes.forEach(paciente => {
            paciente.classList.remove("invisivel")
        });
    }
})