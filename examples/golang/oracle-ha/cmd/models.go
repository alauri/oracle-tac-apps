/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command ``delete`` is used to delete one or more records from the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
*/
package cmd

type Conds struct {
    id int
}

type Arguments struct {
    arg int
}
