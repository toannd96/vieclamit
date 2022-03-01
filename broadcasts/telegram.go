package broadcasts

import (
	"fmt"
	"strings"
	"time"
	"vieclamit/repository"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Telegram struct {
	Token string
	Repo  repository.Repository
}

func (t *Telegram) NewTelegram() {
	pref := tb.Settings{
		Token:  t.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		fmt.Println(err)
	}

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, `👋 Tôi là bot việc làm IT
✅ Tôi có thể tìm kiếm tin tuyển dụng việc làm IT theo từ khóa địa điểm, kỹ năng, công ty
🔎 Để tôi giúp bạn hiểu cách hoạt động /help`)
	})

	bot.Handle("/help", func(m *tb.Message) {
		bot.Send(m.Sender, `✅ Từ khóa bạn nhập không phân biệt chữ hoa chữ thường, phải có dấu

✅ Từ khóa tên địa điểm 
							👉 /location <tên địa điểm>
							Ví dụ:
								👍 /location hà nội
								👎 /location ha noi

✅ Từ khóa tên công ty 
							👉 /company <tên công ty>
							Ví dụ:
								👍 /company smartosc
								👍 /company giao hàng tiết kiệm
								👎 /company giao hang tiet kiem

✅ Từ khóa tên kỹ năng
							👉 /skill <tên kỹ năng>
							Ví dụ:
								👍 /skill golang`)
	})

	bot.Handle("/location", func(m *tb.Message) {
		location := strings.TrimSpace(m.Text[9:])
		if location == "" {
			bot.Send(m.Sender, "💡 Nhập tên thành phố có công việc bạn muốn tìm. Ví dụ: /location Hà nội")
			return
		}
		recruitments, err := t.Repo.FindByLocation(location)
		if err != nil {
			fmt.Println(err)
		}
		for _, recruitment := range *recruitments {
			output := getTemplate(
				recruitment.Title,
				recruitment.Company,
				recruitment.Location,
				recruitment.Salary,
				recruitment.JobDeadline.Format("02/01/2006"),
				recruitment.UrlJob,
				recruitment.UrlCompany,
			)
			bot.Send(m.Sender, output, &tb.SendOptions{
				ParseMode:             "Markdown",
				DisableWebPagePreview: true,
			})
		}
	})

	bot.Handle("/company", func(m *tb.Message) {
		company := strings.TrimSpace(m.Text[8:])
		if company == "" {
			bot.Send(m.Sender, "💡 Nhập tên công ty có công việc bạn muốn tìm. Ví dụ: /company smartosc")
			return
		}
		recruitments, err := t.Repo.FindByCompany(company)
		if err != nil {
			fmt.Println(err)
		}
		for _, recruitment := range *recruitments {
			output := getTemplate(
				recruitment.Title,
				recruitment.Company,
				recruitment.Location,
				recruitment.Salary,
				recruitment.JobDeadline.Format("02/01/2006"),
				recruitment.UrlJob,
				recruitment.UrlCompany,
			)
			bot.Send(m.Sender, output, &tb.SendOptions{
				ParseMode:             "Markdown",
				DisableWebPagePreview: true,
			})
		}
	})

	bot.Handle("/skill", func(m *tb.Message) {
		skill := strings.TrimSpace(m.Text[6:])
		if skill == "" {
			bot.Send(m.Sender, "💡 Nhập tên kỹ năng bạn muốn tìm. Ví dụ: /skill php")
			return
		}
		recruitments, err := t.Repo.FindBySkill(skill)
		if err != nil {
			fmt.Println(err)
		}
		for _, recruitment := range *recruitments {
			output := getTemplate(
				recruitment.Title,
				recruitment.Company,
				recruitment.Location,
				recruitment.Salary,
				recruitment.JobDeadline.Format("02/01/2006"),
				recruitment.UrlJob,
				recruitment.UrlCompany,
			)
			bot.Send(m.Sender, output, &tb.SendOptions{
				ParseMode:             "Markdown",
				DisableWebPagePreview: true,
			})
		}
	})

	bot.Start()
}

func getTemplate(title, company, location, salary, jobDeadline, urlJob, urlCompany string) string {
	return fmt.Sprintf("*%s - %s*\n"+"🏢 %s\n"+"💰 %s\n"+"⏳ %s\n"+"👉 [%s](%s)\n"+"👉 [%s](%s)\n", title, company, location, salary, jobDeadline, "Xem tin tuyển dụng", urlJob, "Xem công ty", urlCompany)
}
