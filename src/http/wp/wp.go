package wp

import "net/http"

func init() {

}
func RunService()  {
	http.ListenAndServe(":8089", nil)
}
