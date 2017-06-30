// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domsnapshot

import (
	json "encoding/json"
	css "github.com/knq/chromedp/cdp/css"
	dom "github.com/knq/chromedp/cdp/dom"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot(in *jlexer.Lexer, out *NameValue) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot(out *jwriter.Writer, in NameValue) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Value != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"value\":")
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NameValue) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NameValue) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NameValue) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NameValue) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot1(in *jlexer.Lexer, out *LayoutTreeNode) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "domNodeIndex":
			out.DomNodeIndex = int64(in.Int64())
		case "boundingBox":
			if in.IsNull() {
				in.Skip()
				out.BoundingBox = nil
			} else {
				if out.BoundingBox == nil {
					out.BoundingBox = new(dom.Rect)
				}
				(*out.BoundingBox).UnmarshalEasyJSON(in)
			}
		case "layoutText":
			out.LayoutText = string(in.String())
		case "inlineTextNodes":
			if in.IsNull() {
				in.Skip()
				out.InlineTextNodes = nil
			} else {
				in.Delim('[')
				if out.InlineTextNodes == nil {
					if !in.IsDelim(']') {
						out.InlineTextNodes = make([]*css.InlineTextBox, 0, 8)
					} else {
						out.InlineTextNodes = []*css.InlineTextBox{}
					}
				} else {
					out.InlineTextNodes = (out.InlineTextNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *css.InlineTextBox
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(css.InlineTextBox)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.InlineTextNodes = append(out.InlineTextNodes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "styleIndex":
			out.StyleIndex = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot1(out *jwriter.Writer, in LayoutTreeNode) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DomNodeIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"domNodeIndex\":")
		out.Int64(int64(in.DomNodeIndex))
	}
	if in.BoundingBox != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"boundingBox\":")
		if in.BoundingBox == nil {
			out.RawString("null")
		} else {
			(*in.BoundingBox).MarshalEasyJSON(out)
		}
	}
	if in.LayoutText != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"layoutText\":")
		out.String(string(in.LayoutText))
	}
	if len(in.InlineTextNodes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"inlineTextNodes\":")
		if in.InlineTextNodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.InlineTextNodes {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.StyleIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"styleIndex\":")
		out.Int64(int64(in.StyleIndex))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LayoutTreeNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LayoutTreeNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LayoutTreeNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LayoutTreeNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot2(in *jlexer.Lexer, out *GetSnapshotReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "domNodes":
			if in.IsNull() {
				in.Skip()
				out.DomNodes = nil
			} else {
				in.Delim('[')
				if out.DomNodes == nil {
					if !in.IsDelim(']') {
						out.DomNodes = make([]*DOMNode, 0, 8)
					} else {
						out.DomNodes = []*DOMNode{}
					}
				} else {
					out.DomNodes = (out.DomNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *DOMNode
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(DOMNode)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.DomNodes = append(out.DomNodes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "layoutTreeNodes":
			if in.IsNull() {
				in.Skip()
				out.LayoutTreeNodes = nil
			} else {
				in.Delim('[')
				if out.LayoutTreeNodes == nil {
					if !in.IsDelim(']') {
						out.LayoutTreeNodes = make([]*LayoutTreeNode, 0, 8)
					} else {
						out.LayoutTreeNodes = []*LayoutTreeNode{}
					}
				} else {
					out.LayoutTreeNodes = (out.LayoutTreeNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *LayoutTreeNode
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(LayoutTreeNode)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.LayoutTreeNodes = append(out.LayoutTreeNodes, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "computedStyles":
			if in.IsNull() {
				in.Skip()
				out.ComputedStyles = nil
			} else {
				in.Delim('[')
				if out.ComputedStyles == nil {
					if !in.IsDelim(']') {
						out.ComputedStyles = make([]*ComputedStyle, 0, 8)
					} else {
						out.ComputedStyles = []*ComputedStyle{}
					}
				} else {
					out.ComputedStyles = (out.ComputedStyles)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *ComputedStyle
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(ComputedStyle)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.ComputedStyles = append(out.ComputedStyles, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot2(out *jwriter.Writer, in GetSnapshotReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.DomNodes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"domNodes\":")
		if in.DomNodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.DomNodes {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil {
					out.RawString("null")
				} else {
					(*v8).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.LayoutTreeNodes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"layoutTreeNodes\":")
		if in.LayoutTreeNodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.LayoutTreeNodes {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.ComputedStyles) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"computedStyles\":")
		if in.ComputedStyles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.ComputedStyles {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetSnapshotReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetSnapshotReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetSnapshotReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetSnapshotReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot3(in *jlexer.Lexer, out *GetSnapshotParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "computedStyleWhitelist":
			if in.IsNull() {
				in.Skip()
				out.ComputedStyleWhitelist = nil
			} else {
				in.Delim('[')
				if out.ComputedStyleWhitelist == nil {
					if !in.IsDelim(']') {
						out.ComputedStyleWhitelist = make([]string, 0, 4)
					} else {
						out.ComputedStyleWhitelist = []string{}
					}
				} else {
					out.ComputedStyleWhitelist = (out.ComputedStyleWhitelist)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.ComputedStyleWhitelist = append(out.ComputedStyleWhitelist, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot3(out *jwriter.Writer, in GetSnapshotParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"computedStyleWhitelist\":")
	if in.ComputedStyleWhitelist == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in.ComputedStyleWhitelist {
			if v14 > 0 {
				out.RawByte(',')
			}
			out.String(string(v15))
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetSnapshotParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetSnapshotParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetSnapshotParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetSnapshotParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot4(in *jlexer.Lexer, out *DOMNode) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nodeType":
			(out.NodeType).UnmarshalEasyJSON(in)
		case "nodeName":
			out.NodeName = string(in.String())
		case "nodeValue":
			out.NodeValue = string(in.String())
		case "backendNodeId":
			(out.BackendNodeID).UnmarshalEasyJSON(in)
		case "childNodeIndexes":
			if in.IsNull() {
				in.Skip()
				out.ChildNodeIndexes = nil
			} else {
				in.Delim('[')
				if out.ChildNodeIndexes == nil {
					if !in.IsDelim(']') {
						out.ChildNodeIndexes = make([]int64, 0, 8)
					} else {
						out.ChildNodeIndexes = []int64{}
					}
				} else {
					out.ChildNodeIndexes = (out.ChildNodeIndexes)[:0]
				}
				for !in.IsDelim(']') {
					var v16 int64
					v16 = int64(in.Int64())
					out.ChildNodeIndexes = append(out.ChildNodeIndexes, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "attributes":
			if in.IsNull() {
				in.Skip()
				out.Attributes = nil
			} else {
				in.Delim('[')
				if out.Attributes == nil {
					if !in.IsDelim(']') {
						out.Attributes = make([]*NameValue, 0, 8)
					} else {
						out.Attributes = []*NameValue{}
					}
				} else {
					out.Attributes = (out.Attributes)[:0]
				}
				for !in.IsDelim(']') {
					var v17 *NameValue
					if in.IsNull() {
						in.Skip()
						v17 = nil
					} else {
						if v17 == nil {
							v17 = new(NameValue)
						}
						(*v17).UnmarshalEasyJSON(in)
					}
					out.Attributes = append(out.Attributes, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pseudoElementIndexes":
			if in.IsNull() {
				in.Skip()
				out.PseudoElementIndexes = nil
			} else {
				in.Delim('[')
				if out.PseudoElementIndexes == nil {
					if !in.IsDelim(']') {
						out.PseudoElementIndexes = make([]int64, 0, 8)
					} else {
						out.PseudoElementIndexes = []int64{}
					}
				} else {
					out.PseudoElementIndexes = (out.PseudoElementIndexes)[:0]
				}
				for !in.IsDelim(']') {
					var v18 int64
					v18 = int64(in.Int64())
					out.PseudoElementIndexes = append(out.PseudoElementIndexes, v18)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "layoutNodeIndex":
			out.LayoutNodeIndex = int64(in.Int64())
		case "documentURL":
			out.DocumentURL = string(in.String())
		case "baseURL":
			out.BaseURL = string(in.String())
		case "contentLanguage":
			out.ContentLanguage = string(in.String())
		case "publicId":
			out.PublicID = string(in.String())
		case "systemId":
			out.SystemID = string(in.String())
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "contentDocumentIndex":
			out.ContentDocumentIndex = int64(in.Int64())
		case "importedDocumentIndex":
			out.ImportedDocumentIndex = int64(in.Int64())
		case "templateContentIndex":
			out.TemplateContentIndex = int64(in.Int64())
		case "pseudoType":
			(out.PseudoType).UnmarshalEasyJSON(in)
		case "isClickable":
			out.IsClickable = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot4(out *jwriter.Writer, in DOMNode) {
	out.RawByte('{')
	first := true
	_ = first
	if in.NodeType != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"nodeType\":")
		(in.NodeType).MarshalEasyJSON(out)
	}
	if in.NodeName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"nodeName\":")
		out.String(string(in.NodeName))
	}
	if in.NodeValue != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"nodeValue\":")
		out.String(string(in.NodeValue))
	}
	if in.BackendNodeID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"backendNodeId\":")
		out.Int64(int64(in.BackendNodeID))
	}
	if len(in.ChildNodeIndexes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"childNodeIndexes\":")
		if in.ChildNodeIndexes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v19, v20 := range in.ChildNodeIndexes {
				if v19 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v20))
			}
			out.RawByte(']')
		}
	}
	if len(in.Attributes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"attributes\":")
		if in.Attributes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v21, v22 := range in.Attributes {
				if v21 > 0 {
					out.RawByte(',')
				}
				if v22 == nil {
					out.RawString("null")
				} else {
					(*v22).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.PseudoElementIndexes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pseudoElementIndexes\":")
		if in.PseudoElementIndexes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v23, v24 := range in.PseudoElementIndexes {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v24))
			}
			out.RawByte(']')
		}
	}
	if in.LayoutNodeIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"layoutNodeIndex\":")
		out.Int64(int64(in.LayoutNodeIndex))
	}
	if in.DocumentURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"documentURL\":")
		out.String(string(in.DocumentURL))
	}
	if in.BaseURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"baseURL\":")
		out.String(string(in.BaseURL))
	}
	if in.ContentLanguage != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"contentLanguage\":")
		out.String(string(in.ContentLanguage))
	}
	if in.PublicID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"publicId\":")
		out.String(string(in.PublicID))
	}
	if in.SystemID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"systemId\":")
		out.String(string(in.SystemID))
	}
	if in.FrameID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"frameId\":")
		out.String(string(in.FrameID))
	}
	if in.ContentDocumentIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"contentDocumentIndex\":")
		out.Int64(int64(in.ContentDocumentIndex))
	}
	if in.ImportedDocumentIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"importedDocumentIndex\":")
		out.Int64(int64(in.ImportedDocumentIndex))
	}
	if in.TemplateContentIndex != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"templateContentIndex\":")
		out.Int64(int64(in.TemplateContentIndex))
	}
	if in.PseudoType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pseudoType\":")
		(in.PseudoType).MarshalEasyJSON(out)
	}
	if in.IsClickable {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isClickable\":")
		out.Bool(bool(in.IsClickable))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DOMNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DOMNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DOMNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DOMNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot5(in *jlexer.Lexer, out *ComputedStyle) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "properties":
			if in.IsNull() {
				in.Skip()
				out.Properties = nil
			} else {
				in.Delim('[')
				if out.Properties == nil {
					if !in.IsDelim(']') {
						out.Properties = make([]*NameValue, 0, 8)
					} else {
						out.Properties = []*NameValue{}
					}
				} else {
					out.Properties = (out.Properties)[:0]
				}
				for !in.IsDelim(']') {
					var v25 *NameValue
					if in.IsNull() {
						in.Skip()
						v25 = nil
					} else {
						if v25 == nil {
							v25 = new(NameValue)
						}
						(*v25).UnmarshalEasyJSON(in)
					}
					out.Properties = append(out.Properties, v25)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot5(out *jwriter.Writer, in ComputedStyle) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Properties) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"properties\":")
		if in.Properties == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Properties {
				if v26 > 0 {
					out.RawByte(',')
				}
				if v27 == nil {
					out.RawString("null")
				} else {
					(*v27).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ComputedStyle) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ComputedStyle) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpDomsnapshot5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ComputedStyle) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ComputedStyle) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpDomsnapshot5(l, v)
}
