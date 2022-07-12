- go run github.com/99designs/gqlgen init
- Substituir o schema.graphqls
- remover arquivos gerados e deixar apenas o schema.graphqls e o gqlgen.yml
- go run github.com/99designs/gqlgen generate
- criar o arquivo server.go:
```GO
package mains

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/viniciussevero/graphql-aluno/graph"
	"github.com/viniciussevero/graphql-aluno/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

```