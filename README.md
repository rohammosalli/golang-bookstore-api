# Simplae go-bookstore-api

### How run the application 

```bash
cd cmd/main/
go build main.go
go run main.go
```

### API Doc

##### GET * Get All
http://localhost:8080/book/



##### GET * GetByID
http://localhost:8080/book/2


##### POST * Create Book
http://localhost:8080/book/
```json
{
  "Name": "Mars",
  "Author": "Roham",
  "Publication": "Royan"
}
```

##### PUT * Update 
http://localhost:8080/book/2

##### Delete * Delete
http://localhost:8080/book/2


##### GET * GetByName

http://localhost:8080/book/name/Mars