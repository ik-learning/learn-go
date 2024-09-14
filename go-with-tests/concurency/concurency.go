package concurrency
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency
type WebsiteChecker func(string) bool

func CheckWebsitesSequence(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = checker(url)
	}
	return results
}

type result struct {
	string
	bool
}

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, checker(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		res := <-resultChannel
		results[res.string] = res.bool
	}
	return results
}
