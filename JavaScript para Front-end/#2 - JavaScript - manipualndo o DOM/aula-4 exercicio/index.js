const botao = document.querySelector("[botao]")
const lista = document.querySelector(".lista")

var amostra = true

botao.addEventListener("click", (evento)=>{
    if(amostra){
        lista.style.visibility = "hidden"
        amostra = false
    }
    else{
        lista.style.visibility = "visible"
        amostra = true
    }
})