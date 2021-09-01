package internal

import (
	"bytes"
	_ "github.com/citilinkru/uuid-msgpack"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
	"testing"
	"time"
)

func encode(r *Report) []byte {
	buf := bytes.NewBuffer(nil)
	e := msgpack.NewEncoder(buf)

	_ = e.EncodeArrayLen(6)
	_ = e.Encode(uuid.UUID(r.Id))
	_ = e.EncodeString(r.NavActionNumber)
	_ = e.EncodeInt64(r.CampaignStartDate.Unix())
	_ = e.EncodeInt64(r.LastModified.Unix())
	_ = e.EncodeInt(int(r.Status))
	_ = e.EncodeInt(int(r.Type))

	return buf.Bytes()
}

func TestFile_DecodeMsgpack(t *testing.T) {
	id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
	tests := []struct {
		name        string
		expectedObj *Report
		bytes       []byte
		wantErr     bool
		wantErrMsg  string
	}{
		{
			name:        "empty",
			expectedObj: &Report{CampaignStartDate: time.Unix(0, 0), LastModified: time.Unix(0, 0)},
		},
		{
			name: "with id",
			expectedObj: &Report{
				Id:                ReportId(id),
				CampaignStartDate: time.Unix(0, 0),
				LastModified:      time.Unix(0, 0),
			},
		},
		{
			name: "with nav action number",
			expectedObj: &Report{
				Id:                ReportId(id),
				NavActionNumber:   "test",
				CampaignStartDate: time.Unix(0, 0),
				LastModified:      time.Unix(0, 0),
			},
		},
		{
			name: "with type",
			expectedObj: &Report{
				Id:                ReportId(id),
				NavActionNumber:   "test",
				CampaignStartDate: time.Unix(0, 0),
				LastModified:      time.Unix(0, 0),
				Type:              ReportTypeDetailed,
			},
		},
		{
			name: "with status",
			expectedObj: &Report{
				Id:                ReportId(id),
				NavActionNumber:   "test",
				CampaignStartDate: time.Unix(0, 0),
				LastModified:      time.Unix(0, 0),
				Type:              ReportTypeDetailed,
				Status:            ReportStatusSuccess,
			},
		},
		{
			name:        "error decoding array length",
			expectedObj: &Report{},
			bytes:       []byte{1},
			wantErr:     true,
			wantErrMsg:  "decode error: msgpack: invalid code 1 decoding array length",
		},
		{
			name:        "error invalid array length",
			expectedObj: &Report{},
			bytes:       []byte{0x95, 0x0, 0x13},
			wantErr:     true,
			wantErrMsg:  "decode error; len doesn't match: 5",
		},
		{
			name:        "error invalid uuid bytes count",
			expectedObj: &Report{},
			bytes:       []byte{0x96, 0xa0, 0xa0, 0xa0, 0xa0, 0xa0},
			wantErr:     true,
			wantErrMsg:  "decode error: can't decode field uuid: invalid bytes count 5 instead of 18",
		},
		{
			name:        "error invalid field navActionNumber bytes",
			expectedObj: &Report{},
			bytes:       []byte{0x96, 216, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x0, 0xa0, 0xa0, 0xa0},
			wantErr:     true,
			wantErrMsg:  "decode error: can't decode field NavActionNumber: msgpack: invalid code 0 decoding bytes length",
		},
		{
			name:        "error invalid field CampaignStartDate bytes",
			expectedObj: &Report{},
			bytes:       []byte{0x96, 216, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xa0, 0xa0, 0xa0, 0xa0},
			wantErr:     true,
			wantErrMsg:  "can't decode field CampaignStartDate: msgpack: invalid code a0 decoding int64",
		},
		{
			name:        "error invalid field LastModified bytes",
			expectedObj: &Report{},
			bytes:       []byte{0x96, 216, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xa0, 0x0, 0xa0, 0xa0},
			wantErr:     true,
			wantErrMsg:  "can't decode field LastModified: msgpack: invalid code a0 decoding int64",
		},
		{
			name:        "error invalid field status bytes",
			expectedObj: &Report{},
			bytes:       []byte{0x96, 216, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xa0, 0x0, 0x0, 0xa0},
			wantErr:     true,
			wantErrMsg:  "can't decode field status: msgpack: invalid code a0 decoding int64",
		},
		{
			name:        "error invalid field type bytes",
			expectedObj: &Report{},
			bytes:       []byte{0x96, 216, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xa0, 0x0, 0x0, 0x0, 0xa0},
			wantErr:     true,
			wantErrMsg:  "can't decode field type: msgpack: invalid code a0 decoding int64",
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				var bufBytes []byte
				if test.wantErr {
					bufBytes = test.bytes
				} else {
					bufBytes = encode(test.expectedObj)
				}

				buf := bytes.NewBuffer(bufBytes)
				dec := msgpack.NewDecoder(buf)
				obj := &Report{}
				err := obj.DecodeMsgpack(dec)

				if test.wantErr {
					assert.EqualErrorf(t, err, test.wantErrMsg, err.Error())
				} else {
					assert.NoError(t, err)
					assert.Equal(t, test.expectedObj, obj)
				}
			},
		)
	}
}
