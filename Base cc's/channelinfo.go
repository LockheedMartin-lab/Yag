{{/*
	This command allows you to view information about a given channel (defaulting to the current channel).
	Usage: `-channelinfo [channel]`.
 
	Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(channel)(-?info)?(\s+|\z)`
*/}}
 
{{ $channel := .Channel }}
{{ $args := parseArgs 0 "**Syntax:** -channel [channel]" (carg "channel" "channel") }}
{{ if $args.IsSet 0 }}
	{{ $channel = $args.Get 0 }}
{{ end }}
 
{{ $types := cslice "Text" "DM" "Voice" "Group DM" "Category" "News" "Store" }}
{{ $nsfw := "No" }}
{{ $parent := "*None set*" }}
{{ if $channel.NSFW }} {{ $nsfw = "Yes" }} {{ end }}
{{ with $channel.ParentID }} {{ $parent = printf "<#%d>" . }} {{ end }}
{{ $createdAt := div $channel.ID 4194304 | add 1420070400000 | mult 1000000 | toDuration | (newDate 1970 1 1 0 0 0).Add }}
{{ sendMessage nil (cembed
	"title" (printf "❯ Info for #%s" $channel.Name)
	"fields" (cslice
		(sdict "name" "❯ ID" "value" (str $channel.ID) "inline" true)
		(sdict "name" "❯ Topic" "value" (or $channel.Topic "*None set*") "inline" true)
		(sdict "name" "❯ Parent Channel" "value" $parent "inline" true)
		(sdict "name" "❯ NSFW" "value" $nsfw "inline" true)
		(sdict "name" "❯ Position" "value" (str (add $channel.Position 1)) "inline" true)
		(sdict "name" "❯ Type" "value" (index $types $channel.Type) "inline" true)
	)
	"color" 14232643
	"footer" (sdict "text" "Created at")
	"timestamp" $createdAt
) }}