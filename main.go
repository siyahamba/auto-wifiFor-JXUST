package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func connectToWifi(ssid string) bool {
	cmd := exec.Command("netsh", "wlan", "connect", "name=\""+ssid+"\"")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to connect to WiFi:", stderr.String())
		return false
	}
	fmt.Println("Successfully connected to WiFi.")
	return true
}

func ping(host string) bool {
	cmd := exec.Command("ping", "-n", "1", host) // Windows ping
	// For Unix/Linux, use: cmd := exec.Command("ping", "-c", "1", host)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}
	// Simple check: if "TTL=" is found, assume ping was successful
	return strings.Contains(out.String(), "TTL=")
}

func con(filename string) string {
	url := "http://eportal.jxust.edu.cn:801/eportal/portal/login"
	uname, passwd, err := readfile(filename)
	params := map[string]string{
		"callback":      "dr1003",
		"login_method":  "1",
		"user_account":  uname,
		"user_password": passwd,
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Host", "eportal.jxust.edu.cn:801")
	req.Header.Set("Referer", "http://eportal.jxust.edu.cn/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}
	fmt.Println("wifi gets up")
	return string(body)
}

type UserInput struct {
	ID      string
	Passwd  string
	Contact int // 假设这个字段用于存储联通方式
}

func initInput(input UserInput) UserInput {
	d := map[int]string{}
	d[1] = "telecom"
	d[2] = "cmcc"
	d[3] = "unicom"
	input.ID = input.ID + "@" + d[input.Contact]
	return input
}

func Input() (UserInput, error) {
	var id string
	var passwd string
	var types int

	for {
		fmt.Print("请输入你的一卡通号：")
		if _, err := fmt.Scan(&id); err != nil {
			return UserInput{}, err // 返回错误（如果输入无法解析为整数）
		}

		fmt.Print("请输入你的登录密码：")
		fmt.Scan(&passwd) // 注意：这里假设密码不包含空格等特殊字符

		fmt.Print("请输入你的联通方式（1表示电信，2表示移动，3表示联通）：")
		if _, err := fmt.Scan(&types); err != nil {
			return UserInput{}, err // 返回错误（如果输入无法解析为整数）
		}

		// 检查types是否在有效范围内
		if types == 1 || types == 2 || types == 3 {
			break // 跳出循环
		}

		fmt.Println("输入不合法请重新输入（1表示电信，2表示移动，3表示联通）：")
	}

	return UserInput{ID: id, Passwd: passwd, Contact: types}, nil
}

func checkfile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			userMess, err := Input()
			if err != nil {
				return err
			}
			userMess = initInput(userMess)
			f, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = fmt.Fprintf(f, "%s\n%s", userMess.ID, userMess.Passwd)
			return err
		}
		return err
	}
	return nil
}

func readfile(filename string) (string, string, error) {
	err := checkfile(filename)
	if err != nil {
		return "", "", err
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", "", fmt.Errorf("打开文件出错: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var id, passwd string
	if scanner.Scan() {
		id = scanner.Text()
	} else if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("读取账号时出错: %w", err)
	}

	if scanner.Scan() {
		passwd = scanner.Text()
	} else if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("读取密码时出错: %w", err)
	}

	return id, passwd, nil
}

func main() {
	filename := "config.txt"
	ssid := "JXUST-WLAN"
	if connectToWifi(ssid) {
		time.Sleep(2 * time.Second)
		con(filename)
		if ping("www.baidu.com") {
			fmt.Println("wifi is up.")
		} else {
			con(filename)
		}
	}
}
