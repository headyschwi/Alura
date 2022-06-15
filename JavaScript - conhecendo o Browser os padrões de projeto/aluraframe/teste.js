let numeros = [3,2,11,20,8,7];

let novosNumeros = numeros.map(numero=>{ numero%2 ? numero*2 : numero
    if(numero % 2){
        return numero*2
    }
    return numero
})


let aprovados = avaliacoes
    .filter(prova => prova.nota >= 7)
    .map(prova => prova.aluno.nome);

console.log(aprovados);















function validaCodigo(codigo) {
    let expressao = /\D{3}-\D{2}-\d{2}/; 
    if(expressao.test(codigo)) {
          alert('Código válido!');
      } else {
          alert('Código inválido');
      }

}

validaCodigo('GWZ-JJ-12'); // válido
validaCodigo('1X1-JJ-12'); // inválido


class Codigo{

    constructor(texto){
        if(!validaCodigo(texto)) throw new Error("código inválido!")
        this._texto = texto
    }

    validaCodigo(texto){
        return /\D{3}-\D{2}-\d{2}/.test(texto)
    }

    get texto(){
        return this._texto
    }
}

numeros.reduce((total,num) => total * num, 1);




    <table>
        <thead>
            <tr>
                <th>Nome</th>
                <th>Endereço</th>
                <th>Salário</th>
            </tr>
        </thead>

        <tbody id="tbody">

            <!-- ELE ESTÁ COM DIFICULDADES AQUI -->
        <tbody>
    <table>

    <script>
        let funcionarios = [
            {
                "nome": "Douglas",
                "endereco" : "Rua da esquina, 123",
                "salario" : "4500"
            },
            {
                "nome": "Felipe",
                "endereco" : "Rua da virada, 456",
                "salario" : "5000"
            },
            {
                "nome": "Silvio",
                "endereco" : "Rua da aresta, 789",
                "salario" : "6000"
            }
        ];

        let funcionariosEmHtml = funcionarios.map((f)=>{
            `
                <tr>
                    <td>${f.nome}</td>
                    <td>${f.endereco}</td>
                    <td>${f.salario}</td>
                </tr>
            `
        }).join()

        let tabela = document.querySelector("#tbody")
        tabela.innerHTML = funcionariosEmHtml
    </script>

let dobro = numeros.map(num => num*2)
let metade = numeros.map(num => num/2)
let raiz = numeros.map(num => num*num)












