package models

import (
	d "alura/aplicativo-web/db"
)

type Produto struct {
	Nome, Descricao string
	Valor           float64
	Id, Quantidade  int
}

func RetornaTodosProdutos() []Produto {
	db := d.ConectaBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &valor, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}
func CriaNovoProduto(nome, descricao string, valor float64, quantidade int) {
	db := d.ConectaBancoDeDados()
	scriptDb, err := db.Prepare("insert into produtos(nome, descricao, valor, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	scriptDb.Exec(nome, descricao, valor, quantidade)
	defer db.Close()
}
func DeletaProduto(id string) {
	db := d.ConectaBancoDeDados()

	scriptDb, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	scriptDb.Exec(id)
	defer db.Close()
}
func EditaProduto(id string) Produto {
	db := d.ConectaBancoDeDados()

	produto, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	var p Produto

	for produto.Next() {
		err := produto.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Valor, &p.Quantidade)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
	return p
}
func AtualizaProduto(nome, descricao string, id, quantidade int, valor float64) {
	db := d.ConectaBancoDeDados()

	scriptDb, err := db.Prepare("update produtos set nome=$1, descricao=$2, valor=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	scriptDb.Exec(nome, descricao, valor, quantidade, id)
	defer db.Close()
}
