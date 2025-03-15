package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1350480328958935111, false)
	EmojiOpen       = NewCustomEmoji("open", 1350480320574783498, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1350480312580182087, false)
	EmojiClose      = NewCustomEmoji("close", 1350480365411893309, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1350480357606035587, false)
	EmojiReason     = NewCustomEmoji("reason", 1350480258062745742, false)
	EmojiSubject    = NewCustomEmoji("subject", 1350480240014655571, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1350480216073441290, false)
	EmojiClaim      = NewCustomEmoji("claim", 1350480374286909512, false)
	EmojiPanel      = NewCustomEmoji("panel", 1350480303663087769, false)
	EmojiRating     = NewCustomEmoji("rating", 1350480284428144670, false)
	EmojiStaff      = NewCustomEmoji("staff", 1350480249472942080, false)
	EmojiThread     = NewCustomEmoji("thread", 1350480229080109177, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1350480383485149254, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1350480294737870899, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1350480348399534121, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
