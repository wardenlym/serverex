package excalibur

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	toml "github.com/pelletier/go-toml"
)

func Test_usingTOML(t *testing.T) {

	type Doc struct {
		AisNumber    int
		TimeIsInline time.Time
		B            string
		C            float32
		D            []int
		F            map[string]int
		GisMapOfMap  map[string]map[string][]int
		HisListofMap []map[string][]int
	}

	doc := Doc{
		AisNumber:    42,
		TimeIsInline: time.Now(),
		B:            "toml is good for human reading and hand writing",
		C:            -3.1415926,
		D:            []int{1, 2, 3, 4},
		F:            map[string]int{"a": 1, "42": 2},
		GisMapOfMap: map[string]map[string][]int{
			"key1": map[string][]int{"bbb": []int{1, 2, 3}},
			"key2": map[string][]int{"ddd": []int{4, 5, 6}},
		},
		HisListofMap: []map[string][]int{
			map[string][]int{"xxxx": []int{1, 2, 3, 4}},
			map[string][]int{"yyyy": []int{5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			map[string][]int{"zzzz": []int{-1, -2, -3, -4}},
		},
	}

	b, err := toml.Marshal(doc)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}

func Test_parseTOML(t *testing.T) {

	b, err := toml.Marshal(betaModeConfig)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}

func Test_codegen_WriteDefaultConfig(t *testing.T) {

	b, err := toml.Marshal(betaModeConfig)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
	err = ioutil.WriteFile("conf/config.beta.toml", b, 0644)
	if err != nil {
		t.Error(err)
	}

	b, err = toml.Marshal(devModeConfig)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
	err = ioutil.WriteFile("conf/config.dev.toml", b, 0644)
	if err != nil {
		t.Error(err)
	}

	b, err = toml.Marshal(approvalModeConfig)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
	err = ioutil.WriteFile("conf/config.approval.toml", b, 0644)
	if err != nil {
		t.Error(err)
	}

	b, err = toml.Marshal(prodModeConfig)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
	err = ioutil.WriteFile("conf/config.prod.toml", b, 0644)
	if err != nil {
		t.Error(err)
	}

	b, err = toml.Marshal(taptapModeConfig)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
	err = ioutil.WriteFile("conf/config.taptap.toml", b, 0644)
	if err != nil {
		t.Error(err)
	}

}
