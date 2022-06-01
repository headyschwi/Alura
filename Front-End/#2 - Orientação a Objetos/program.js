import {Cliente} from "./Cliente.js";
import {ContaCorrente} from "./ContaCorrente.js";

let cliente = new Cliente()

const conta = new ContaCorrente(cliente, 81);
const conta2 = new ContaCorrente("Rogério", 81);
const conta3 = new ContaCorrente(cliente, 81);
const conta4 = new ContaCorrente(cliente, 81);
const conta5 = new ContaCorrente(cliente, 81);



console.log(conta5.id);
console.log(ContaCorrente.TotalIds);