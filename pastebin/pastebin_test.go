package pastebin

import (
	"testing"
)

func TestUnmarshalPaste(t *testing.T) {
	rawList := `<paste>
<paste_key>dvTfMxH7</paste_key>
<paste_date>1544546506</paste_date>
<paste_title>kuba</paste_title>
<paste_size>5</paste_size>
<paste_expire_date>1547138506</paste_expire_date>
<paste_private>0</paste_private>
<paste_format_long>Go</paste_format_long>
<paste_format_short>go</paste_format_short>
<paste_url>https://pastebin.com/dvTfMxH7</paste_url>
<paste_hits>1</paste_hits>
</paste>
<paste>
<paste_key>fnd0tjJ9</paste_key>
<paste_date>1544542503</paste_date>
<paste_title></paste_title>
<paste_size>5</paste_size>
<paste_expire_date>1547134503</paste_expire_date>
<paste_private>0</paste_private>
<paste_format_long>None</paste_format_long>
<paste_format_short>text</paste_format_short>
<paste_url>https://pastebin.com/fnd0tjJ9</paste_url>
<paste_hits>8</paste_hits>
</paste>`

	list, err := parseListXML(rawList)
	if err != nil {
		t.Errorf("parsing list failed: %s", rawList)
	}

	if len(list) != 2 {
		t.Errorf("expected to have 2 pastes after unmarshalling, got %d", len(list))
	}
}
