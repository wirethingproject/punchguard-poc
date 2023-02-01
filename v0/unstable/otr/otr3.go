package otr

import (
	"crypto/rand"
	"log"

	"github.com/coyim/otr3"

	"github.com/punchguard/v0"
)

type Otr3 struct {
	punchguard.BaseOtr
	Conversation *otr3.Conversation
}

func NewConversation() *otr3.Conversation {
	priv := new(otr3.DSAPrivateKey)
	priv.Generate(rand.Reader)

	c := new(otr3.Conversation)
	c.SetOurKeys([]otr3.PrivateKey{priv})

	// set the Policies.
	c.Policies.RequireEncryption()
	// c.Policies.AllowV2()
	c.Policies.AllowV3()
	// c.Policies.SendWhitespaceTag()
	// c.Policies.WhitespaceStartAKE()

	// You can also setup a debug mode
	c.SetDebug(true)

	// TODO Authentication
	return c
}

func (o *Otr3) Init(id string) error {
	// log.Printf("%T.Init", o)

	if err := o.InitBase(id); err != nil {
		return err
	}

	o.Conversation = NewConversation()

	return nil
}

func (o *Otr3) Query() ([]string, error) {
	var result []string

	if o.GetId() != "o1" {
		return result, nil
	}

	return o.Encode("")
}

func (o *Otr3) End() ([]string, error) {
	var result []string
	toSend, err := o.Conversation.End()
	if err != nil {
		log.Printf("%T.Hi: error %v", o, err)
	}
	for _, s := range toSend {
		result = append(result, string(s))
	}
	return result, nil
}

func (o *Otr3) Encode(plain string) ([]string, error) {
	var result []string

	// log.Printf("%T.Encode: %v pre IsEncrypted '%v' '%v'", o, o.GetId(), o.Conversation.IsEncrypted(), plain)
	toSend, err := o.Conversation.Send([]byte(plain))

	// log.Printf("%T.Encode: %v pos IsEncrypted '%v' len %v", o, o.GetId(), o.Conversation.IsEncrypted(), len(toSend))

	if !o.Conversation.IsEncrypted() && o.IsReady {
		o.IsReady = false
	}

	if err != nil {
		log.Printf("%T.Encode: error %v", o, err)
	}

	for _, s := range toSend {
		result = append(result, string(s))
	}

	return result, nil
}

func (o *Otr3) Decode(msg string) (string, []string, error) {
	var result []string

	// log.Printf("%T.Decode: %v pre IsEncrypted '%v'", o, o.GetId(), o.Conversation.IsEncrypted())
	plain, toSend, err := o.Conversation.Receive([]byte(msg))
	// log.Printf("%T.Decode: %v pos IsEncrypted '%v' '%v' len %v", o, o.GetId(), o.Conversation.IsEncrypted(), plain, len(toSend))

	if o.Conversation.IsEncrypted() && !o.IsReady {
		o.Ready()
	}

	if !o.Conversation.IsEncrypted() && o.IsReady {
		o.IsReady = false
	}

	if err != nil {
		log.Printf("%T.Encode: error %v", o, err)
	}

	for _, s := range toSend {
		result = append(result, string(s))
	}

	return string(plain), result, nil
}
