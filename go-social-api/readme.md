## Learn Go by building Project API 


## Go mod init

```go
go mod init github.com/theisaachome/social
```

## Project Structure

```sh
go get 
```

## Implementing Repository pattern 


### Create UserStore (UserRepo)

- internal/store/users.go


```go
type UserStore struct{
	db *sql.DB
}
func (s *UserStore)Create(ctx context.Context) error{
	return  nil
}
```
### Create PostStore (PostRepository)
- internal/store/posts.go
```go
package store

import (
	"context"
	"database/sql"
)

type PostStore struct{
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context) error{
	return nil
}
```

### Create storage.go file
```go
package store

import (
	"context"
	"database/sql"
)


type Storage struct{
	Posts interface{
		Create (context.Context) error
	}
	Users interface{
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage{
	return Storage{
		Posts: &PostStore{db: db},
		Users: &UserStore{db: db},
	}
}
```
### update api.go to accept database 
```go

type application struct {
	config config
	store store.Storage
}
```
### at main.go file create store instance

```go


func main(){
	cfg :=config{addr: env.GetString("ADDR", ":8080")}

	store :=store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux :=app.mount()
	log.Fatal(app.run(mux))

}
```



### PostgreSQL driver

```go 
go get github.com/lib/pq
```
