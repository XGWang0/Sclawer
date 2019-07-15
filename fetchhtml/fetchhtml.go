package fetchhtml

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"logo"
	"net/url"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
	SmzdmRootUrl string = "https://www.smzdm.com/jingxuan"
	//https:///www.smzdm.com/jingxuan/p1
)

type SMZDM struct {
	Title       string
	Price       string
	Link        string
	Vote        int
	Unvote      int
	CommentCont int
	CommentLink string
	DataTime    string
	Vendor      string
}

var ItemList []SMZDM

var SMZDMDocList []*goquery.Document
var AppendLock sync.Mutex

func HandleItemPage(li *goquery.Selection, smzdm *SMZDM, commentcnt, votecnt int) {
	var pg_wg sync.WaitGroup
	itemLink, _ := li.Find("a[onclick]").First().Attr("href")
	//logo.Log.Debug(itemLink)
	pg_wg.Add(1)
	go func(url string) {
		//fmt.Println("---------------------", url)
		var (
			itemTitle string
			itemPrice string
		)
		defer func() {
			smzdm.Title = strings.TrimSpace(itemTitle)
			smzdm.Price = strings.TrimSpace(itemPrice)
			smzdm.Link = itemLink

			pg_wg.Done()
		}()
		SMZDMItemDoc, err := goquery.NewDocument(url)
		if err != nil {
			logo.Log.Fatal(err.Error())
			os.Exit(1)
		}
		divSel := SMZDMItemDoc.Find("div[class=title-box]")
		itemTitle = divSel.Find("[class=title\\ J_title]").Text()
		itemPrice = divSel.Find("span").Text()
		//fmt.Println("Title, Price", itemTitle, itemPrice)

	}(itemLink)

	// Fetch comment, vote, unvote count
	li.Find("span[class=feed-btn-group]").Each(func(i int, span *goquery.Selection) {
		itemspansels := span.Find("span[class=unvoted-wrap]>span")
		itemVote := itemspansels.First().Text()
		itemUnVote := itemspansels.Last().Text()
		smzdm.Vote, _ = strconv.Atoi(itemVote)
		smzdm.Unvote, _ = strconv.Atoi(itemUnVote)
		//fmt.Printf("Vote: %#v\nUnVote:%#v\n", itemVote, itemUnVote)

		itemCommentsel := span.SiblingsFiltered("a[class=z-group-data]")
		itemCommentCount := strings.TrimSpace(itemCommentsel.Text())
		itemCommentLink, _ := itemCommentsel.Attr("href")
		smzdm.CommentCont, _ = strconv.Atoi(itemCommentCount)
		smzdm.CommentLink = itemCommentLink
		//fmt.Printf("Comment: %#v\nComment Link: %#v\n", itemCommentCount, itemCommentLink)

	})
	itemvendorsel := li.Find("span[class=feed-block-extras]")
	smzdm.DataTime = strings.TrimSpace(itemvendorsel.Contents().Not("a").Text())
	smzdm.Vendor = strings.TrimSpace(itemvendorsel.Find("a").Text())
	pg_wg.Wait()
	//fmt.Printf("One item:%#v\n", smzdm)

}

func HadelSinglePage(url string, commentcnt, votecnt int) {
	//SMZDMDoc, err := goquery.NewDocument(SmzdmRootUrl)
	SMZDMDoc, err := goquery.NewDocument(url)
	if err != nil {
		logo.Log.Fatal(err.Error())
		os.Exit(1)
	}

	var li_wg sync.WaitGroup

	SMZDMDoc.Find("ul[id=feed-main-list]").Each(func(i int, ul *goquery.Selection) {
		smzdm := SMZDM{}
		//fmt.Println(ul.Find("a").Has("onclick"))
		// Get vote unvote and comments for each product
		ul.Find("li[class=feed-row-wide]").Each(func(i int, li *goquery.Selection) {
			li_wg.Add(1)
			go func(li *goquery.Selection) {
				defer func() {
					AppendLock.Lock()
					if smzdm.Title != "" {
						if smzdm.CommentCont > commentcnt || smzdm.Vote >= votecnt {
							ItemList = append(ItemList, smzdm)
						}
					}
					AppendLock.Unlock()
					li_wg.Done()
				}()
				HandleItemPage(li, &smzdm, commentcnt, votecnt)

			}(li)
			li_wg.Wait()
		})
	})

}

func HandelAllUrl(page, commentcnt, votecnt int) {
	logo.Log.Debug("ffffffffffffffffffffff")
	ItemList = ItemList[:0]
	var url_wg sync.WaitGroup
	for i := 1; i <= page; i++ {
		url_wg.Add(1)
		u, _ := url.Parse(SmzdmRootUrl)
		u.Path = path.Join(u.Path, "p"+strconv.Itoa(i))
		eachurl := u.String()
		go func(url string) {
			defer url_wg.Done()
			HadelSinglePage(eachurl, commentcnt, votecnt)
		}(eachurl)
	}
	url_wg.Wait()
	SortItemList()
	//PringItemList()
}

func SortItemList() {
	sort.SliceStable(ItemList, func(i, j int) bool {
		if ItemList[i].CommentCont > ItemList[j].CommentCont {
			return true
		} else if ItemList[i].CommentCont < ItemList[j].CommentCont {
			return false
		}
		if ItemList[i].Vote > ItemList[j].Vote {
			return true
		} else if ItemList[i].Vote < ItemList[j].Vote {
			return false
		}
		return true
	})
}

func PringItemList() {
	/*
		sort.SliceStable(ItemList, func(i, j int) bool {
			if ItemList[i].CommentCont > ItemList[j].CommentCont {
				return true
			} else if ItemList[i].CommentCont < ItemList[j].CommentCont {
				return false
			}
			if ItemList[i].Vote > ItemList[j].Vote {
				return true
			} else if ItemList[i].Vote < ItemList[j].Vote {
				return false
			}
			return true
		})
	*/
	SortItemList()
	fmt.Printf("Total Filter %d Items\n", len(ItemList)+1)
	for i, value := range ItemList {
		fmt.Printf("------------------------------------------\n")
		fmt.Printf("No.%d [%s %s]\n", i+1, value.Title, value.Price)
		fmt.Printf("Product Link  : %s\n", value.Link)
		fmt.Printf("Deliver Time  : %s\n", value.DataTime)
		fmt.Printf("Comment Count : %d\n", value.CommentCont)
		fmt.Printf("[UN]Vote Count : (V)%d, (UV)%d)\n", value.Vote, value.Unvote)
		fmt.Printf("Product Vendor  : %#v\n", value.Vendor)
	}
	fmt.Printf("\n------------------------------------------\n")
}
