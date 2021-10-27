package main

func main() {
	err := consulRegister("192.168.0.102", 8501, "api-user", []string{"api", "user"}, "api-user")
	if err != nil {
		panic(err)
	}
	//consulDiscover()
}
