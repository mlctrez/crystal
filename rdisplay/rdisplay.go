package rdisplay

import "net"

const (
	Null = iota
	// Moves cursor to the top left character position. No data is changed. Identical to Control Q,0,0.
	CursorHome
	// Display is blanked, no data is changed.
	HideDisplay
	// Restores blanked display; nothing else is changed.
	RestoreDisplay
	// Cursor is not shown; nothing else is changed.
	HideCursor
	// Shows a non-blinking underline cursor at the printing location.
	ShowUnderlineCursor
	// Shows a blinking block cursor at the printing location.
	ShowBlockCursor
	// Shows a blinking block cursor at the printing location. This cursor inverts the character rather
	// than replacing the character with a block. This cursor style is the default cursor at power-up.
	ShowInvertingBlockCursor
	// Moves the cursor back one space and erases the character in that space. Will wrap from the left-most
	// column to the right-most column of the line above. Will wrap from the left-most column of the first
	// row to the right-most column of the last row.
	Backspace
	_
	// Moves the cursor down one row. If SCROLL is on and the cursor is at the bottom row, the display
	// will scroll up one row and the bottom row will be cleared. If SCROLL is off, and the cursor is
	// at the bottom row, it will wrap up to the same character position on the top row.
	LineFeed
	// Deletes the character at the current cursor position. Cursor is not moved.
	DeleteInPlace
	// Clears the display and returns cursor to Home position (upper left). All data is erased.
	FormFeed
	// Moves cursor to the left-most column of the current row.
	CarriageReturn
	// Send "Control-N" , followed by a byte from 0-100 for the backlight brightness.
	// 0=OFF; 100=ON , intermediate values will vary the brightness.
	// There are a total of 25 possible brightness levels.
	BacklightControl
	// Send "Control O", followed by a byte from 0-100 for the contrast setting of the displayed characters.
	// 0 = very light, 100 = very dark, 50 is typical.
	// There are a total of 25 possible contrast levels.
	ContrastControl
	_
	// Send "Control Q" followed by one byte for the column (0-19 for a 20x4 display, or 0-15 for a 16x2 display),
	// and a second byte for the row (0-3 for a 4x20 or 0-1 for a 2x16). The upper-left position is 0,0.
	// The lower-right position is 15,1 for a 16x2, and 19,3 for a 20x4.
	SetCursorPosition

	HorizontalBarGraph
	// Turns Scroll feature ON. Then a Line Feed (Control J) command from the bottom row will scroll
	// the display up by one row, independent of Wrap . If Wrap is also on, a wrap occuring on the bottom
	// row will cause the display to scroll up one row. Scroll is on at power up.
	ScrollOn
	// Turns Scroll feature OFF. Then a Line Feed (Control J) command from the bottom row will move the
	// cursor to the top row of of the same column, independent of Wrap. If Wrap is also on, a wrap occuring
	// on the bottom row will also wrap vertically to the top row. Scroll is on at power up.
	ScrollOff
	SetScrollingMarqueeCharacters
	EnableScrollingMarquee
	// Turns Wrap feature ON. Then, a printable character received when the cursor is at the right-most
	// column will cause the cursor to move down one row, to the left-most column. If the cursor is
	// already at the right-most column of the bottom row, it will wrap to the top row if Scroll is OFF,
	// or the display will scroll up one row if Scroll is ON.
	WrapOn
	// Turns Wrap feature OFF. Then , a printable character received when the cursor is at the right-most
	// column will cause the cursor to disappear (as it will be off the right edge of the screen) and any
	// subsequent characters will be ignored until some other command moves the cursor back onto the display.
	// This function is independent of Scroll.
	WrapOff
	SetCustomCharacterBitmap
	Reboot
	EscapeSequencePrefix
	LargeBlockNumber
	_
	SendControllerData
	ShowInformationScreen

	CustomCharacter0 = 128
	CustomCharacter1 = CustomCharacter0 + 1
	CustomCharacter2 = CustomCharacter0 + 2
	CustomCharacter3 = CustomCharacter0 + 3
	CustomCharacter4 = CustomCharacter0 + 4
	CustomCharacter5 = CustomCharacter0 + 5
	CustomCharacter6 = CustomCharacter0 + 6
	CustomCharacter7 = CustomCharacter0 + 7
)

type Display struct {
	conn net.Conn
}

func Connect(addr string) (d *Display, err error) {
	conn, err := net.Dial("tcp", "10.0.0.230:8266")
	if err != nil {
		return nil, err
	}
	d = &Display{conn: conn}
	return
}

func (d *Display) Close() error {
	if d.conn != nil {
		return d.conn.Close()
	}
	return nil
}

func (d *Display) Write(data []byte) error {
	_, err := d.conn.Write(data)
	return err
}

func (d *Display) WriteByte(b byte) error {
	return d.Write([]byte{b})
}

func (d *Display) Print(s string) error {
	return d.Write([]byte(s))
}

func (d *Display) Clear() error {
	return d.WriteByte(FormFeed)
}
