const form = document.getElementById("novoItem") 
const lista = document.getElementById("lista")
const itens = JSON.parse(localStorage.getItem("itens")) || []  // -> [{"nome"}]

itens.forEach( (elemento) => {    
    criaElemento(elemento)
} )

form.addEventListener("submit", (evento) => {  
    evento.preventDefault()

    const nome = evento.target.elements['nome']
    const quantidade = evento.target.elements['quantidade']

    const existe = itens.find(elemento => elemento.nome === nome.value)
    
    const novoItem = {
    "nome": nome.value,
    "quantidade": quantidade.value
    }

    if(existe){
        novoItem.id = existe.id;
        atualizaElemento(novoItem);
        itens[itens.findIndex(item => item.id === existe.id)] = novoItem;
    }
    else{
        novoItem.id = itens[itens.length-1] ? (itens[itens.length-1]).id+1 : 0
        criaElemento(novoItem)
        itens.push(novoItem)
    }

    localStorage.setItem("itens", JSON.stringify(itens))

    nome.value = ""
    quantidade.value = ""
})

function criaElemento(item) {  
    const novoItem = document.createElement('li')
    novoItem.classList.add("item")

    const numeroItem = document.createElement('strong')
    numeroItem.innerHTML = item.quantidade
    numeroItem.dataset.id = item.id
    novoItem.appendChild(numeroItem)

    novoItem.innerHTML += item.nome

    novoItem.appendChild(botaoDeleta(item.id));
    lista.appendChild(novoItem)
    //novoItem.appendChild(botaoDeleta());
}

function atualizaElemento(item){
    document.querySelector("[data-id='"+ item.id+"']").innerHTML = item.quantidade
}

function botaoDeleta(id){
    const elementoBotao = document.createElement('button');
    elementoBotao.innerText = 'X';
    elementoBotao.addEventListener("click", function(){
        deletaElemento(this.parentNode, id);
    })
    return elementoBotao
}

function deletaElemento(tag, id){
    tag.remove();
    itens.splice(itens.findIndex(element => element.id === id), 1)

    localStorage.setItem("itens", JSON.stringify(itens))
}