package image

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type transport struct {
	rt http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	return t.rt.RoundTrip(req)
}

func getBody(client *http.Client, url string) (io.ReadCloser, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		res.Body.Close()
		return nil, errors.New(string(body))
	}
	return res.Body, nil
}

func getImgUrls(r io.Reader, selector string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	var urls []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		img := s.Find("img")
		srcset, exist := img.Attr("srcset")
		var url string
		if exist {
			tmp := strings.Split(srcset, ",")
			last := tmp[len(tmp)-1]
			url = strings.Split(strings.TrimSpace(last), " ")[0]
		} else {
			url, exist = img.Attr("src")
			if !exist {
				logrus.Warnf("%s not found", selector)
			}
		}
		urls = append(urls, url)
	})
	return urls, nil
}

func download(client *http.Client, url string, outdir string) error {
	logrus.Infof("dowloading: %s", url)
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	filename := path.Base(url)
	f, err := os.Create(path.Join(outdir, filename))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func Download(rootUrl string, selector string, outdir string, limit int, concurrency int) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	body, err := getBody(&client, rootUrl)
	if err != nil {
		return err
	}
	defer body.Close()

	urls, err := getImgUrls(body, selector)
	if err != nil {
		return err
	}
	u, err := url.Parse(rootUrl)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(outdir, 0755); err != nil {
		return err
	}

	if concurrency <= 0 {
		concurrency = 1
	}

	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for p := range c {
				parsedURL, err := url.Parse(p)
				if err != nil {
					logrus.Error(err)
				}
				if parsedURL.Scheme == "" {
					u.Path = p
					p = u.String()
				}

				if err := download(&client, p, outdir); err != nil {
					logrus.Errorf("%s download err: %v", p, err)
				}
			}
		}()
	}

	for i, p := range urls {
		if limit > 0 && i == limit {
			break
		}
		c <- p

	}
	close(c)

	wg.Wait()

	return nil
}
