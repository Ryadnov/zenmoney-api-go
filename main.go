package zenmoneyapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Ryadnov/zenmoney-api-go/entities"
)

type Api struct {
	BaseUrl string
	Token   string
}

func (a *Api) Sync(requestDiff entities.Diff) (*entities.Diff, error) {
	return sync(a.BaseUrl, a.Token, requestDiff)
}

func (a *Api) FullSync() (*entities.Diff, error) {
	requestDiff := entities.Diff{
		ServerTimestamp: "0",
	}

	return a.Sync(requestDiff)
}

func sync(url string, token string, requestDiff entities.Diff) (*entities.Diff, error) {
	////fmt.Printf("%#v \n", tokenResponse)
	////fmt.Printf("%#v \n", tokenResponse.AccessToken)
	//
	//diff := Diff{
	//	CurrentClientTimestamp: time.Now().UTC().Unix(),
	//	ServerTimestamp:        time.Now().UTC().Unix() - 10*60,
	//}

	requestDiff.CurrentClientTimestamp = entities.GetCurrentTimestamp()

	//'currentClientTimestamp' => time(),
	//	'serverTimestamp'        => time() - 10 * 60,
	jsonStr, err := json.Marshal(requestDiff)
	if err != nil {
		return nil, err
		//fmt.Printf("%#v \n", err)
	}
	//fmt.Printf("\n json: %s \n", string(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
		//fmt.Printf("%#v \n", err)
		//fmt.Printf("%#v \n", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	//client := &http.Client{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
		//fmt.Printf("%#v \n", err)
	}
	defer resp.Body.Close()

	//dump(req)
	//dump2(resp)

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	//ou

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseDiff *entities.Diff

	err = json.Unmarshal(body, responseDiff)
	if err != nil {
		return nil, err
	}

	return responseDiff, nil
}
