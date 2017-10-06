TARGET=myra-upload


$(TARGET): clean
	go build -ldflags="-s -w" -o $@
	upx --brute $@

clean:
	rm -f $(TARGET)


.PHONY:clean
