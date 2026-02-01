package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// type Produk struct {
// 	ID    int    `json:"id"`
// 	Nama  string `json:"nama"`
// 	Harga int    `json:"harga"`
// 	Stok  int    `json:"stok"`
// }

// var produk = []Produk{
// 	{ID: 1, Nama: "rizqi", Harga: 5000, Stok: 60},
// 	{ID: 1, Nama: "Doni", Harga: 6000, Stok: 30},
// 	{ID: 1, Nama: "Ari", Harga: 8000, Stok: 20},
// }

// type Category struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// var category = []Category{
// 	{ID: 1, Name: "Komedi", Description: "Tentang Lucu-lucuan"},
// 	{ID: 2, Name: "Horor", Description: "Tentang takut-takutan"},
// 	{ID: 3, Name: "Drama", Description: "Tentang lebay-lebayan"},
// }

// func getProdukByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, p := range produk {
// 		if p.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}

// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// // PUT localhost:8080/api/produk/{id}
// func updateProduk(w http.ResponseWriter, r *http.Request) {
// 	// get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

// 	// ganti int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	// get data dari request
// 	var updateProduk Produk
// 	err = json.NewDecoder(r.Body).Decode(&updateProduk)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	// loop produk, cari id, ganti sesuai data dari request
// 	for i := range produk {
// 		if produk[i].ID == id {
// 			updateProduk.ID = id
// 			produk[i] = updateProduk

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(updateProduk)
// 			return
// 		}
// 	}
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// func deleteProduk(w http.ResponseWriter, r *http.Request) {
// 	// get id
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
// 	// ganti id int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}
// 	// loop produk cari ID, dapet index yang mau dihapus
// 	for i, p := range produk {
// 		if p.ID == id {
// 			// bikin slice baru dengan data sebelum dan sesudah index
// 			produk = append(produk[:i], produk[i+1:]...)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "sukses delete",
// 			})

// 			return
// 		}
// 	}

// 	http.Error(w, "Produk belum ada", http.StatusNotFound)

// }

// func getCategoryByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "INvalid request", http.StatusBadRequest)
// 		return
// 	}

// 	for _, p := range category {
// 		if p.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}
// 	http.Error(w, "Category Belum ada", http.StatusNotFound)
// }

// func updateCategory(w http.ResponseWriter, r *http.Request) {
// 	//get id dari url
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}
// 	// loop produk cari ID, dapet index yang mau dihapus
// 	for i, p := range produk {
// 		if p.ID == id {
// 			// bikin slice baru dengan data sebelum dan sesudah index
// 			category = append(category[:i], category[i+1:]...)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "sukses delete",
// 			})

// 			return
// 		}
// 	}
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// func deleteCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "INvalid request", http.StatusBadRequest)
// 		return
// 	}

// 	//cari data
// 	for i := range category {
// 		if category[i].ID == id {
// 			category = append(category[:i], category[i+1:]...)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "sukses delete",
// 			})

// 			return
// 		}
// 	}
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	//product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	//category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
	http.HandleFunc("/api/category", categoryHandler.HandleCategory)

	//localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "Ok",
			"message": "Api Running",
		})

	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal runnung")
	}
}
