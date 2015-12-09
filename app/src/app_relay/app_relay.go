/**
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Jane Lee <jane@webconn.me>
 * Copyright (c) 2015 Edward Kim <edward@webconn.me>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
    "github.com/webconnme/go-webconn"
)

import (
	"fmt"
	"log"
	"ioctl"
	"unsafe"
	"os"
)

type GpioS struct {
	io int
	mode int
	value int
}

const (
	IOCTL_GPIO_SET_OUTPUT = 7239681
	IOCTL_GPIO_GET_OUTPUT = 7239682
)

var client webconn.Webconn

var file *os.File

func GPIOOpen() {
	var err error
	file, err = os.OpenFile("/dev/ioctl_gpio", os.O_RDWR | os.O_SYNC, 0777)
	if err != nil {
		log.Fatal("open", err)
	}
	log.Println("GPIO Open...");

}

func GPIOClose() {
	file.Close()
	fmt.Println("GPIO Close...")
}

func GpioOut(buf []byte) error {
	g := GpioS {28, 0, 0}
	header := unsafe.Pointer(&g)

    state := string(buf)

    if state == "on" {
        g.value = 1
    }else if state == "off" {
        g.value = 0
    }

    ioctl.IOCTL(uintptr(file.Fd()), IOCTL_GPIO_SET_OUTPUT, uintptr(header))

    fmt.Println(">>>power stat : ", g.value)

    return nil
}

func main() {
	GPIOOpen()
    defer GPIOClose()

    client = webconn.NewClient("http://nor.kr:3002/v01/relay/80")
    client.AddHandler("power", GpioOut)

    client.Run()
}
