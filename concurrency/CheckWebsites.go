package concurrency


type WebsiteChecker func (string) bool 
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls [] string )map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		go func() {

			resultsChannel <- result{url, wc(url)}
		}()//go routines and anonymous functions,
		
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}
	
	return results

}