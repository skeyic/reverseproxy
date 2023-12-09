package fowarder

import (
	"github.com/vulcand/oxy/forward"
)

var (
	TheDirectForwarder, _ = forward.New()
)
