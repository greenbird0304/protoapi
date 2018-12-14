package phpyii2

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/yoozoo/protoapi/util"

	"github.com/yoozoo/protoapi/generator/data"
)

// NewMessage return a pointer of new Message struct
func NewMessage(msg *data.MessageData, baseNameSpace string, enums []*data.EnumData) *Message {
	nameSpace := baseNameSpace + "\\models"
	filePath := strings.Replace(nameSpace, "\\", "/", -1)

	// todo: now only handle msg in current package
	// should handle import from different packages
	_, p := data.GetMessageProtoAndFile(msg.Name)
	if p.IsFileToGenerate {
		msg.Name = msg.Name[strings.LastIndex(msg.Name, ".")+1:]
	}

	for _, f := range msg.Fields {
		_, p = data.GetMessageProtoAndFile(f.DataType)
		if p.IsFileToGenerate {
			f.DataType = f.DataType[strings.LastIndex(f.DataType, ".")+1:]
		}
	}

	filePath = filePath + "/" + strings.Title(msg.Name) + ".php"

	o := &Message{msg, nameSpace, filePath, enums}
	return o
}

// Message is struct of php message class
type Message struct {
	*data.MessageData
	NameSpace string
	FilePath  string
	Enums     []*data.EnumData
}

func (p *Message) IsObject(fieldType string) bool {
	switch fieldType {
	case data.StringFieldType,
		data.DoubleFieldType,
		data.IntFieldType,
		data.BooleanFieldType:
		return false
	default:
		// check if is enum
		for _, enum := range p.Enums {
			if enum.Name == fieldType {
				return false
			}
		}
		return true
	}
}

func (p *Message) Gen(result map[string]string) error {
	buf := bytes.NewBufferString("")

	tplContent := data.LoadTpl("/generator/template/yii2/models/message.gophp")

	funcMap := template.FuncMap{
		"isObject":  p.IsObject,
		"className": util.GetPHPClassName,
	}

	tpl, err := template.New("message").Funcs(funcMap).Parse(tplContent)
	if err != nil {
		return err
	}
	err = tpl.Execute(buf, p)
	if err != nil {
		return err
	}

	result[p.FilePath] = buf.String()
	return nil
}
