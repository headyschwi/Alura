import { Cliente } from "./Cliente.js";

export class ContaCorrente{
    static TotalIds = 0;
    agencia;
    _id = 0;
    _saldo = 0;
    _cliente;
    
    get id(){
        return this._id;
    }
    get saldo(){
        return this._saldo;
    }
    set cliente(value){
        if(value instanceof Cliente){
            this._cliente = value;
        }
    }
    get cliente(){
        return this._cliente
    }

    constructor(cliente, agencia){
        this.cliente = cliente;
        this.agencia = agencia;
        this._id = ContaCorrente.TotalIds + 1;
        ContaCorrente.TotalIds++;
    }

    sacar(valor){
        if(valor <= 0 || valor >= this._saldo) return;
        this._saldo -= valor;
        return valor;
    }
    
    depositar(valor){
        if(valor <= 0) return false;
        this._saldo += valor;
        return true;
    }

    transferir(valor, contaDestino){
        let validator = this.sacar(valor);
        if(!validator) return false;
        contaDestino.depositar(valor);
        return true;
    }
}