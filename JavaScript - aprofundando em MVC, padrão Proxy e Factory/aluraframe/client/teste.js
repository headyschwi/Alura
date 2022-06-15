class Relogio {

    constructor() {
        this._segundos = 0;

        setInterval(() => {
            console.log(++this._segundos);
          }, 1000);
    }
}

var relogio = new Relogio();


let funcionario = {
    email: 'abc@abc.com'
};


new Proxy(funcionario, {
    get(target, prop, receiver){
        console.log("Armadilha aqui");
        return `**${target[prop]}**`
    }
})

class Funcionario {

    constructor(email) {
        this._email = email;
    }

    get email() {
        return this._email;
    }

    set email(email) {
        this._email = email;
    }
}

new Proxy(new Funcionario("aaa@gmail.com"), {
    get(target, prop, receiver){

        console.log("Armadilha aqui")
        return Reflect.get(target, prop, receiver)
    },

    set(target, prop, value, receiver){
        console.log(`${prop} | Antigo valor: ${target[prop]}, novo valor: ${value}`)
        return Reflect.set(target, prop, value, receiver)
    }
})




let funcionario2 = {email: 'abc@abc.com'};
new Proxy(funcionario2, {
    set(target, prop, value, receiver){
        console.log(`Valor da antigo: ${target[prop]}, valor novo: ${value}`)
        return Reflect.set(target, prop, value, receiver)
    }
})













let dadosServidor = [
    [
        [new Date(), 1, 100],
        [new Date(), 2, 100]
    ],
    [
        [new Date(), 1, 150],
        [new Date(), 2, 300]
    ],
    [
        [new Date(), 3, 50],
        [new Date(), 1, 100]
    ]        
];


let listaDeNegociacoes = dadosServidor.reduce((novoArray, array) => novoArray.concat(array), [])
.map(dado => new Negociacao(new Date(dado.data), dado.quantidade, dado.valor ));






