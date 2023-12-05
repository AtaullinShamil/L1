# Какой размер у структуры struct{}{}?

Рарзмер структуры struct{} равен нулю. Это можно проверить воспользовавшись функцией Size из пакета reflect

```
func main() {
   var s struct{}
   fmt.Println(reflect.TypeOf(s).Size())
}
```

При попытке проверить struct{}{}, компилятор выдает ошибку