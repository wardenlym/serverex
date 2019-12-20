package fswatcher

// func Test_watcher(t *testing.T) {
// 	path := os.TempDir() + "/test.txt"
// 	ioutil.WriteFile(path, []byte("testing"), 0664)

// 	modifyed := make(chan string)
// 	w := New(path, func(name string) {
// 		fmt.Println(name, "changed")
// 		modifyed <- ""
// 	})
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ioutil.WriteFile(path, []byte("42"), 0664)
// 	}()
// 	s := <-modifyed
// 	fmt.Println(s)
// 	w.Stop()
// 	os.Remove(path)

// }
