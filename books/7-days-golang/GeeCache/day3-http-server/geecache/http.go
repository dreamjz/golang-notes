package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/geecache/"

var _ http.Handler = (*HTTPPool)(nil)

// HTTPPool implements PeerPicker for a pool of HTTP peers
type HTTPPool struct {
	self     string // 自己的地址，主机+端口
	basePath string // 节点间通讯地址的前缀
}

func NewHTTPPool(addr string) *HTTPPool {
	return &HTTPPool{
		self:     addr,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...any) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if !strings.HasPrefix(path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + path)
	}
	p.Log("%s %s", req.Method, path)

	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}
