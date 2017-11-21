// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wanikani.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	wanikani.proto

It has these top-level messages:
	Meaning
	Reading
	Radical
	Kanji
	Vocabulary
	Subject
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Reading_Type int32

const (
	Reading_ONYOMI  Reading_Type = 1
	Reading_KUNYOMI Reading_Type = 2
	Reading_NANORI  Reading_Type = 3
)

var Reading_Type_name = map[int32]string{
	1: "ONYOMI",
	2: "KUNYOMI",
	3: "NANORI",
}
var Reading_Type_value = map[string]int32{
	"ONYOMI":  1,
	"KUNYOMI": 2,
	"NANORI":  3,
}

func (x Reading_Type) Enum() *Reading_Type {
	p := new(Reading_Type)
	*p = x
	return p
}
func (x Reading_Type) String() string {
	return proto1.EnumName(Reading_Type_name, int32(x))
}
func (x *Reading_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Reading_Type_value, data, "Reading_Type")
	if err != nil {
		return err
	}
	*x = Reading_Type(value)
	return nil
}
func (Reading_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Vocabulary_PartOfSpeech int32

const (
	Vocabulary_NOUN              Vocabulary_PartOfSpeech = 1
	Vocabulary_NUMERAL           Vocabulary_PartOfSpeech = 2
	Vocabulary_INTRANSITIVE_VERB Vocabulary_PartOfSpeech = 3
	Vocabulary_ICHIDAN_VERB      Vocabulary_PartOfSpeech = 4
	Vocabulary_TRANSITIVE_VERB   Vocabulary_PartOfSpeech = 5
	Vocabulary_NO_ADJECTIVE      Vocabulary_PartOfSpeech = 6
	Vocabulary_GODAN_VERB        Vocabulary_PartOfSpeech = 7
	Vocabulary_NA_ADJECTIVE      Vocabulary_PartOfSpeech = 8
	Vocabulary_I_ADJECTIVE       Vocabulary_PartOfSpeech = 9
	Vocabulary_SUFFIX            Vocabulary_PartOfSpeech = 10
	Vocabulary_ADVERB            Vocabulary_PartOfSpeech = 11
	Vocabulary_SURU_VERB         Vocabulary_PartOfSpeech = 12
	Vocabulary_PREFIX            Vocabulary_PartOfSpeech = 13
	Vocabulary_PROPER_NOUN       Vocabulary_PartOfSpeech = 14
	Vocabulary_EXPRESSION        Vocabulary_PartOfSpeech = 15
	Vocabulary_ADJECTIVE         Vocabulary_PartOfSpeech = 16
	Vocabulary_INTERJECTION      Vocabulary_PartOfSpeech = 17
	Vocabulary_COUNTER           Vocabulary_PartOfSpeech = 18
	Vocabulary_PRONOUN           Vocabulary_PartOfSpeech = 19
	Vocabulary_CONJUNCTION       Vocabulary_PartOfSpeech = 20
)

var Vocabulary_PartOfSpeech_name = map[int32]string{
	1:  "NOUN",
	2:  "NUMERAL",
	3:  "INTRANSITIVE_VERB",
	4:  "ICHIDAN_VERB",
	5:  "TRANSITIVE_VERB",
	6:  "NO_ADJECTIVE",
	7:  "GODAN_VERB",
	8:  "NA_ADJECTIVE",
	9:  "I_ADJECTIVE",
	10: "SUFFIX",
	11: "ADVERB",
	12: "SURU_VERB",
	13: "PREFIX",
	14: "PROPER_NOUN",
	15: "EXPRESSION",
	16: "ADJECTIVE",
	17: "INTERJECTION",
	18: "COUNTER",
	19: "PRONOUN",
	20: "CONJUNCTION",
}
var Vocabulary_PartOfSpeech_value = map[string]int32{
	"NOUN":              1,
	"NUMERAL":           2,
	"INTRANSITIVE_VERB": 3,
	"ICHIDAN_VERB":      4,
	"TRANSITIVE_VERB":   5,
	"NO_ADJECTIVE":      6,
	"GODAN_VERB":        7,
	"NA_ADJECTIVE":      8,
	"I_ADJECTIVE":       9,
	"SUFFIX":            10,
	"ADVERB":            11,
	"SURU_VERB":         12,
	"PREFIX":            13,
	"PROPER_NOUN":       14,
	"EXPRESSION":        15,
	"ADJECTIVE":         16,
	"INTERJECTION":      17,
	"COUNTER":           18,
	"PRONOUN":           19,
	"CONJUNCTION":       20,
}

func (x Vocabulary_PartOfSpeech) Enum() *Vocabulary_PartOfSpeech {
	p := new(Vocabulary_PartOfSpeech)
	*p = x
	return p
}
func (x Vocabulary_PartOfSpeech) String() string {
	return proto1.EnumName(Vocabulary_PartOfSpeech_name, int32(x))
}
func (x *Vocabulary_PartOfSpeech) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Vocabulary_PartOfSpeech_value, data, "Vocabulary_PartOfSpeech")
	if err != nil {
		return err
	}
	*x = Vocabulary_PartOfSpeech(value)
	return nil
}
func (Vocabulary_PartOfSpeech) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Meaning struct {
	Meaning          *string `protobuf:"bytes,1,opt,name=meaning" json:"meaning,omitempty"`
	IsPrimary        *bool   `protobuf:"varint,2,opt,name=is_primary" json:"is_primary,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Meaning) Reset()                    { *m = Meaning{} }
func (m *Meaning) String() string            { return proto1.CompactTextString(m) }
func (*Meaning) ProtoMessage()               {}
func (*Meaning) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Meaning) GetMeaning() string {
	if m != nil && m.Meaning != nil {
		return *m.Meaning
	}
	return ""
}

func (m *Meaning) GetIsPrimary() bool {
	if m != nil && m.IsPrimary != nil {
		return *m.IsPrimary
	}
	return false
}

type Reading struct {
	Reading          *string       `protobuf:"bytes,1,opt,name=reading" json:"reading,omitempty"`
	IsPrimary        *bool         `protobuf:"varint,2,opt,name=is_primary" json:"is_primary,omitempty"`
	Type             *Reading_Type `protobuf:"varint,3,opt,name=type,enum=proto.Reading_Type" json:"type,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Reading) Reset()                    { *m = Reading{} }
func (m *Reading) String() string            { return proto1.CompactTextString(m) }
func (*Reading) ProtoMessage()               {}
func (*Reading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reading) GetReading() string {
	if m != nil && m.Reading != nil {
		return *m.Reading
	}
	return ""
}

func (m *Reading) GetIsPrimary() bool {
	if m != nil && m.IsPrimary != nil {
		return *m.IsPrimary
	}
	return false
}

func (m *Reading) GetType() Reading_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Reading_ONYOMI
}

