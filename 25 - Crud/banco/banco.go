package banco

import ("database/sql"

		_"github.com/go-sql-driver/mysql" //druver de conexão com o mysql

)

//conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8"

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

