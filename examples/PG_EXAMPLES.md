# PG

## Using enum with go

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

type Status string

const (
    Pending Status = "pending"
    Approved Status = "approved"
    Rejected Status = "rejected"
)

func main() {
    // Connect to the database
    db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create a table with an enum column
    _, err = db.Exec(`
        CREATE TYPE status AS ENUM ('pending', 'approved', 'rejected');
        CREATE TABLE mytable (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            status status NOT NULL
        );
    `)
    if err != nil {
        panic(err)
    }

    // Insert a row with an enum value
    _, err = db.Exec("INSERT INTO mytable (name, status) VALUES ($1, $2)", "John", Approved)
    if err != nil {
        panic(err)
    }

    // Query rows with an enum value
    rows, err := db.Query("SELECT id, name, status FROM mytable")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var status Status
        err = rows.Scan(&id, &name, &status)
        if err != nil {
            panic(err)
        }
        fmt.Printf("id=%d name=%s status=%s\n", id, name, status)
    }
}
```

## Insert with return ID

```go
var userid int
err := db.QueryRow(`INSERT INTO users(name, favorite_fruit, age) VALUES('beatrice', 'starfruit', 93) RETURNING id`).Scan(&userid)
```
