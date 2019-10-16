package main

import (
   "server/route"
)

var a string

func main() {
   a = "G"
   print(route.RouterName)
   f1()
}

func f1() {
   a := "O"
   print(a)
   f2()
}

func f2() {
   print(a)
}