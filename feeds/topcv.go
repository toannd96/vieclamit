package feeds

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"

	"vieclamit/common"
	"vieclamit/models"
	"vieclamit/repository"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const (
	topcvBasePath = "https://www.topcv.vn"
	topcvJobsPath = "/tim-viec-lam-it-phan-mem-c10026"

	collection = "vieclamit"
)

// totalPageTopCV get total page
func totalPageTopCV() (int, error) {
	var numJob int

	url := fmt.Sprintf("%s%s", topcvBasePath, topcvJobsPath)
	doc, err := common.GetNewDocument(url)
	if err != nil {
		return 0, err
	}

	doc.Find("div.job-header span b.text-highlight").Each(func(index int, content *goquery.Selection) {
		numJob, _ = strconv.Atoi(strings.ReplaceAll(content.Text(), ",", ""))
	})

	numElement := doc.Find("div.job-item").Length()

	totalPage := int(math.Ceil(float64(numJob) / float64(numElement)))

	return totalPage, nil
}

// dataOnePage extract all item in a page
func dataOnePage(url string, repo repository.Repository) error {
	var recruitment models.Recruitment
	var urlJob, jobDeadlineString string

	doc, err := common.GetNewDocument(url)
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
		count, err := repo.FindByUrl(urlJob, collection)
		if err != nil {
			fmt.Println(err)
		}

		// if url job not exists
		if count == 0 {

			// job deadline
			body.Find("p.deadline strong").Each(func(index int, deadline *goquery.Selection) {
				numDay, _ := strconv.Atoi(deadline.Text())
				jobDeadlineString = time.Now().AddDate(0, 0, numDay).Format("02/01/2006")

			})

			// check the job deadline with the current time
			expired, _ := common.CheckTimeBefore(jobDeadlineString)

			//  expired is false is job deadline after current time
			if !expired {
				fmt.Printf("Extract %s\n", urlJob)

				recruitment.UrlJob = urlJob

				jobDeadline, _ := common.ParseTime(jobDeadlineString)
				recruitment.JobDeadline = jobDeadline

				// title
				body.Find("h3.title span.transform-job-title").Each(func(index int, title *goquery.Selection) {
					recruitment.Title = title.Text()
				})

				// company
				body.Find("p.company a[href]").Each(func(index int, href *goquery.Selection) {
					// name
					recruitment.Company = href.Text()

					// url on topcv
					link, _ := href.Attr("href")

					doc, _ := common.GetNewDocument(link)
					if !strings.Contains(link, "brand") {

						// url on company
						// ex: https://www.topcv.vn/viec-lam/ky-su-tich-hop-he-thong-luong-len-den-1000/595334.html
						doc.Find("p.website a[href]").Each(func(index int, urlCompany *goquery.Selection) {
							recruitment.UrlCompany = urlCompany.Text()
						})
					} else {
						// url on company brand
						// ex: https://www.topcv.vn/brand/educa/tuyen-dung/product-designer-manh-ve-ux-thu-nhap-hap-dan-j587195.html
						doc.Find("a.color-premium").Each(func(index int, urlCompany *goquery.Selection) {
							recruitment.UrlCompany = urlCompany.Text()
						})
					}
				})

				// salary
				body.Find("label.salary").Each(func(index int, salary *goquery.Selection) {
					recruitment.Salary = salary.Text()
				})

				// location
				body.Find("label.address").Each(func(index int, address *goquery.Selection) {
					recruitment.Location = address.Text()
				})

				// save in to mongodb
				errSave := repo.Insert(recruitment, collection)
				if errSave != nil {
					fmt.Println(errSave)
				}
			} else {
				fmt.Printf("job deadline %s before time today\n", jobDeadlineString)
			}

		} else {
			fmt.Printf("Exists %s\n", urlJob)
		}
	})

	return nil
}

// TopCV crawl all page it jobs
func TopCV(repo repository.Repository) {
	sem := semaphore.NewWeighted(int64(2 * runtime.NumCPU()))
	group, ctx := errgroup.WithContext(context.Background())

	totalPage, _ := totalPageTopCV()
	for page := 1; page <= totalPage; page++ {
		url := fmt.Sprintf("%s%s?page=%d", topcvBasePath, topcvJobsPath, page)
		err := sem.Acquire(ctx, 1)
		if err != nil {
			continue
		}
		group.Go(func() error {
			defer sem.Release(1)

			err := dataOnePage(url, repo)
			if err != nil {
				fmt.Println(err)
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Crawl completed")
}

// screenshotJDTopCV takes a screenshot of job descript topcv
func ScreenshotJDTopCV(url string) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte

	if strings.Contains(url, "brand") {
		// ex: https://www.topcv.vn/brand/smartosc/tuyen-dung/it-comtor-j592057.html
		if err := chromedp.Run(ctx, common.ElementScreenshot(url, "div.section-body", []string{"div.box-apply", "div.box-seo-job-detail"}, &buf)); err != nil {
			fmt.Println(err)
		}
	} else {
		// ex: https://www.topcv.vn/viec-lam/blockchain-developers-luong-1-000-4-000-hcm/590697.html
		if err := chromedp.Run(ctx, common.ElementScreenshot(url, "div.box-info-job div.col-md-8", []string{"div.box-how-to-apply"}, &buf)); err != nil {
			fmt.Println(err)
		}
	}

	if err := ioutil.WriteFile("screenshot_descript.png", buf, 0644); err != nil {
		fmt.Println(err)
	}
}
