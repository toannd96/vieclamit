package common

import (
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

// ElementScreenshot takes a screenshot of a specific element.
func ElementScreenshot(url, sel string, element []string, res *[]byte) chromedp.Tasks {
	actions := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("html", chromedp.ByQuery),
	}

	actions = append(actions, chromedp.ActionFunc(func(ctx context.Context) error {
		_, exp, err := runtime.Evaluate(RemoveSpecificElement(element)).Do(ctx)
		if err != nil {
			return err
		}

		if exp != nil {
			return exp
		}

		return nil
	}))

	actions = append(actions, chromedp.Tasks{
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	})

	return actions
}

// RemoveSpecificElement remove specific element
func RemoveSpecificElement(element []string) string {
	return fmt.Sprintf(`[%s].forEach(item => {
		const element = document.querySelector(item)
		if(element) element.remove()
	})`, fmt.Sprintf(
		"'%s'",
		strings.Join(element, `', '`),
	),
	)
}
