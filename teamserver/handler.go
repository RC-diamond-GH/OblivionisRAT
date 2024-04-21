package main

func GetBytes(data []byte, length int) []byte {
	if length > len(data) {
		length = len(data)
	}
	data = data[:length]
	return data
}

func POST_handler(body []byte) []byte {

}
