# Что выведет данная программа и почему?

```
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```

Программа выведет[100 2 3 4 5].
Это происходит, потому что v[0] - это ссылка, а v - значение.
Если мы хотим менять значение среза, нам нужно передавать его по ссылке
```
func someAction(v *[]int8, b int8) {
	(*v)[0] = 100
	*v = append(*v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(&a, 6)
	fmt.Println(a)
}
```