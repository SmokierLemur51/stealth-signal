package data

/*
We will be working under the assumption that messages are only intended for specific
recipients. Each group will have members, some can post messages, some can only read.

Each message will be specifically encrypted with the public key of the recipient in mind.
This will take extra computing power, but it will be the best way to get the job done.

So a member will create a message, that message will then be sent to the server to be processed
into a signed/encrypted messsage for each group member it was meant for. It will then wait for
a request on that group/members endpoint.
*/

// "github.com/ProtonMail/gopenpgp/v2"

func CreateKeyPair() error { return nil }
