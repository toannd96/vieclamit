package feeds

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"vieclamit/common"
	"vieclamit/models"
	"vieclamit/repository"

	"github.com/PuerkitoBio/goquery"
)

const (
	topcvBasePath = "https://www.topcv.vn"
	topcvJobsPath = "/tim-viec-lam-it-phan-mem-c10026"
)

func GetTotalPageTopCV() (float64, error) {
	var numPage int

	url := fmt.Sprintf("%s%s", topcvBasePath, topcvJobsPath)
	doc, err := common.GetNewDocument(url)
	if err != nil {
		return 0, err
	}

	doc.Find("div.job-header span b.text-highlight").Each(func(index int, content *goquery.Selection) {
		numPage, _ = strconv.Atoi(strings.ReplaceAll(content.Text(), ",", ""))
	})

	numElement := doc.Find("div.job-item").Length()

	totalPage := math.Ceil(float64(numPage) / float64(numElement))

	fmt.Println(totalPage)

	return totalPage, nil
}

func GetDataOnePage(repo repository.Repository) error {
	var recruitment models.Recruitment
	var urlJob string

	doc, err := common.GetNewDocument("https://www.topcv.vn/tim-viec-lam-it-phan-mem-c10026")
	if err != nil {
		return err
	}

	doc.Find("div.body").Each(func(index int, body *goquery.Selection) {

		// url job
		body.Find("h3.title a[href]").Each(func(index int, href *goquery.Selection) {
			link, _ := href.Attr("href")
			urlJob = common.RemoveCharacterInString(link, "?")
		})

		// check url job exists in mongodb
		count, err := repo.FindByUrl(urlJob, "vieclamit")
		if err != nil {
			fmt.Println(err)
		}

		// if not exists
		if count == 0 {
			recruitment.UrlJob = urlJob

			body.Find("div.content").Each(func(index int, content *goquery.Selection) {

				// title
				content.Find("h3.title span.transform-job-title").Each(func(index int, title *goquery.Selection) {
					recruitment.Title = title.Text()
				})

				// job deadline
				content.Find("p.deadline strong").Each(func(index int, deadline *goquery.Selection) {
					numDay, _ := strconv.Atoi(deadline.Text())
					recruitment.JobDeadline = time.Now().AddDate(0, 0, numDay).Format("02/01/2006")
				})

				// company
				content.Find("p.company a[href]").Each(func(index int, href *goquery.Selection) {
					// name
					recruitment.Company = href.Text()

					// url on topcv
					link, _ := href.Attr("href")

					doc, _ := common.GetNewDocument(link)
					if !strings.Contains(link, "brand") {

						// url on company
						doc.Find("p.website a[href]").Each(func(index int, urlCompany *goquery.Selection) {
							recruitment.UrlCompany = urlCompany.Text()
						})
					} else {

						// url on company brand
						doc.Find("a.color-premium").Each(func(index int, urlCompany *goquery.Selection) {
							recruitment.UrlCompany = urlCompany.Text()
						})
					}
				})
			})

			body.Find("div.label-content").Each(func(index int, labelContent *goquery.Selection) {

				// salary
				labelContent.Find("label.salary").Each(func(index int, salary *goquery.Selection) {
					recruitment.Salary = salary.Text()
				})

				// location
				labelContent.Find("label.address").Each(func(index int, address *goquery.Selection) {
					recruitment.Location = address.Text()
				})
			})

			// save in to mongodb
			errSave := repo.Insert(recruitment, "vieclamit")
			if errSave != nil {
				fmt.Println(errSave)
			}

		} else {
			fmt.Printf("Exists %s\n", urlJob)
		}
	})

	fmt.Println("Done")

	return nil
}
