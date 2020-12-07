package main

import (
    "crypto/sha1"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "math"
    "math/rand"
    "os"
    "strconv"
)
type Ad struct {
    Id string `json:"id"`
    Text string `json:"text"`
    Tags []string `json:"tags"`
    Rate float64  `json:"rate"`
}
func main (){
    adsNumberParce, err := strconv.ParseFloat(os.Args[1], 32);
    if err != nil {
      panic(err)
    }
    adsNumber := int(adsNumberParce)
    tagsNames := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
    tagsMaskSize := 8
    tagsMaskMax := 255
    for i:=0; i<adsNumber; i++ {
        var tags []string
        tagsMask := uint32(math.Round(float64(tagsMaskMax) * rand.Float64()))
        for t:=0; t<tagsMaskSize; t++ {
            if (tagsMask & (1<<t))>0 {
                tags = append(tags, tagsNames[t])
            }
        }
        rate := rand.Float64()
        sha := sha1.Sum([]byte(strconv.Itoa(i)))
        id := hex.EncodeToString(sha[:])
        text := fmt.Sprintf("Статья с тегами %s и рейтингом %f.", tags, rate)
        json, _ := json.Marshal(Ad{id, text, tags, rate})
        fmt.Println(string(json))
    }
}//}