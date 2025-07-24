package main

import (
    "fmt"
    "runtime"
    "net/http"
    "net/url"
    "strconv"
    "github.com/PuerkitoBio/goquery"
    "github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/huh/spinner"
	"golang.org/x/term"
	"github.com/lu4p/cat"
	"strings"
)
func UserAgent() string{
    os := runtime.GOOS
    arch := runtime.GOARCH
    // Set the user agent string based on the OS and architecture
    var userAgent string
    switch os {
    case "windows":
        userAgent = fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; %s; rv:128.0) Gecko/20100101 Firefox/128.0", arch)
    case "linux":
        userAgent = fmt.Sprintf("Mozilla/5.0 (X11; Linux %s; rv:128.0) Gecko/20100101 Firefox/128.0", arch)
    case "darwin":
        userAgent = fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7; %s; rv:128.0) Gecko/20100101 Firefox/128.0", arch)
    default:
        userAgent = "Mozilla/5.0 (Unknown OS; Unknown Architecture; rv:128.0) Gecko/20100101 Firefox/128.0"
    }
    
    return userAgent
}
func LinkExtract(query string, times int) {
    list, _ := cat.File("list.txt")
    lines := strings.Split(list, "\n") // Corrected this line
    for _, line := range lines {
        // Split each line into its components
        parts := strings.Split(line, ";")
        if len(parts) < 3 {
            fmt.Println("Invalid line format:", line)
            continue
        }

        // Extract the URL, number of results, class, and base URL
        URL := strings.TrimSpace(parts[0]) + query
        class := strings.TrimSpace(parts[1])
        baseURL := strings.TrimSpace(parts[2])

        width, _, _ := term.GetSize(0)
        // This takes the HTML from the URL
        resp, err := http.Get(URL)
        if err != nil {
            fmt.Println("Error fetching URL:", err)
            return // Added return to exit on error
        }
        // Take the HTML body from the response and put it in a variable "doc"
        doc, _ := goquery.NewDocumentFromReader(resp.Body)
        count := 0
        // Scan the HTML body
        doc.Find(class).Each(func(i int, s *goquery.Selection) {
            if count < times {
                link, exists := s.Find("a").Attr("href") // It searches for <a> and from the "href" attribute takes the link
                linklist := baseURL + link
                if exists {
                    fmt.Println(
                        lipgloss.NewStyle().
                            Width(width - 5).
                            BorderStyle(lipgloss.RoundedBorder()).
                            BorderForeground(lipgloss.Color("63")).
                            Padding(0, 1).Render(linklist),
                    )
                    count++
                }
            }
        })
    }
}

func main() {
    var query string
    var ResNum string  
    game := huh.NewInput().Title("Search a game").Prompt("> ").Value(&query)
    resultsnumb := huh.NewInput().Title("how many results(per website) you want?").Prompt("> ").Value(&ResNum)
    
    huh.NewForm(huh.NewGroup(game)).Run()
    huh.NewForm(huh.NewGroup(resultsnumb)).Run()
    conResN, _ := strconv.Atoi(ResNum)
    
    fmt.Printf("Your links:\n")
    action := func() {
		LinkExtract(url.QueryEscape(query), conResN)
	}
	if err := spinner.New().Title("Preparing your links...").Action(action).Run(); err != nil {
		fmt.Println("Failed:", err)
		return
	}
    for{}
}

