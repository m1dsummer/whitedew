package whitedew

import (
	"bytes"
	"fmt"
)

type MessageChain struct {
	buffer bytes.Buffer
}

func (m *MessageChain) Prepare() *MessageChain {
	m.buffer = bytes.Buffer{}
	return m
}

// Plain
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%BA%AF%E6%96%87%E6%9C%AC
func (m *MessageChain) Plain(str string) *MessageChain {
	m.buffer.WriteString(str)
	return m
}

// Face
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#qq-%E8%A1%A8%E6%83%85
func (m *MessageChain) Face(faceId int) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:face,id=%d]", faceId))
	return m
}

// Image
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%9B%BE%E7%89%87
func (m *MessageChain) Image(url string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:image,file=%s]", url))
	return m
}

// Record
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E8%AF%AD%E9%9F%B3
func (m *MessageChain) Record(url string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:record,file=%s]", url))
	return m
}

// Video
//
// https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%9F%AD%E8%A7%86%E9%A2%91
func (m *MessageChain) Video(url string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:video,file=%s]", url))
	return m
}

// At
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E6%9F%90%E4%BA%BA
func (m *MessageChain) At(uid int64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:at,qq=%d]", uid))
	return m
}

// Rps
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%8C%9C%E6%8B%B3%E9%AD%94%E6%B3%95%E8%A1%A8%E6%83%85
func (m *MessageChain) Rps() *MessageChain {
	m.buffer.WriteString("[CQ:rps]")
	return m
}

// Dice
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E6%8E%B7%E9%AA%B0%E5%AD%90%E9%AD%94%E6%B3%95%E8%A1%A8%E6%83%85
func (m *MessageChain) Dice() *MessageChain {
	m.buffer.WriteString("[CQ:dice]")
	return m
}

// Shake
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E7%AA%97%E5%8F%A3%E6%8A%96%E5%8A%A8%E6%88%B3%E4%B8%80%E6%88%B3-
func (m *MessageChain) Shake() *MessageChain {
	m.buffer.WriteString("[CQ:shake]")
	return m
}

// Poke
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E6%88%B3%E4%B8%80%E6%88%B3
func (m *MessageChain) Poke(_type int, id int64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:poke,type=%d,id=%d]", _type, id))
	return m
}

func (m *MessageChain) AtAll() *MessageChain {
	m.buffer.WriteString("[CQ:at,qq=all]")
	return m
}

// Anonymous
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%8C%BF%E5%90%8D%E5%8F%91%E6%B6%88%E6%81%AF-
func (m *MessageChain) Anonymous() *MessageChain {
	m.buffer.WriteString("[CQ:anonymous]")
	return m
}

// Share
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E9%93%BE%E6%8E%A5%E5%88%86%E4%BA%AB
func (m *MessageChain) Share(title string, url string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:share,url=%s,title=%s]", url, title))
	return m
}

// Contact
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E6%8E%A8%E8%8D%90%E5%A5%BD%E5%8F%8B
func (m *MessageChain) Contact(_type string, uid int64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:contact,type=%s,id=%d]", _type, uid))
	return m
}

// Location
//
// https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E4%BD%8D%E7%BD%AE
func (m *MessageChain) Location(lat float64, lon float64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:location,lat=%f,lon=%f]", lat, lon))
	return m
}

// Music
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E9%9F%B3%E4%B9%90%E5%88%86%E4%BA%AB-
func (m *MessageChain) Music(_type string, id int64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:music,type=%s,id=%d]", _type, id))
	return m
}

// CustomMusic
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E9%9F%B3%E4%B9%90%E8%87%AA%E5%AE%9A%E4%B9%89%E5%88%86%E4%BA%AB-
func (m *MessageChain) CustomMusic(url, audio, title string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:music,type=custom,url=%s,audio=%s,title=%s]", url, audio, title))
	return m
}

// Reply
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%9B%9E%E5%A4%8D
func (m *MessageChain) Reply(mid int64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:reply,id=%d]", mid))
	return m
}

// Forward
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%90%88%E5%B9%B6%E8%BD%AC%E5%8F%91-
func (m *MessageChain) Forward(mid int64) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:forward,id=%d]", mid))
	return m
}

// TODO: implement Node method
// Node
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#%E5%90%88%E5%B9%B6%E8%BD%AC%E5%8F%91%E8%87%AA%E5%AE%9A%E4%B9%89%E8%8A%82%E7%82%B9
//func (m *MessageChain)Node() {
//
//}

// XML
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#xml-%E6%B6%88%E6%81%AF
func (m *MessageChain) XML(msg string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:xml,data=%s]", msg))
	return m
}

// JSON
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md#json-%E6%B6%88%E6%81%AF
func (m *MessageChain) JSON(msg string) *MessageChain {
	m.buffer.WriteString(fmt.Sprintf("[CQ:json,data=%s", msg))
	return m
}

func (m MessageChain) String() string {
	return m.buffer.String()
}
