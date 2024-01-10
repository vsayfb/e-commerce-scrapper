package channel

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
	"github.com/gorilla/websocket"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/search"
)

func ReceiveProducts(msg string, page uint8, conn *websocket.Conn) {

	s := search.New(msg, page)

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

}
