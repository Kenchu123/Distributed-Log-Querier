package constant

import "time"

var CONFIG_PATH string = ".dsgrep/config.yml"

// socket constants
var BUFFER_SIZE int = 16777216
var CHUNK_SIZE int = 4096
var CONNECTION_TIMEOUT time.Duration = 5 * time.Second
