package main

type Client struct {
	name string
}

func teste(c Client) {
	println(c.name)
}

func main() {
	client := Client{
		name: "Gustavo",
	}
	teste(client)
}
