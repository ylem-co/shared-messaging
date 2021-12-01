# Datamin messages
## Module contents
The module contains messages exchanged within the whole Datamin system and methods for marshaling/unmarshaling.

## Installation
In order to include this module in your application, add the following to your .bashrc and .bash_profile:

```bash
export GOPRIVATE=github.com/datamin-io
```

Also, add the following to your ~/.gitconfig:
```bash
[url "git@github.com:"]
        insteadOf = https://github.com/
```

Download the module:

```bash
$ go get github.com/datamin-io/messaging@main
```

Then import as usual:

```golang
import "github.com/datamin-io/messaging"
```


## Usage
### On the sending end
```golang
package main

import "github.com/datamin-io/messaging"

// create an actual message, e.g.
msg := &messaging.RunQueryTask{
    Query: "SELECT * FROM some_table ORDER BY id DESC LIMIT 100"
}

// wrap it into an envelope
envelope := messaging.NewEnvelope(msg);

// add headers with meta-information if necessary
// note the assigment — WithHeader() returns a new instance
envelope = envelope.WithHeader("X-Load-Balancer-Cost", "100")

// serialize it before sending to a transport
codec := &messaging.MessageCodec{}
serialized, err := codec.Encode(envelope)

// finally send it
sometransport.Send(serialized)
```

### On the receiving end

```golang
package main

import (
    "fmt"

    "github.com/datamin-io/messaging"
)

// receive a serialized message from the transport
serialized, err := sometransport.Receive()

// deserialize it
codec := &messaging.MessageCodec{}
envelope, err := codec.Decode(e)

// determine type of the wrapped message by checking "X-Message-Name" header, which is added automatically on the sending side

switch envelope.Headers[messaging.HEADER_MESSAGE_NAME] {
    case messaging.RUN_QUERY_TASK:
        msg := envelope.Msg.(*RunQueryTask)
        fmt.Printf(msg.Query)
}

```


### Using together with lovoo/goka
https://github.com/lovoo/goka

Goka is a compact yet powerful distributed stream processing library for Apache Kafka written in Go.

`MessageCodec` struct implements [Goka's Codec interface](https://github.com/lovoo/goka/blob/master/codec.go),
therefore this module can be easily integrated with Goka as follows:

```golang
package main

import (
	"fmt"

	"github.com/datamin-io/messaging"
	"github.com/lovoo/goka"
)

func main() {
    emitter, err := goka.NewEmitter("127.0.0.1", goka.Stream("some-topic"), new(messaging.MessageCodec))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	defer emitter.Finish()

    envelope := messaging.NewEnvelope(
        &messaging.RunQueryTask{
            Query: "SELECT * FROM some_table ORDER BY id DESC LIMIT 100"
        },
    );

    // the message will get serialized automatically and sent to the topic "some-topic"
    err = emitter.EmitSync("", envelope)
    if err != nil {
        log.Fatalf("error emitting message: %v", err)
    }

    
}
```
