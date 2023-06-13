package main

//
//import (
//	"encoding/json"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//type EchoResponse struct {
//	Message string `json:"message"`
//}
//
//// Hàm routing echo, gồm hai params
//// wr http.ResponseWriter : dùng để ghi phản hồi về client
//// r *http.Request : dùng để đọc yêu cầu từ client
//func echo(wr http.ResponseWriter, r *http.Request) {
//	// Đọc thông điệp mà client gửi tới trong r.Body
//	body, err := ioutil.ReadAll(r.Body)
//	// Kiểm tra lỗi khi đọc body
//	if err != nil {
//		http.Error(wr, "Error reading request body", http.StatusBadRequest)
//		return
//	}
//	defer r.Body.Close()
//
//	// Tạo một struct để chứa thông điệp và gửi về client
//	response := EchoResponse{
//		Message: string(body),
//	}
//
//	// Chuyển đổi struct thành JSON
//	jsonResponse, err := json.Marshal(response)
//	// Kiểm tra lỗi khi chuyển đổi thành JSON
//	if err != nil {
//		http.Error(wr, "Error creating JSON response", http.StatusInternalServerError)
//		return
//	}
//
//	// Ghi phản hồi về client dưới dạng JSON
//	wr.Header().Set("Content-Type", "application/json")
//	wr.Write(jsonResponse)
//}
//
//// Hàm main của chương trình
//func main() {
//	// Mapping URL ứng với hàm routing echo
//	http.HandleFunc("/", echo)
//
//	// Địa chỉ http://127.0.0.1:8080/
//	err := http.ListenAndServe(":8080", nil)
//	// Ghi log ra lỗi nếu bị trùng port
//	if err != nil {
//		log.Fatal(err)
//	}
//}