type Radical struct {
	Japanese         *string    `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	Meanings         []*Meaning `protobuf:"bytes,2,rep,name=meanings" json:"meanings,omitempty"`
	CharacterImage   *string    `protobuf:"bytes,3,opt,name=character_image" json:"character_image,omitempty"`
	Mnemonic         *string    `protobuf:"bytes,4,opt,name=mnemonic" json:"mnemonic,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Radical) Reset()                    { *m = Radical{} }
func (m *Radical) String() string            { return proto1.CompactTextString(m) }
func (*Radical) ProtoMessage()               {}
func (*Radical) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Radical) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Radical) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Radical) GetCharacterImage() string {
	if m != nil && m.CharacterImage != nil {
		return *m.CharacterImage
	}
	return ""
}

func (m *Radical) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

type Kanji struct {
	Japanese            *string    `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	Readings            []*Reading `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty"`
	Meanings            []*Meaning `protobuf:"bytes,3,rep,name=meanings" json:"meanings,omitempty"`
	MeaningMnemonic     *string    `protobuf:"bytes,4,opt,name=meaning_mnemonic" json:"meaning_mnemonic,omitempty"`
	MeaningHint         *string    `protobuf:"bytes,5,opt,name=meaning_hint" json:"meaning_hint,omitempty"`
	ReadingMnemonic     *string    `protobuf:"bytes,6,opt,name=reading_mnemonic" json:"reading_mnemonic,omitempty"`
	ReadingHint         *string    `protobuf:"bytes,7,opt,name=reading_hint" json:"reading_hint,omitempty"`
	ComponentSubjectIds []int32    `protobuf:"varint,8,rep,name=component_subject_ids" json:"component_subject_ids,omitempty"`
	XXX_unrecognized    []byte     `json:"-"`
}

