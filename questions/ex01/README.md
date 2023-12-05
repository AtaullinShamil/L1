# Какой самый эффективный способ конкатенации строк?

Самым эффеективным способоб конкатенации строк является использование strings.Builder.
При его использовании не происходит не происходит повторного выделения памяти и копирования данных.
```
func main() {
	var builder strings.Builder

	builder.WriteString("Привет ")
	builder.WriteString("Мир!")
	fmt.Println(builder.String())
}
```
