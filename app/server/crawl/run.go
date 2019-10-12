package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"sync"
	"weserver/utils"
)
type Goods struct {
	Title string
	Price string
	Supplier string
	Cover string
}
type JdGood struct {
	ID int
	Title string
	Price string
	Cover string
	Supplier string
}
type Urls struct{
	Url string
	Pages int
}
func New(url string)*Urls{
	return &Urls{Url:url}
}
func (url *Urls)Start(){
	wg:=sync.WaitGroup{}
    for i:=1;i<=url.Pages*2-1;{
    	wg.Add(1)
    	url_page:=url.Url+"&page="+strconv.Itoa(i)
    	go select_goods(url_page)
    	i=i*2-1
	}

    //var goods []Goods
    //var good Goods
    //doc.Find("#J_goodsList").Find(".gl-warp").Find(".gl-item").Each(func(i int, selection *goquery.Selection) {
	//	  good.Title,_=selection.Find(".gl-i-wrap .p-name").Find("a").Attr("title")
	//	  good.Price=selection.Find(".gl-i-wrap .p-price").Find("i").Text()
	//	  good.CommentNum=selection.Find(".gl-i-wrap .p-commit").Find("a").Text()
	//	  good.Supplier=selection.Find(".gl-i-wrap .p-shop").Find("J_im_icon").Find("a").Text()
	//	  goods=append(goods,good)
	//})
    //fmt.Print(goods)
}
func select_goods(url string){
  res,err:=http.Get(url)
  if err!=nil{
  	log.Fatal(err)
  }
  defer res.Body.Close()
  doc,err1:=goquery.NewDocumentFromReader(res.Body)
  if err1!=nil{
  	log.Fatal(err1)
  }
	var goods []Goods
	var good Goods
	doc.Find("#J_goodsList").Find(".gl-warp").Find(".gl-item").Each(func(i int, selection *goquery.Selection) {
		good.Title,_=selection.Find(".gl-i-wrap .p-name").Find("a").Attr("title")
		good.Price=selection.Find(".gl-i-wrap .p-price").Find("i").Text()
		//good_url,_:=selection.Find(".gl-i-wrap .p-commit").Find("strong").Find("a").Attr("href")
		good.Supplier=selection.Find(".gl-i-wrap .p-shop").Find(".J_im_icon").Find("a").Text()

		good.Cover,_=selection.Find(".gl-i-wrap .p-img").Find("a").Html()
		goods=append(goods,good)
		//jd_good:=JdGood{
		//	Title:good.Title,
		//	Price:good.Price,
		//	Cover:good.Cover,
		//	Supplier:good.Supplier,
		//}
		//sqls.Db.Create(&jd_good)
	})
}
func (url *Urls)Page()*Urls{
	res,err:=http.Get(url.Url)
	if err!=nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode!=200{
		log.Fatal("status code error %d %s",res.StatusCode,res.Status)
	}
	doc,err:=goquery.NewDocumentFromReader(res.Body)
	if err!=nil{
		log.Fatal(err)
	}
	page:=doc.Find("#J_filter").Find("#J_topPage").Find(".fp-text").Find("i").Text()
	url.Pages=utils.StringToInt(page)
	return url
}