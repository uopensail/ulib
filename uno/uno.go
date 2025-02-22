package uno

import (
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/uopensail/ulib/sample"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
)

type _Column struct {
	Addr   int32
	Column string
	Type   sample.DataType
}

type _Columns struct {
	addrs []int32
	cols  []string
	types []sample.DataType
}

type _Expression struct {
	columns  _Columns
	varSlice []unsafe.Pointer
}

type Evaluator struct {
	expression unsafe.Pointer
	columns    map[string]map[string]_Column
	types      map[string]sample.DataType
}

func NewEvaluator(condition string, types map[string]sample.DataType) (*Evaluator, error) {
	err := check(condition)
	if err != nil {
		return nil, err
	}
	code, err := parse(condition)
	if err != nil {
		return nil, err
	}

	expr := uno_create_expression(code)
	columns := make(map[string]map[string]_Column)
	cols := (*_Expression)(expr).columns
	for i := 0; i < len(cols.addrs); i++ {
		items := strings.Split(cols.cols[i], ".")
		tmp, ok := columns[items[0]]
		if !ok {
			tmp = make(map[string]_Column)
		}
		tmp[items[1]] = _Column{
			Addr:   cols.addrs[i],
			Column: cols.cols[i],
			Type:   cols.types[i],
		}
		columns[items[0]] = tmp
	}
	return &Evaluator{expression: expr, columns: columns}, nil
}

func (e *Evaluator) Fill(table string, features sample.Features, value []unsafe.Pointer) {
	columns, ok := e.columns[table]
	if !ok {
		zlog.LOG.Error("table not exists", zap.String("table", table))
		return
	}
	for name, column := range columns {
		fea := features.Get(name)
		if fea == nil {
			continue
		}

		if fea.Type() != column.Type {
			zlog.LOG.Error("column type mismatch")
			continue
		}
		if column.Type == sample.Int64Type {
			data, _ := fea.GetInt64()
			value[column.Addr] = unsafe.Pointer(&data)
		} else if column.Type == sample.Int64sType {
			data, _ := fea.GetInt64s()
			value[column.Addr] = unsafe.Pointer(&data)
		} else if column.Type == sample.Float32Type {
			data, _ := fea.GetFloat32()
			value[column.Addr] = unsafe.Pointer(&data)
		} else if column.Type == sample.Float32sType {
			data, _ := fea.GetFloat32s()
			value[column.Addr] = unsafe.Pointer(&data)
		} else if column.Type == sample.StringType {
			data, _ := fea.GetString()
			value[column.Addr] = unsafe.Pointer(&data)
		} else if column.Type == sample.StringsType {
			data, _ := fea.GetStrings()
			value[column.Addr] = unsafe.Pointer(&data)
		}
	}
}

func (e *Evaluator) FillInt64(table, col string, data int64, value []unsafe.Pointer) error {
	columns, ok := e.columns[table]
	if !ok {
		return fmt.Errorf("table %s not found", table)
	}

	column, ok := columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == sample.Int64Type {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillInt64s(table, col string, data []int64, value []unsafe.Pointer) error {
	columns, ok := e.columns[table]
	if !ok {
		return fmt.Errorf("table %s not found", table)
	}
	column, ok := columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == sample.Int64sType {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillFloat32(table, col string, data float32, value []unsafe.Pointer) error {
	columns, ok := e.columns[table]
	if !ok {
		return fmt.Errorf("table %s not found", table)
	}
	column, ok := columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == sample.Float32Type {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillFloat32s(table, col string, data []float32, value []unsafe.Pointer) error {
	columns, ok := e.columns[table]
	if !ok {
		return fmt.Errorf("table %s not found", table)
	}
	column, ok := columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == sample.Float32sType {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillString(table, col string, data string, value []unsafe.Pointer) error {
	columns, ok := e.columns[table]
	if !ok {
		return fmt.Errorf("table %s not found", table)
	}
	column, ok := columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == sample.StringType {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillStrings(table, col string, data []string, value []unsafe.Pointer) error {
	columns, ok := e.columns[table]
	if !ok {
		return fmt.Errorf("table %s not found", table)
	}
	column, ok := columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == sample.StringsType {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) Allocate() []unsafe.Pointer {
	ret := make([]unsafe.Pointer, len((*_Expression)(e.expression).varSlice))
	copy(ret, (*_Expression)(e.expression).varSlice)
	return ret
}

func (e *Evaluator) Eval(slice []unsafe.Pointer) int32 {
	return uno_eval(e.expression, slice)
}

func (e *Evaluator) PreEval(slice []unsafe.Pointer) {
	uno_preeval(e.expression, slice)
}

func (e *Evaluator) BatchEval(slices [][]unsafe.Pointer) []int32 {
	return uno_batch_eval(e.expression, slices)
}

func (e *Evaluator) Clean(slice []unsafe.Pointer) {
	uno_clean_varslice(e.expression, slice)
}

func (e *Evaluator) Release() {
	uno_release_expression(e.expression)
}

func check(condition string) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf(perr.(string))
		}
	}()
	s := antlr.NewInputStream(condition)
	lexer := NewunoLexer(s)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewunoParser(tokens)
	parser.Start_()
	return
}

func parse(condition string, types map[string]sample.DataType) (code string, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf(perr.(string))
		}
	}()
	s := antlr.NewInputStream(condition)
	lexer := NewunoLexer(s)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewunoParser(tokens)
	listener := NewListener(types)
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Start())
	root := listener.booleans.Pop()
	root = root.Simplify()
	nodes := root.ToList()
	for i := 0; i < len(nodes); i++ {
		nodes[i].SetId(int32(i))
	}

	type jsonNodes struct {
		Nodes []Expression `json:"nodes"`
	}
	bytes, jerr := json.Marshal(&jsonNodes{Nodes: nodes})

	if jerr != nil {
		err = jerr
		code = ""
		return
	}
	code = string(bytes)
	return
}
