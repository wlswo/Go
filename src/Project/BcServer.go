package Project

import (
	b "bytes"
	"encoding/json"
	f "fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Hash   []byte `json:"Hash"`
	Height int64  `json:"Height"`
}

func StartBCServer() {
	//서버 키면 지정한 경로에 Level DB 생성
	db, err := leveldb.OpenFile("/Users/byunjaejin/Go/level_DB", nil)
	if err != nil {
		panic(err)
	}

	//서버 키면 블록체인 구조체 생성
	bc := NewBlockchain()
	Block := bc.Blocks[len(bc.Blocks)-1]
	bytes, _ := json.Marshal(Block)
	buff := b.NewBuffer(bytes)
	Block.BPrint()
	resp, err := http.Post("http://192.168.10.57:4000/pbft", "application/json", buff)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// localhsot:80/test 에 접속시
	// 넘어오는 값은 트랜잭션 내용

	http.HandleFunc("/create_bc", func(res http.ResponseWriter, req *http.Request) {

		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			TxData := []byte(respBody)

			DataForSign := &Data{}

			err = json.Unmarshal(TxData, DataForSign)

			/*
				1. DataForSign 의 UserId 와 Sign 을 뽑아온다.
				2. LevelDB에서 UserId 가 가진 공개키를 가져온다.
				3. 가져온 공개키로 Sign 을 Verify() 한다.
				4. bool 값에 따라 처리한다.
			*/
			//1. UserId, Sign 값 가져오기
			UserID := DataForSign.UserID
			Sign := DataForSign.Sign

			//2.levelDB에서 ID에 맞는공개키 가져오기
			data, err := db.Get([]byte(UserID), nil)

			bc.CreateBc(TxID) //블록 생성
		}
		defer req.Body.Close()

		//b, _ := ioutil.ReadFile("BlockFile.json") //file read
		//res.Write([]byte(b)) //웹 브라우저에 응답

		/*--------- PBFT 서버에 POST 전달 --------*/
		Block := bc.Blocks[len(bc.Blocks)-1]
		bytes, _ := json.Marshal(Block)
		buff := b.NewBuffer(bytes)

		resp, err := http.Post("http://192.168.10.57:4000/pbft", "application/json", buff)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		respBody, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			str := string(respBody)
			println(str)
		}
		/*--------------------------------------*/
	})

	//pbft -> Here Responce 합의 완료 답장
	/*
		{
			Hash :
			Height :
		}
	*/
	http.HandleFunc("/reply", func(res http.ResponseWriter, req *http.Request) {
		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			data := &Result{}
			err := json.Unmarshal([]byte(respBody), data)

			if err != nil {
				f.Println("에러  : 합의 답장 실패")
			}

			f.Println("------Reply------")
			f.Printf("Hash : %x\n", data.Hash)
			f.Printf("Height : %d\n", data.Height)
			f.Println("---------------")

		}
		defer req.Body.Close()
	})

	http.ListenAndServe(":80", nil) //80번 포트에서 웹 서버 실행
}

// 바이트를 문자열로

func decode(b []byte) string {
	return string(b[:len(b)])
}

/*
//Input DB			    Key				Value
			db.Put([]byte(string(data.Height)), []byte(data.Hash), nil)
			//Key 로 Vaule Get
			DBdata, _ := db.Get([]byte(string(data.Height)), nil)
			f.Println(decode(DBdata))

*/
