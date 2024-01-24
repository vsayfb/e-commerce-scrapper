package channel

import (
	"encoding/json"
	"fmt"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
	"github.com/gorilla/websocket"
	"github.com/vsayfb/e-commerce-scrapper/cache"
	"github.com/vsayfb/e-commerce-scrapper/cache/redis"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/search"
)

func ReceiveProducts(keyword string, page uint8, conn *websocket.Conn) {
	c := cache.New(redis.New())

	key := fmt.Sprintf("%v-%v", keyword, page)

	result, err := c.GetProducts(key)

	if err == nil {
		conn.WriteJSON(result)
	} else {
		s := search.New(keyword, page)

		ch := make(chan product.Product)

		quit := make(chan bool)

		tree := treemap.NewWith(utils.Float64Comparator)

		s.SearchAsync(ch)

		go func() {
			for {
				select {
				case p, open := <-ch:

					if open {
						tree.Put(p.NumPrice, p)

						if tree.Size()%10 == 0 {

							val := tree.Values()

							if writeErr := conn.WriteJSON(val); writeErr != nil {
								fmt.Println("Write error", writeErr)

								return
							}

						}
					} else {
						quit <- true
					}
				}
			}
		}()

		<-quit

		go func() {
			bytes, err := json.Marshal(tree.Values())

			if err != nil {
				fmt.Print("Marshal error", err)
			} else {
				c.Add(key, bytes)
			}
		}()

	}
}
