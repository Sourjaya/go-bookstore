package utils

/*ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
*/
import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
