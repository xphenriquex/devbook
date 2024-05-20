package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick from
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?`,
		publicacaoId,
	)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p
	inner join usuarios u on u.id = p.autor_id
	inner join seguidores s on p.autor_id = s.usuario_id
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc
	`,
		usuarioId, usuarioId,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoId uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set titulo = ?, conteudo = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from publicacoes where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) BuscarPorUsusario(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select distinct p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		where u.id = ?
		`,
		usuarioId,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Curtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set curtidas = curtidas + 1 where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Descurtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes set curtidas = 
		CASE WHEN curtidas > 0 THEN curtidas - 1 
		ELSE 0 END 
		where id = ?`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
