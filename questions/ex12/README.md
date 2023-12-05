# Что выведет данная программа и почему?
```
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```

Программа выведет 0. Это происходит из-за того, что переменная **n** объявляется в двух областях видимости. Внутри конструкции if n == 2, но печатается **n**, объявленная в начале программы.