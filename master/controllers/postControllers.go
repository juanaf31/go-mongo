package controllers
import (
	"GoMongo/master/usecases"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type PostHandler struct{
	postUsecase usecases.PostUsecase
}

func PostController(r *mux.Router,service usecases.PostUsecase){
	postHandler := PostHandler{postUsecase: service}

	r.HandleFunc("/posts",postHandler.GetPost).Methods(http.MethodGet)
}

func(c *PostHandler) GetPost(w http.ResponseWriter, r *http.Request)  {
	posts,err := c.postUsecase.GetPosts()
	if err != nil {
		log.Println(err)
	}
	byteData,err := json.Marshal(posts)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "`+ err.Error() + `"}`))
	}else{
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}

	// json.NewEncoder(w).Encode(posts)

}