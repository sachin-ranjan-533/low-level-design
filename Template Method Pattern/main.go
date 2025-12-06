package main

import "template-method-pattern/templates"

func main() {
	payToFriend := templates.NewPayToFriend(101, 102, 10.0)
	templates.SendMoney(payToFriend, 101, 102, 10.0)
}
