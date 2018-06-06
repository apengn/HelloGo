package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var timeout = time.Duration(5 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

//通过url 获取数据
func Fetcher(url string) ([]byte, error) {

	t := http.Transport{
		Dial:                  dialTimeout,
		MaxIdleConns:          10,
		IdleConnTimeout:       timeout,
		ResponseHeaderTimeout: timeout}

	client := http.Client{Transport: &t}
	resp, err := client.Get(url)


	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("http status:", resp.StatusCode)
		return nil, errors.New(fmt.Sprintln("http status:", resp.StatusCode))
	}
	bufioRead := bufio.NewReader(resp.Body)
	e := determineEncoding(bufioRead)
	//文本格式转换
	transformReader := transform.NewReader(bufioRead, e.NewDecoder())
	//把reader转换为字节数组
	bytes, err := ioutil.ReadAll(transformReader)

	return bytes, err
}

//判断html的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//取出网页中1024个字节用来猜测字符编码
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("guess decode failed,err:%s", err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
