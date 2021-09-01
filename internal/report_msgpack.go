package internal

import (
	"fmt"
	_ "github.com/citilinkru/uuid-msgpack"
	"github.com/google/uuid"
	"gopkg.in/vmihailenco/msgpack.v2"
	"time"
)

// DecodeMsgpack декодирует отчет из Msgpack
func (a *Report) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}
	if length != 6 {
		return fmt.Errorf("decode error; len doesn't match: %d", length)
	}

	id := uuid.UUID{}
	if err = d.Decode(&id); err != nil { // 1
		return fmt.Errorf("decode error: can't decode field uuid: %w", err)
	}
	a.Id = ReportId(id)

	if v, err := d.DecodeString(); err != nil { // 2
		return fmt.Errorf("decode error: can't decode field NavActionNumber: %w", err)
	} else {
		a.NavActionNumber = v
	}

	if v, err := d.DecodeInt64(); err != nil { // 3
		return fmt.Errorf("can't decode field CampaignStartDate: %w", err)
	} else {
		a.CampaignStartDate = time.Unix(v, 0)
	}

	if v, err := d.DecodeInt64(); err != nil { // 4
		return fmt.Errorf("can't decode field LastModified: %w", err)
	} else {
		a.LastModified = time.Unix(v, 0)
	}

	if v, err := d.DecodeInt(); err != nil { // 5
		return fmt.Errorf("can't decode field status: %w", err)
	} else {
		a.Status = ReportStatus(v)
	}

	if v, err := d.DecodeInt(); err != nil { // 6
		return fmt.Errorf("can't decode field type: %w", err)
	} else {
		a.Type = ReportType(v)
	}

	return nil
}
