Learning go from github.com/ardanlabs/gotraining.

First Program: variables declaration and initialization.


    1. Variables declared will have their zero value.

        `var a int` -> **int 0**.

    2. If variable is declared and initailized, use short variable declaration operator.

        `a := "Hello World!"` --> **string "Hello World!"**.

    run:
        `go run language_mechanics/variables/main.go`

Second Program: Struct type declaration, initialization and anonymous declaration
    1. Stuct initialized without values will have their zero values
        `type user struct = { 
            name string
            age int
            email string
        }
        u1 := user{}` --> **{0}**
    2. Struct declared using var is anonymous
       1. 
            `user := {
                name string
                age int 
                email string
            }{
                name: "madhu",
                age: 10,
                email: "madhureddy88@gmail.com"
            }` --> **{"madhu", 10, "madhureddy88@gmail.com"}
        2.  
            `var user struct {
                name string
                age int
                email string
            }
            user.name = "madhu"
            user.age = 10
            user.email = "madhureddy88@gmail.com"
            `--> **{"madhu", 10, "madhureddy88@gmail.com"}


    run:
        `go run language_mechanics/struct_type/main.go`