func (m *Kanji) Reset()                    { *m = Kanji{} }
func (m *Kanji) String() string            { return proto1.CompactTextString(m) }
func (*Kanji) ProtoMessage()               {}
func (*Kanji) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Kanji) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Kanji) GetReadings() []*Reading {
	if m != nil {
		return m.Readings
	}
	return nil
}

func (m *Kanji) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Kanji) GetMeaningMnemonic() string {
	if m != nil && m.MeaningMnemonic != nil {
		return *m.MeaningMnemonic
	}
	return ""
}

func (m *Kanji) GetMeaningHint() string {
	if m != nil && m.MeaningHint != nil {
		return *m.MeaningHint
	}
	return ""
}

func (m *Kanji) GetReadingMnemonic() string {
	if m != nil && m.ReadingMnemonic != nil {
		return *m.ReadingMnemonic
	}
	return ""
}

func (m *Kanji) GetReadingHint() string {
	if m != nil && m.ReadingHint != nil {
		return *m.ReadingHint
	}
	return ""
}

func (m *Kanji) GetComponentSubjectIds() []int32 {
	if m != nil {
		return m.ComponentSubjectIds
	}
	return nil
}

type Vocabulary struct {
	Japanese            *string                   `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	Readings            []*Reading                `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty"`
	Meanings            []*Meaning                `protobuf:"bytes,3,rep,name=meanings" json:"meanings,omitempty"`
	MeaningExplanation  *string                   `protobuf:"bytes,4,opt,name=meaning_explanation" json:"meaning_explanation,omitempty"`
	ReadingExplanation  *string                   `protobuf:"bytes,5,opt,name=reading_explanation" json:"reading_explanation,omitempty"`
	Sentences           []*Vocabulary_Sentence    `protobuf:"bytes,6,rep,name=sentences" json:"sentences,omitempty"`
	ComponentSubjectIds []int32                   `protobuf:"varint,7,rep,name=component_subject_ids" json:"component_subject_ids,omitempty"`
	PartsOfSpeech       []Vocabulary_PartOfSpeech `protobuf:"varint,8,rep,name=parts_of_speech,enum=proto.Vocabulary_PartOfSpeech" json:"parts_of_speech,omitempty"`
	Audio               *string                   `protobuf:"bytes,9,opt,name=audio" json:"audio,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *Vocabulary) Reset()                    { *m = Vocabulary{} }
func (m *Vocabulary) String() string            { return proto1.CompactTextString(m) }
func (*Vocabulary) ProtoMessage()               {}
func (*Vocabulary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Vocabulary) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Vocabulary) GetReadings() []*Reading {
	if m != nil {
		return m.Readings
	}
	return nil
}

func (m *Vocabulary) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Vocabulary) GetMeaningExplanation() string {
	if m != nil && m.MeaningExplanation != nil {
		return *m.MeaningExplanation
	}
	return ""
}

func (m *Vocabulary) GetReadingExplanation() string {
	if m != nil && m.ReadingExplanation != nil {
		return *m.ReadingExplanation
	}
	return ""
}

func (m *Vocabulary) GetSentences() []*Vocabulary_Sentence {
	if m != nil {
		return m.Sentences
	}
	return nil
}

func (m *Vocabulary) GetComponentSubjectIds() []int32 {
	if m != nil {
		return m.ComponentSubjectIds
	}
	return nil
}

func (m *Vocabulary) GetPartsOfSpeech() []Vocabulary_PartOfSpeech {
	if m != nil {
		return m.PartsOfSpeech
	}
	return nil
}

func (m *Vocabulary) GetAudio() string {
	if m != nil && m.Audio != nil {
		return *m.Audio
	}
	return ""
}

type Vocabulary_Sentence struct {
	Japanese         *string `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	English          *string `protobuf:"bytes,2,opt,name=english" json:"english,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Vocabulary_Sentence) Reset()                    { *m = Vocabulary_Sentence{} }
func (m *Vocabulary_Sentence) String() string            { return proto1.CompactTextString(m) }
func (*Vocabulary_Sentence) ProtoMessage()               {}
func (*Vocabulary_Sentence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Vocabulary_Sentence) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Vocabulary_Sentence) GetEnglish() string {
	if m != nil && m.English != nil {
		return *m.English
	}
	return ""
}

type Subject struct {
	Id               *int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Level            *int32      `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	Slug             *string     `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	DocumentUrl      *string     `protobuf:"bytes,4,opt,name=document_url" json:"document_url,omitempty"`
	Radical          *Radical    `protobuf:"bytes,5,opt,name=radical" json:"radical,omitempty"`
	Kanji            *Kanji      `protobuf:"bytes,6,opt,name=kanji" json:"kanji,omitempty"`
	Vocabulary       *Vocabulary `protobuf:"bytes,7,opt,name=vocabulary" json:"vocabulary,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto1.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Subject) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Subject) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Subject) GetSlug() string {
	if m != nil && m.Slug != nil {
		return *m.Slug
	}
	return ""
}

