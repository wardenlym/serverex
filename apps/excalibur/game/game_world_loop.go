package game

import (
	"time"
)

func (g *GameService) loop() {

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	func() {
		for {
			select {
			case <-g.loop_stoper:
				return
			case <-ticker.C:
				g.update()
			}
		}
	}()
}

func (g *GameService) update() {

}

func (g *GameService) stop_loop() {
	close(g.loop_stoper)
}
