package table_test

import (
	_ "embed"
	"testing"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/darabuchi/utils/table"
)

func TestTable(t *testing.T) {
	tb := table.NewTable().AddLine(table.NewLine().SetText("这是一行"))

	b, err := tb.ToImg()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = utils.FileWrite("test.png", b.String())
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
}
