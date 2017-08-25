// http://ascii.jp/elem/000/001/243/1243667/
// 第2回　低レベルアクセスへの入り口（1）：io.Writer
package main

//import "os"
import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)


func main() {
	/* ---ファイル書き込みの処理--- */
//	file, err := os.Create("test.txt")
//	if err != nil {
//		panic(err)
//	}
//	file.Write([]byte("os.File example\n"))
//	file.Close()
	/* ---ファイル書き込みの処理--- */

	/* ---画面出力の処理--- */
	//os.Stdout.Write([]byte("os.Stdout example\n"))
	/* ---画面出力の処理--- */

	/* ---バッファを利用して内容を記憶するときもio.Writerを使う--- */
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.buffer example\n"))
	//buffer.WriteString("bytes.Buffer example\n") // ←のようにすればキャストは不要
	//io.WriteString(buffer, "bytes.Buffer example\n") // これでも出来るらしいがbufferの型が合わない
	fmt.Println(buffer.String())
	/* ---バッファを利用して内容を記憶するときもio.Writerを使う--- */

	/* ---インターネットアクセス--- */
	// net.Dialが返すnet.Connにはio.Writerが実装されている
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
	/* ---インターネットアクセス--- */


}