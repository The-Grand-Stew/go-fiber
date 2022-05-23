package controllers

import (
	"{{ .AppName }}/app/queries"

	"github.com/gofiber/fiber/v2"
)

var {{ .DomainName | ToLower }} *queries.{{ .DomainName }}

func Create{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func Get{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func GetAll{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func Update{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil
}

func Delete{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil
}
