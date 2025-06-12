package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // O driver para SQLite
)

func main() {
	// Conecta ao banco de dados SQLite. O arquivo será criado se não existir.
	db, err := sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// (Opcional) Cria a tabela de itens se ela não existir
	createTableSQL := `CREATE TABLE IF NOT EXISTS items (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"quantity" INTEGER
	  );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	// Inicia o Gin
	r := gin.Default()

	// Diz ao Gin onde encontrar os arquivos HTML
	r.LoadHTMLGlob("templates/*")

	// Rota principal que renderiza a página
	r.GET("/", func(c *gin.Context) {
		// Por enquanto, apenas renderizamos a página sem dados do banco
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Controle de Estoque Profissional",
		})
	})

	log.Println("Servidor rodando em http://localhost:8080")
	r.Run(":8080")
}