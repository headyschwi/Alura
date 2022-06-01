export class Cliente{
    Nome;
    _CPF;

    get cpf(){
      return this._CPF  
    }

    constructor(nome, cpf){
        this.Nome = nome;
        this._CPF = cpf
    }
}
