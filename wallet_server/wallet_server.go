package walletserver

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/samcheck/go-blockchain/wallet"
)

const TEMPLATE_DIR = "wallet_server/templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "text/html")
		t, _ := template.ParseFiles(path.Join(TEMPLATE_DIR, "index.html"))
		t.Execute(w, "")

	default:
		log.Printf(("ERROR: Invalid HTTP METHOD"))
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf(("ERROR: Invalid HTTP METHOD"))
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
