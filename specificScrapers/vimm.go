package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Vimm() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("vimm.json")
	specificInfo := config.SpecificInfo
	itemKeys := config.Search.ItemKeys

	// imgCookies := helpers.GetSiteCookies("https://vimm.net")
	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {
		item := variables.Item{
			Name: h.ChildText(itemKeys.Name),
			Link: h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}
		// a cookie is needed to load the image, otherwise vimm returns a default image, the cookie is generated by a post request with some tokens. Either figure out where those tokens come from, or open a browser with go-rod on the home page, and get the cookies from the browser
		// imgUrl := "https://vimm.net/image.php?type=box&id=" + strings.Split(item.Link, "/vault/")[1]

		// item.Thumbnail = helpers.GetImageBase64(imgUrl, imgCookies)

		if item.Name != "" {
			item.Metadata = map[string]string{
				"console": h.ChildText(specificInfo["console"]),
			}
			//sometimes the region is not referenced
			flag := strings.Split(h.ChildAttr(specificInfo["region"], "src"), "/flags/")
			if len(flag) >= 2 {
				item.Metadata["region"] = strings.ReplaceAll(flag[1], ".png", "")
			}
			results = append(results, item)
		}
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
