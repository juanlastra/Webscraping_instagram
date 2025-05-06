package extraer_Instagram

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/MrAndiw/instagram-go-scraper/instagram"
)

type RequestBody struct {
	Link string `json:"link"`
}

type ResponseData struct {
	Usuario    string `json:"usuario"`
	Posts      int    `json:"posts"`
	Seguidores int    `json:"seguidores"`
	Seguidos   int    `json:"seguidos"`
}

func Corregir_link(url string) string {
	re := regexp.MustCompile(`https:\/\/www\.instagram\.com\/([^\/]+)`)

	match := re.FindStringSubmatch(url)

	if len(match) > 1 {
		return match[1]
	}

	return ""
}

// función para extraer datos
func Extraer_info_ig(usuario string) ResponseData {
	ig := instagram.Init()

	Instagram := instagram.NewInstagram(ig)

	Usuario := Instagram.GetFullBio(Corregir_link(usuario))

	info := ResponseData{
		Usuario:    Usuario.Title,
		Posts:      Usuario.Post,
		Seguidores: Usuario.Followers,
		Seguidos:   Usuario.Following,
	}

	return info
}

func Enviar_info(w http.ResponseWriter, r *http.Request) {

	var link RequestBody

	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, "Error al parsear la solicitud", http.StatusBadRequest)
		return
	}

	data := Extraer_info_ig(link.Link)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}

// Middleware para habilitar CORS
func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite todo origen. En producción sería mejor especificar.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
