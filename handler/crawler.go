package handler


import (
	"github.com/labstack/echo/v4"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"crawler-website/helper"
	"strings"
	"regexp"
	"html"
)

type crawlerHandler struct {
	
}

func InitCrawlerHandler() *crawlerHandler {
	return &crawlerHandler{}
}

func (h crawlerHandler) FetchData(links string) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		httpClient := &http.Client{
			Timeout: 60 * time.Second,
		}

		scraper := colly.NewCollector(
			colly.Async(true),
			colly.AllowURLRevisit(),
			colly.MaxDepth(1),
		)

		scraper.WithTransport(httpClient.Transport)

		scraper.SetRequestTimeout(60 * time.Second)

		var htmlStr strings.Builder
		scraper.OnResponse(func(e *colly.Response) {
			htmlStr.Write(e.Body)
		})

		scraper.OnHTML("script", func(e *colly.HTMLElement) {
			scriptText := e.Text
			scriptText = html.UnescapeString(scriptText)
			prettyScript := prettifyScript(scriptText)
			htmlStr.WriteString("<script>")
			htmlStr.WriteString(prettyScript)
			htmlStr.WriteString("</script>")
		})

		scraper.OnError(func(e *colly.Response, err error) {
			stringErr := fmt.Sprintf("failed with response '%s'\n", err)
			response := helper.APIResponse(stringErr, http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		})
		
		scraper.OnScraped(func(e *colly.Response) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr.String()))
			if err != nil {
				fmt.Println("Error loading HTML:", err)
			}

			doc.Find("noscript").Each(func(i int, noscript *goquery.Selection) {
				noscript.AfterHtml("\n")
			})
			
			doc.Find("div").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})
			
			doc.Find("main").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})
			
			doc.Find("footer").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})
			
			doc.Find("span").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})
			
			doc.Find("img").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})
			
			doc.Find("script").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			doc.Find("button").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			doc.Find("li").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			doc.Find("ul").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			doc.Find("section").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			doc.Find("p").Each(func(i int, s *goquery.Selection) {
				s.AfterHtml("\n")
			})

			prettyHTML, err := doc.Html()
			if err != nil {
				fmt.Println("Error getting HTML:", err)
			}
			currentTime := time.Now()
			filePath := fmt.Sprintf("hasil_crawling_%s.html", currentTime.Format("2006-01-02_15-04-05"))
			err = ioutil.WriteFile(filePath, []byte(prettyHTML), 0644)
			if err != nil {
				stringErr := fmt.Sprintf("Error save HTML:'%s'\n", err)
				response := helper.APIResponse(stringErr, http.StatusBadRequest, "error", nil)
				c.JSON(http.StatusBadRequest, response)
				return
			}
			
			stringprint := fmt.Sprintf("File '%s' berhasil disimpan", filePath)
			stringsuccess := fmt.Sprintf("Berhasil melakukan crawler ke website %s", links)
			
			response := helper.APIResponse(stringsuccess, http.StatusOK, "success", map[string]string{
				"message": stringprint,
			})
			
			c.JSON(http.StatusOK, response)
		})
		
		scraper.Visit(links)

		scraper.Wait()
		return nil
	}

}

func prettifyScript(script string) string {
	script = strings.ReplaceAll(script, "  ", " ")
	script = strings.ReplaceAll(script, "\n\n", "\n")
	re := regexp.MustCompile(`{\s*`)
	script = re.ReplaceAllString(script, "{\n")
	re = regexp.MustCompile(`\s*}`)
	script = re.ReplaceAllString(script, "\n}")
	re = regexp.MustCompile(`;\s*`)
	script = re.ReplaceAllString(script, ";\n")
	re = regexp.MustCompile(`,\s*`)
	script = re.ReplaceAllString(script, ", ")
	lines := strings.Split(script, "\n")
	for i, line := range lines {
		lines[i] = "\t" + line
	}
	script = strings.Join(lines, "\n")
	return script
}