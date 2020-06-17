package goroseadapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type Adapter struct {
	Client *Client
}

func NewAdapter() *Adapter {
	return &Adapter{
		Client: NewFromEnvionment(),
	}
}

func loadPolicyLine(line *CasbinRule, model model.Model) {
	const prefixLine = ", "
	var sb strings.Builder

	sb.WriteString(line.P_type)
	if len(line.V0) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(line.V0)
	}
	if len(line.V1) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(line.V1)
	}
	if len(line.V2) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(line.V2)
	}
	if len(line.V3) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(line.V3)
	}
	if len(line.V4) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(line.V4)
	}
	if len(line.V5) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(line.V5)
	}

	persist.LoadPolicyLine(sb.String(), model)
}

func (a *Adapter) LoadPolicy(model model.Model) error {
	retrieve, err := a.Client.Retrieve(&CasbinRule{})
	if err != nil {
		return err
	}
	lines := make([]*CasbinRule, 0)
	err = json.Unmarshal(retrieve, &lines)
	if err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return nil
}

func (a *Adapter) SavePolicy(model model.Model) error {
	return nil
}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := &CasbinRule{
		P_type: ptype,
	}
	l := len(rule)
	if l > 0 {
		line.V0 = rule[0]
	}
	if l > 1 {
		line.V1 = rule[1]
	}
	if l > 2 {
		line.V2 = rule[2]
	}
	if l > 3 {
		line.V3 = rule[3]
	}
	if l > 4 {
		line.V4 = rule[4]
	}
	if l > 5 {
		line.V5 = rule[5]
	}
	retrieve, err := a.Client.Create(line)
	if err != nil {
		return err
	}
	fmt.Println(string(retrieve))
	return nil
}

func savePolicyLine(ptype string, rule []string) *CasbinRule {
	line := &CasbinRule{P_type: ptype}

	l := len(rule)
	if l > 0 {
		line.V0 = rule[0]
	}
	if l > 1 {
		line.V1 = rule[1]
	}
	if l > 2 {
		line.V2 = rule[2]
	}
	if l > 3 {
		line.V3 = rule[3]
	}
	if l > 4 {
		line.V4 = rule[4]
	}
	if l > 5 {
		line.V5 = rule[5]
	}

	return line
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	retrieve, err := a.Client.Retrieve(line)
	if err != nil {
		return err
	}
	lines := make([]*CasbinRule, 0)
	err = json.Unmarshal(retrieve, &lines)
	if err != nil {
		return err
	}
	if len(lines) == 0 {
		return nil
	}
	const prefixLine = ","
	var sb strings.Builder
	for _, line := range lines {
		sb.WriteString(fmt.Sprintf("%d", line.Id))
		sb.WriteString(prefixLine)
	}
	bytes, err := a.Client.Delete(sb.String())
	fmt.Println(string(bytes))
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
