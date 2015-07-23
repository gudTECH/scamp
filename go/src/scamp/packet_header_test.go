package scamp

import "testing"
import "bytes"
import "encoding/json"

func TestEncodeEnvelope(t *testing.T){
  expected := []byte("\"JSON\"\n")

  buf := new(bytes.Buffer)
  encoder := json.NewEncoder(buf)

  val := ENVELOPE_JSON
  err := encoder.Encode(val)

  if err != nil {
    t.Errorf("got unexpected err `%s`\n", err)
    t.FailNow()
  }
  if !bytes.Equal(expected, buf.Bytes()){
    t.Errorf("expected `%s` but got `%s`", expected, buf.Bytes())
    t.FailNow()
  }
}

func TestWritePacketHeader(t *testing.T){
  packetHeader := PacketHeader {
    Action: "hello.helloworld",
    Envelope: ENVELOPE_JSON,
    MessageId: "0123456789012345",
    Version: 1,
  }
  expected := []byte("{\"action\":\"hello.helloworld\",\"envelope\":\"JSON\",\"message_id\":\"0123456789012345\",\"type\":\"REQUEST\",\"version\":1}\n")

  buf := new(bytes.Buffer)
  packetHeader.Write(buf)

  if !bytes.Equal(expected, buf.Bytes()) {
    t.Errorf("expected\n`%s`\n`%v`\ngot\n`%s`\n`%v`\n", expected, expected, buf.Bytes(), buf.Bytes())
  }
}

