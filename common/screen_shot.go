package common

import "github.com/chromedp/chromedp"

// ElementScreenShot takes a screenshot of a specific element.
func ElementScreenShot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}
