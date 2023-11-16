package bot

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/vvjke314/MPPR/lab1/internal/chooser"
	"github.com/vvjke314/MPPR/lab1/internal/config"
	"github.com/vvjke314/MPPR/lab1/internal/repository"
	tele "gopkg.in/telebot.v3"
)

type Bot interface {
	StartBot()
}

type TelegramBot struct {
	cfg config.Config
	b   *tele.Bot
	pg  repository.PostgresRepo
}

func NewTelegramBot() TelegramBot {
	cfg := config.NewViperConfig("config", "./configs/")
	err := cfg.ReadConfig()
	if err != nil {
		panic(err)
	}

	pref := tele.Settings{
		Token:  cfg.GetValue("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	pg, err := repository.NewPostgresRepo(context.Background(), cfg.GetValue("DB_CONNECT"))
	if err != nil {
		panic(err)
	}

	return TelegramBot{
		cfg: cfg,
		b:   b,
		pg:  pg,
	}
}

func (tb TelegramBot) StartBot() {
	defer tb.pg.Close()

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnStart := menu.Text("/start")

	menu.Reply(
		menu.Row(btnStart),
	)

	tb.b.Use(tb.ExistenceCheck)
	//tb.b.Use(tb.UserState)

	tb.b.Handle("/start", func(c tele.Context) error {
		a1btn := menu.Text("Привет! Давай начнем")
		menu.Reply(
			menu.Row(a1btn),
		)
		return c.Send("Привет, я могу помочь тебе с выбором парфюмерии!", menu)
	})

	tb.b.Handle(tele.OnText, func(c tele.Context) error {
		chatId := strconv.Itoa(int(c.Chat().ID))
		state, err := tb.pg.GetState(chatId) // Получаем текущее состояние
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(state)

		// Развилка
		if state == 2 && c.Text() == "Для себя" {
			tb.pg.UpdateProgress(chatId, state, c.Text())
			state = 6
			c.Send("Какой ваш возраст?")
			tb.pg.UpdateState(chatId, 6)
		}

		switch state {
		case 0:
			// Сам вопрос
			menu.Reply()

			c.Send("Какой у вас бюджет?", menu)
		case 1:
			// Обновляем данные с текущего стейта
			budget, err := strconv.Atoi(c.Text())
			if err != nil {
				log.Printf("[%v STATE] Can't parse budget due to: %s", state, err)
				return err
			}
			tb.pg.UpdateProgress(chatId, state, strconv.Itoa(budget))

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			a1btn := menu.Text("Для себя")
			a2btn := menu.Text("Подарок")
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)

			c.Send("Для чего вам духи?", menu)
		case 2:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			a4btn := menu.Text(question.Answer4)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
				menu.Row(a4btn),
			)

			c.Send(question.QuestionText, menu)
		case 3:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)

			c.Send(question.QuestionText, menu)
		case 4:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			a4btn := menu.Text(question.Answer4)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
				menu.Row(a4btn),
			)

			c.Send(question.QuestionText, menu)
		case 5:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(21)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			menu.Reply(
				menu.Row(a1btn),
			)
			c.Send(question.QuestionText, menu)
		case 6:
			// Обновляем данные с текущего стейта
			age, err := strconv.Atoi(c.Text())
			if err != nil {
				log.Printf("[%v STATE] Can't parse age due to: %s", state, err)
				return err
			}
			tb.pg.UpdateProgress(chatId, state, strconv.Itoa(age))

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)
			c.Send(question.QuestionText, menu)
		case 7:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
			)
			c.Send(question.QuestionText, menu)
		case 8:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
			)
			c.Send(question.QuestionText, menu)
		case 9:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			a4btn := menu.Text(question.Answer4)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
				menu.Row(a4btn),
			)
			c.Send(question.QuestionText, menu)
		case 10:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
			)
			c.Send(question.QuestionText, menu)
		case 11:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)
			c.Send(question.QuestionText, menu)
		case 12:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)
			c.Send(question.QuestionText, menu)
		case 13:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			a4btn := menu.Text(question.Answer4)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
				menu.Row(a4btn),
			)
			c.Send(question.QuestionText, menu)
		case 14:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)
			c.Send(question.QuestionText, menu)
		case 15:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			a4btn := menu.Text(question.Answer4)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
				menu.Row(a4btn),
			)
			c.Send(question.QuestionText, menu)
		case 16:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)
			c.Send(question.QuestionText, menu)
		case 17:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
			)
			c.Send(question.QuestionText, menu)
		case 18:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
			)
			c.Send(question.QuestionText, menu)
		case 19:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			a2btn := menu.Text(question.Answer2)
			a3btn := menu.Text(question.Answer3)
			menu.Reply(
				menu.Row(a1btn),
				menu.Row(a2btn),
				menu.Row(a3btn),
			)
			c.Send(question.QuestionText, menu)
		case 20:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			// Обновляем кнопки и задаем вопрос для перехода к след стейту
			question, err := tb.pg.GetQuestion(state + 1)
			if err != nil {
				log.Printf("[%v STATE] %s", state, err)
				return err
			}
			a1btn := menu.Text(question.Answer1)
			menu.Reply(
				menu.Row(a1btn),
			)
			c.Send(question.QuestionText, menu)
		case 21:
			// Обновляем данные с текущего стейта
			tb.pg.UpdateProgress(chatId, state, c.Text())

			result, err := tb.pg.GetResult(chatId)
			if err != nil {
				log.Printf("[%v STATE] Can't get next state due to: %s", state, err)
				return err
			}

			choose := chooser.Choose(result)
			c.Send(choose)
			tb.pg.DeleteUser(chatId)
			return c.Send("Введите /start чтобы начать заново")
		}

		// Получаем следующий стейт
		nextState, err := tb.pg.NextState(chatId)
		if err != nil {
			log.Printf("[%v STATE] Can't get next state due to: %s", state, err)
			return err
		}

		// Обновляем стейт пользователя
		err = tb.pg.UpdateState(chatId, nextState)
		if err != nil {
			log.Println(err)
			return err
		}
		state = nextState
		return nil
	})

	tb.b.Start()
}

// ExistenceCheck checks in DB chat existence [Middleware]
func (tb TelegramBot) ExistenceCheck(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		chatId := strconv.Itoa(int(c.Chat().ID))
		check, err := tb.pg.CheckUserExistence(chatId)
		if err != nil {
			log.Println(err)
			return err
		}
		if check != true {
			log.Printf("New user[%s] started chatting", chatId)
			tb.pg.AddUser(chatId)
		}
		return next(c) // continue execution chain
	}
}
