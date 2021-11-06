package quiz

import (
	"github.com/dennybiasiolli/go-quiz/auth"
	"github.com/gofiber/fiber/v2"
)

func QuizRoutesRegister(router fiber.Router) {
	quizzes := router.Group("/quizzes")
	quizzes.Get("/", QuizList)
	quizzes.Get("/:id/", QuizDetail)
	quizzes.Use(auth.IsAdminMiddleware()).Post("/", QuizCreate)
	quizzes.Use(auth.IsAdminMiddleware()).Put("/:id/", QuizUpdate)
	quizzes.Use(auth.IsAdminMiddleware()).Delete("/:id/", QuizDelete)

	questions := router.Group("/questions")
	questions.Use(auth.IsAdminMiddleware()).Post("/", QuestionCreate)
	questions.Use(auth.IsAdminMiddleware()).Put("/:id/", QuestionUpdate)
	questions.Use(auth.IsAdminMiddleware()).Delete("/:id/", QuestionDelete)
}
