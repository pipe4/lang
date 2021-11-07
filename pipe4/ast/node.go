package ast

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func (l NodeList) ToYaml() (string, error) {
	buf, err := json.Marshal(l)
	if err != nil {
		return "", fmt.Errorf("failed print ast tree: %w", err)
	}
	var m []interface{}
	err = json.Unmarshal(buf, &m)
	if err != nil {
		return "", fmt.Errorf("failed print ast tree: %w", err)
	}
	out := &strings.Builder{}
	yamlEncoder := yaml.NewEncoder(out)
	yamlEncoder.SetIndent(2)

	if err := yamlEncoder.Encode(m); err != nil {
		return "", fmt.Errorf("failed print ast tree: %w", err)
	}
	str := out.String()
	return str, nil
}
func (l *NodeList) FromYaml(source string) error {
	var m []interface{}
	if err := yaml.Unmarshal([]byte(source), &m); err != nil {
		return fmt.Errorf("failed to read ast from json: yaml.Unmarshal: %w", err)
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to read ast from json: json.Marshal: %w", err)
	}

	err = json.Unmarshal(buf, l)
	if err != nil {
		return fmt.Errorf("failed to read ast from json: json.Unmarshal: %w", err)
	}
	return nil

}
func (l NodeList) Get(name string) (Node, bool) {
	for _, node := range l {
		if node.Ident.Name == name {
			return node, true
		}
	}
	return Node{}, false
}
func (l NodeList) GetString(name string) string {
	if node, ok := l.Get(name); ok && node.Type.BodyOneOf == BodyString {
		return node.Type.String
	}
	return ""
}

func (t Node) MarshalJSON() ([]byte, error) {
	s := map[string]interface{}{}

	if t.Ident.Name != "" {
		s["Ident"] = t.Ident.Name
	}
	if !t.Comment.IsEmpty() {
		s["Comment"] = t.Comment
	}
	if !t.Type.IsEmpty() {
		s["Type"] = t.Type
	}
	if !t.Default.IsEmpty() {
		s["Default"] = t.Default
	}

	return json.Marshal(s)
}

func (t *Node) UnmarshalJSON(data []byte) error {
	*t = Node{}
	s := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("failed to unmarshal Node from %s: %w", data, err)
	}
	if v, ok := s["Ident"]; ok {
		if err := json.Unmarshal(v, &t.Ident.Name); err != nil {
			return fmt.Errorf("failed to unmarshal Ident from %s: %w", v, err)
		}
	}
	if v, ok := s["Comment"]; ok {
		if err := json.Unmarshal(v, &t.Comment); err != nil {
			return fmt.Errorf("failed to unmarshal Comment from %s: %w", v, err)
		}
	}
	if v, ok := s["Type"]; ok {
		if err := json.Unmarshal(v, &t.Type); err != nil {
			return fmt.Errorf("failed to unmarshal Type from %s: %w", v, err)
		}
	}
	if v, ok := s["Default"]; ok {
		if err := json.Unmarshal(v, &t.Default); err != nil {
			return fmt.Errorf("failed to unmarshal Default from %s: %w", v, err)
		}
	}

	return nil
}

func (t Comment) IsEmpty() bool {
	return t.Text == "" && len(t.Tags) == 0
}
