// Code generated by Thrift Compiler (0.14.2). DO NOT EDIT.

package zipkincore

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

const CLIENT_SEND = "cs"
const CLIENT_RECV = "cr"
const SERVER_SEND = "ss"
const SERVER_RECV = "sr"
const MESSAGE_SEND = "ms"
const MESSAGE_RECV = "mr"
const WIRE_SEND = "ws"
const WIRE_RECV = "wr"
const CLIENT_SEND_FRAGMENT = "csf"
const CLIENT_RECV_FRAGMENT = "crf"
const SERVER_SEND_FRAGMENT = "ssf"
const SERVER_RECV_FRAGMENT = "srf"
const HTTP_HOST = "http.host"
const HTTP_METHOD = "http.method"
const HTTP_PATH = "http.path"
const HTTP_ROUTE = "http.route"
const HTTP_URL = "http.url"
const HTTP_STATUS_CODE = "http.status_code"
const HTTP_REQUEST_SIZE = "http.request.size"
const HTTP_RESPONSE_SIZE = "http.response.size"
const LOCAL_COMPONENT = "lc"
const ERROR = "error"
const CLIENT_ADDR = "ca"
const SERVER_ADDR = "sa"
const MESSAGE_ADDR = "ma"

func init() {
}