func (m *Subject) GetDocumentUrl() string {
	if m != nil && m.DocumentUrl != nil {
		return *m.DocumentUrl
	}
	return ""
}

func (m *Subject) GetRadical() *Radical {
	if m != nil {
		return m.Radical
	}
	return nil
}

func (m *Subject) GetKanji() *Kanji {
	if m != nil {
		return m.Kanji
	}
	return nil
}

func (m *Subject) GetVocabulary() *Vocabulary {
	if m != nil {
		return m.Vocabulary
	}
	return nil
}

func init() {
	proto1.RegisterType((*Meaning)(nil), "proto.Meaning")
	proto1.RegisterType((*Reading)(nil), "proto.Reading")
	proto1.RegisterType((*Radical)(nil), "proto.Radical")
	proto1.RegisterType((*Kanji)(nil), "proto.Kanji")
	proto1.RegisterType((*Vocabulary)(nil), "proto.Vocabulary")
	proto1.RegisterType((*Vocabulary_Sentence)(nil), "proto.Vocabulary.Sentence")
	proto1.RegisterType((*Subject)(nil), "proto.Subject")
	proto1.RegisterEnum("proto.Reading_Type", Reading_Type_name, Reading_Type_value)
	proto1.RegisterEnum("proto.Vocabulary_PartOfSpeech", Vocabulary_PartOfSpeech_name, Vocabulary_PartOfSpeech_value)
}

