package quiz

import (
	"errors"

	"github.com/dennybiasiolli/go-quiz/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func QuizList(c *fiber.Ctx) error {
	db := common.GetDB()
	var quizzes []Quiz
	var count int64
	err := db.Find(&quizzes).Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"results": QuizzesSerializer(quizzes),
		"count":   count,
	})
}

func QuizDetail(c *fiber.Ctx) error {
	db := common.GetDB()
	var quiz Quiz
	id := c.Params("id")
	err := db.Preload("Questions").First(&quiz, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(QuizSerializer(quiz))
}

func QuizCreate(c *fiber.Ctx) error {
	input := new(Quiz)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	db := common.GetDB()
	if err := db.Omit(clause.Associations).Create(&input).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(QuizSerializer(*input))
}

func QuestionCreate(c *fiber.Ctx) error {
	input := new(Question)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	db := common.GetDB()
	if err := db.Omit(clause.Associations).Create(&input).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(QuestionSerializer(*input))
}
