// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Order {
	id : String!
	Price : Float!
	Tax : Float!
	FinalPrice : Float!
}

type OrderInput {
	id : String!
	Price : Float!
	Tax : Float!
}

type Mutation {
	createOrder(input: OrderInput): Order
}