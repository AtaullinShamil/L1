# Что такое интерфейсы, как они применяются в Go?

Интерфейсы в Go представляют собой способ определения поведения типа. Они определяют набор методов, которые должен реализовать тип, чтобы соответствовать интерфейсу. Интерфейсы в Go реализуются неявно, то есть нет необходимости специально указывать, что структура или другой тип реализует определенный интерфейс

```
type Stringer interface {
   String() string
}
```

В этом примере интерфейс Stringer определяет один метод String(), который не принимает аргументов и возвращает строку.

Чтобы тип реализовал интерфейс, он должен реализовать все методы, определенные в интерфейсе. Например, структура Article может реализовать интерфейс Stringer, определив метод String():

```
type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

```

Теперь Article может быть использован в качестве Stringer:

```
func Print(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Shamil",
	}
	Print(a)
}

```
