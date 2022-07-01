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

type UserKey struct {
	UserId string `json:"UserId"`
	PubKey []byte `json:"PubKey"`
}

var cntHeight int64
var flag = 0

func StartBCServer() {
	//서버 키면 지정한 경로에 Level DB 생성
	db, err := leveldb.OpenFile("/Users/byunjaejin/Go/level_DB", nil)
	if err != nil {
		panic(err)
	}

	//서버 키면 제네시스 블럭을 PBFT 로 전송

	bc := NewBlockchain()
	/*
		Block := bc.Blocks[len(bc.Blocks)-1]
		bytes, _ := json.Marshal(Block)
		buff := b.NewBuffer(bytes)
		Block.BPrint()

		resp, err := http.Post("http://192.168.10.57:4000/pbft", "application/json", buff)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
	*/

	// localhsot:80/create_bc 에 접속시
	// 넘어오는 값은 트랜잭션 내용
	http.HandleFunc("/create_bc", func(res http.ResponseWriter, req *http.Request) {

		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			TxData := []byte(respBody)

			DataForSign := &Data{}

			err = json.Unmarshal(TxData, DataForSign)
			if err != nil {
				println("Json Unmarshal Fail")
			}
			/*
				1. DataForSign 의 UserId 와 Sign 을 뽑아온다.
				2. LevelDB에서 UserId 가 가진 공개키를 가져온다.
				3. 가져온 공개키로 Sign 을 Verify() 한다.
				4. bool 값에 따라 처리한다.
			*/
			//1. UserId, Sign 값 가져오기
			UserID := DataForSign.UserID
			Sign := []byte{0}   //DataForSign.Sign
			HashId := []byte{0} //DataForSign.HashId
			//2.levelDB에서 ID에 맞는공개키 가져오기
			data, err := db.Get([]byte(UserID), nil)
			if err != nil {
				println("공개키 못 찾았음")
			}
			//3. Verify()
			if !Verify(data, Sign, HashId) {
				bytes, _ := json.Marshal(DataForSign)
				buff := b.NewBuffer([]byte(bytes))
				resp, err := http.Post("http://localhost:81/create_tx", "application/json", buff)

				if err != nil {
					panic(err)
				}

				defer resp.Body.Close()
			}

		}
		defer req.Body.Close()

	})

	/* 트랜잭션 검증이 끝난후 Tx서버에서 TxID를 넘겨 블록 생성 요청  */
	http.HandleFunc("/newblock", func(res http.ResponseWriter, req *http.Request) {
		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			TxID := []byte(respBody)
			bc.CreateBc(TxID) //블록 생성
		}

		/*--------- PBFT 서버에 POST 전달 --------*/
		if flag == 1 {
			flag = 0 //전송대기
			Block := bc.Blocks[cntHeight+1]
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

			} else {

				f.Println("------Reply------")
				f.Printf("Hash : %x\n", data.Hash)
				f.Printf("Height : %d\n", data.Height)
				f.Println("---------------")

				cntHeight = data.Height
				flag = 1 //전송 실행 상태 On
			}
		}

		defer req.Body.Close()

	})

	//회원의 ID에 맞는 공개키를 levelDB에 저장
	http.HandleFunc("/save_key", func(res http.ResponseWriter, req *http.Request) {
		respBody, err := ioutil.ReadAll(req.Body)
		UserKey := &UserKey{}
		if err == nil {
			err := json.Unmarshal(respBody, UserKey)
			if err != nil {
				println("유저키 언마샬 실패")
			}

			//Input DB			Key		        Value
			db.Put([]byte(UserKey.UserId), UserKey.PubKey, nil)
			//Key 로 Vaule Get
			DBdata, _ := db.Get([]byte(UserKey.UserId), nil)
			f.Println("저장된 유저 아이디와 공개키 정보")
			f.Printf("UserId : %s\n", UserKey.UserId)
			f.Printf("PubKey : %x\n", DBdata)

		}

	})

	http.ListenAndServe(":80", nil) //80번 포트에서 웹 서버 실행
}
