package handler

import (
	"github.com/gofiber/fiber/v2"
	"main/model"
)

var cards = make([]model.Card, 0)

func GetAllCards(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "data": &cards})
}

func GetCard(c *fiber.Ctx) error {
	id := c.Params("id")
	var card = model.Card{}
	for _, n := range cards {
		if n.Id == id {
			card = n
		}
	}

	if card.Content == "" {
		return c.JSON(fiber.Map{"status": "error", "message": "карты с таким Id не найдено", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "data": card})
}

func CreateCard(c *fiber.Ctx) error {
	card := new(model.Card)

	if err := c.BodyParser(card); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	if card.Content == "" || card.Id == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product"})
	}
	cards = append(cards, *card)
	return c.JSON(fiber.Map{"status": "success", "message": "Created card", "data": card})
}

func DeleteCard(c *fiber.Ctx) error {
	id := c.Params("id")
	var cardIndex = -1
	for i, n := range cards {
		if n.Id == id {
			cardIndex = i
		}
	}
	if cardIndex == -1 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Карты с таким id не существует"})
	}

	// 1. Копировать последний элемент в индекс i.
	cards[cardIndex] = cards[len(cards)-1]

	// 2. Удалить последний элемент (записать нулевое значение).
	cards[len(cards)-1] = model.Card{}

	// 3. Усечь срез.
	cards = cards[:len(cards)-1]

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted card", "data": cards})
}
