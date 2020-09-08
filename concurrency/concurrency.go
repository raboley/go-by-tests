package concurrency

type WebsiteChecker func(string) bool
type CheckedWebsites struct {
	url       string
	reachable bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	allCheckedWebsites := make(map[string]bool)
	checkedWebsitesChannel := make(chan CheckedWebsites)

	for _, url := range urls {
		go func(u string) {
			checkedWebsitesChannel <- CheckedWebsites{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		checkedWebsite := <-checkedWebsitesChannel
		allCheckedWebsites[checkedWebsite.url] = checkedWebsite.reachable
	}

	return allCheckedWebsites
}
