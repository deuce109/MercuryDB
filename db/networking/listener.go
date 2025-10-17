package networking

type Listener interface {
	Listen()
	Close()
}

var listeners = []Listener{}

func StartListeners() {
	for _, l := range listeners {
		if l != nil {
			go l.Listen()
		}
	}
}

func CloseListeners() {
	for _, l := range listeners {
		if l != nil {
			l.Close()
		}
	}
}