func init() { proto1.RegisterFile("wanikani.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0x16, 0x3f, 0xc6, 0x70, 0x20, 0xe0, 0x0c, 0x89, 0xae, 0x95, 0xe8, 0xde, 0x4b, 0x2d, 0x55,
	0x42, 0xaa, 0xc2, 0x82, 0x4d, 0xd7, 0x94, 0x38, 0xad, 0x93, 0xc6, 0x46, 0x36, 0x8e, 0xd2, 0x95,
	0x35, 0x31, 0x13, 0x98, 0xc4, 0x8c, 0x2d, 0xdb, 0xa4, 0x65, 0xd3, 0x27, 0x6a, 0xdf, 0xa3, 0x0f,
	0xd1, 0x87, 0xa9, 0x66, 0x3c, 0x24, 0xe4, 0x6f, 0xd9, 0x95, 0xed, 0x73, 0xbe, 0xf9, 0xbe, 0xf3,
	0x1d, 0x9f, 0x33, 0xd0, 0xfe, 0x8a, 0x19, 0xbd, 0xc5, 0x8c, 0x0e, 0x92, 0x34, 0xce, 0x63, 0xa4,
	0x88, 0x87, 0x31, 0x00, 0xf5, 0x9c, 0x60, 0x46, 0xd9, 0x1c, 0x75, 0x40, 0x5d, 0x16, 0xaf, 0x7a,
	0xa9, 0x57, 0xea, 0x37, 0x10, 0x02, 0xa0, 0x59, 0x90, 0xa4, 0x74, 0x89, 0xd3, 0xb5, 0x5e, 0xee,
	0x95, 0xfa, 0x75, 0xe3, 0x3b, 0xa8, 0x2e, 0xc1, 0x33, 0x89, 0x4f, 0x8b, 0xd7, 0xd7, 0xf1, 0xe8,
	0x0d, 0x54, 0xf3, 0x75, 0x42, 0xf4, 0x4a, 0xaf, 0xd4, 0x6f, 0x0f, 0xbb, 0x85, 0xf8, 0x40, 0x52,
	0x0c, 0xa6, 0xeb, 0x84, 0x18, 0xef, 0xa0, 0xca, 0x9f, 0x08, 0xa0, 0xe6, 0xd8, 0x5f, 0x9c, 0x73,
	0x4b, 0x2b, 0xa1, 0x26, 0xa8, 0x67, 0x7e, 0xf1, 0x51, 0xe6, 0x09, 0x7b, 0x64, 0x3b, 0xae, 0xa5,
	0x55, 0x8c, 0x05, 0xa8, 0x2e, 0x9e, 0xd1, 0x10, 0x47, 0x48, 0x83, 0xfa, 0x0d, 0x4e, 0x30, 0x23,
	0x19, 0x91, 0x05, 0xf4, 0xa0, 0x2e, 0x1d, 0x64, 0x7a, 0xb9, 0x57, 0xe9, 0x37, 0x87, 0x6d, 0x29,
	0xb8, 0xf1, 0xf8, 0x0f, 0x74, 0xc2, 0x05, 0x4e, 0x71, 0x98, 0x93, 0x34, 0xa0, 0x4b, 0x3c, 0x2f,
	0x2a, 0x6b, 0x70, 0xb2, 0x25, 0x23, 0xcb, 0x98, 0xd1, 0x50, 0xaf, 0xf2, 0x88, 0xf1, 0xbb, 0x04,
	0xca, 0x19, 0x66, 0x37, 0xf4, 0x65, 0x21, 0x69, 0xfd, 0xa9, 0xd0, 0xa6, 0x39, 0xdb, 0xa5, 0x54,
	0x5e, 0x2c, 0x45, 0x07, 0x4d, 0x22, 0x82, 0xc7, 0xca, 0x68, 0x0f, 0x5a, 0x9b, 0xcc, 0x82, 0xb2,
	0x5c, 0x57, 0x44, 0x54, 0x07, 0x4d, 0x6a, 0x3e, 0xe0, 0x6b, 0x1b, 0xfc, 0x26, 0x23, 0xf0, 0xaa,
	0x88, 0xfe, 0x0b, 0xfb, 0x61, 0xbc, 0x4c, 0x62, 0x46, 0x58, 0x1e, 0x64, 0xab, 0xab, 0x1b, 0x12,
	0xe6, 0x01, 0x9d, 0x65, 0x7a, 0xbd, 0x57, 0xe9, 0x2b, 0xc6, 0x0f, 0x05, 0xe0, 0x22, 0x0e, 0xf1,
	0xd5, 0x2a, 0xc2, 0xe9, 0xfa, 0x2f, 0x79, 0x3c, 0x84, 0xee, 0xc6, 0x09, 0xf9, 0x96, 0x44, 0x98,
	0xe1, 0x9c, 0xc6, 0x4c, 0xda, 0x3c, 0x84, 0xee, 0xa6, 0xec, 0xed, 0x64, 0xe1, 0xf6, 0x08, 0x1a,
	0x19, 0x61, 0x39, 0x61, 0x21, 0xc9, 0xf4, 0x9a, 0x20, 0x3f, 0x90, 0xe4, 0x0f, 0x55, 0x0f, 0x3c,
	0x09, 0x79, 0xdd, 0xac, 0xca, 0xcd, 0xa2, 0xf7, 0xd0, 0x49, 0x70, 0x9a, 0x67, 0x41, 0x7c, 0x1d,
	0x64, 0x09, 0x21, 0xe1, 0x42, 0x74, 0xa1, 0x3d, 0xfc, 0xef, 0x39, 0xe7, 0x04, 0xa7, 0xb9, 0x73,
	0xed, 0x09, 0x14, 0xda, 0x01, 0x05, 0xaf, 0x66, 0x34, 0xd6, 0x1b, 0xbc, 0xaa, 0x83, 0x23, 0xa8,
	0xdf, 0x4b, 0x3e, 0xef, 0x58, 0x07, 0x54, 0xc2, 0xe6, 0x11, 0xcd, 0x16, 0x62, 0xf8, 0x1b, 0xc6,
	0xaf, 0x32, 0xb4, 0x1e, 0xd1, 0xd5, 0xa1, 0x6a, 0x3b, 0xbe, 0x5d, 0x0c, 0xb8, 0xed, 0x9f, 0x9b,
	0xee, 0xe8, 0xb3, 0x56, 0x46, 0xfb, 0xb0, 0x6b, 0xd9, 0x53, 0x77, 0x64, 0x7b, 0xd6, 0xd4, 0xba,
	0x30, 0x83, 0x0b, 0xd3, 0xfd, 0xa0, 0x55, 0x90, 0x06, 0x2d, 0x6b, 0xfc, 0xc9, 0x3a, 0x1e, 0xd9,
	0x45, 0xa4, 0x8a, 0xba, 0xd0, 0x79, 0x0a, 0x53, 0x38, 0xcc, 0x76, 0x82, 0xd1, 0xf1, 0xa9, 0x39,
	0xe6, 0x61, 0xad, 0x86, 0xda, 0x00, 0x1f, 0x9d, 0xfb, 0x63, 0xaa, 0x40, 0x8c, 0xb6, 0x10, 0x75,
	0xd4, 0x81, 0xa6, 0xb5, 0x15, 0x68, 0xf0, 0x1d, 0xf3, 0xfc, 0x93, 0x13, 0xeb, 0x52, 0x03, 0xfe,
	0x3e, 0x3a, 0x16, 0x47, 0x9b, 0x68, 0x07, 0x1a, 0x9e, 0xef, 0xfa, 0x05, 0x53, 0x8b, 0xa7, 0x26,
	0xae, 0xc9, 0x61, 0x3b, 0x9c, 0x63, 0xe2, 0x3a, 0x13, 0xd3, 0x0d, 0x84, 0xa7, 0x36, 0x97, 0x35,
	0x2f, 0x27, 0xae, 0xe9, 0x79, 0x96, 0x63, 0x6b, 0x1d, 0x7e, 0xf6, 0x41, 0x42, 0x13, 0x76, 0xec,
	0xa9, 0xe9, 0x8a, 0x88, 0x63, 0x6b, 0xbb, 0xbc, 0x09, 0x63, 0xc7, 0xe7, 0x31, 0x0d, 0xf1, 0x8f,
	0x89, 0xeb, 0x08, 0xaa, 0x2e, 0xe7, 0x1e, 0x3b, 0xf6, 0xa9, 0x6f, 0x17, 0xd0, 0x3d, 0xe3, 0x67,
	0x09, 0x54, 0xaf, 0xf8, 0xaf, 0x08, 0xa0, 0x4c, 0x67, 0xa2, 0xe7, 0x0a, 0xff, 0x41, 0x11, 0xb9,
	0x23, 0x91, 0xe8, 0xb8, 0x82, 0x5a, 0x50, 0xcd, 0xa2, 0xd5, 0x5c, 0x2e, 0xf5, 0x1e, 0xb4, 0x66,
	0x71, 0xb8, 0x5a, 0xf2, 0xa1, 0x58, 0xa5, 0x91, 0x9c, 0xbb, 0xff, 0x41, 0x4d, 0x8b, 0x2b, 0x44,
	0xcc, 0xda, 0xd6, 0x5c, 0xcb, 0x8b, 0xe5, 0x10, 0x94, 0x5b, 0xbe, 0xf8, 0x62, 0xbd, 0x9a, 0xc3,
	0x96, 0x4c, 0x17, 0x97, 0xc1, 0x5b, 0x80, 0xbb, 0xfb, 0x61, 0x11, 0xab, 0xd6, 0x1c, 0xee, 0x3e,
	0x9b, 0xa2, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa8, 0xa5, 0x9a, 0x6f, 0x05, 0x00, 0x00,
}
